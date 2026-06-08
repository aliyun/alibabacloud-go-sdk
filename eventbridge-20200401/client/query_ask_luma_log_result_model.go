// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAskLumaLogResult interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*AskLumaLogEntry) *QueryAskLumaLogResult
	GetEntries() []*AskLumaLogEntry
	SetHasMore(v bool) *QueryAskLumaLogResult
	GetHasMore() *bool
	SetLastKey(v string) *QueryAskLumaLogResult
	GetLastKey() *string
}

type QueryAskLumaLogResult struct {
	Entries []*AskLumaLogEntry `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	HasMore *bool              `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	LastKey *string            `json:"LastKey,omitempty" xml:"LastKey,omitempty"`
}

func (s QueryAskLumaLogResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAskLumaLogResult) GoString() string {
	return s.String()
}

func (s *QueryAskLumaLogResult) GetEntries() []*AskLumaLogEntry {
	return s.Entries
}

func (s *QueryAskLumaLogResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryAskLumaLogResult) GetLastKey() *string {
	return s.LastKey
}

func (s *QueryAskLumaLogResult) SetEntries(v []*AskLumaLogEntry) *QueryAskLumaLogResult {
	s.Entries = v
	return s
}

func (s *QueryAskLumaLogResult) SetHasMore(v bool) *QueryAskLumaLogResult {
	s.HasMore = &v
	return s
}

func (s *QueryAskLumaLogResult) SetLastKey(v string) *QueryAskLumaLogResult {
	s.LastKey = &v
	return s
}

func (s *QueryAskLumaLogResult) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
