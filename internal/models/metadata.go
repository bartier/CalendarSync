package models

import (
	"fmt"
	"hash/fnv"
)

// Metadata describes the metadata which is added to events read from the source.
// The data is either calculated from the original event or given in the config file.
type Metadata struct {
	// SyncID is a unique ID which links the original event with the synced copy/copies
	SyncID string `json:"SyncID"`
	// OriginalEventUri is an URI which points to the original event which was synced. This is usually an URL.
	OriginalEventUri string `json:"OriginalEventUri"`
	// SourceID contains the unique hash of the source which this event was imported from
	SourceID string `json:"SourceID"`
}

func Hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func NewEventID(seed string) string {
	// We hash the event id, as we need some common denominator for the event IDs
	// We can't use the original event id as the event id in the sink, because the allowed formats differ
	// between the adapters.
	return fmt.Sprint(Hash(seed))
}

func NewEventMetadata(syncId, originalEventUri, sourceID string) *Metadata {
	return &Metadata{
		SyncID:           NewEventID(syncId),
		OriginalEventUri: originalEventUri,
		SourceID:         sourceID,
	}
}
