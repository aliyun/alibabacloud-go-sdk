// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceSpansLogEventList interface {
	dara.Model
	String() string
	GoString() string
	SetLogEventTagEntryList(v []*TagEntry) *TraceSpansLogEventList
	GetLogEventTagEntryList() []*TagEntry
}

type TraceSpansLogEventList struct {
	LogEventTagEntryList []*TagEntry `json:"logEventTagEntryList,omitempty" xml:"logEventTagEntryList,omitempty" type:"Repeated"`
}

func (s TraceSpansLogEventList) String() string {
	return dara.Prettify(s)
}

func (s TraceSpansLogEventList) GoString() string {
	return s.String()
}

func (s *TraceSpansLogEventList) GetLogEventTagEntryList() []*TagEntry {
	return s.LogEventTagEntryList
}

func (s *TraceSpansLogEventList) SetLogEventTagEntryList(v []*TagEntry) *TraceSpansLogEventList {
	s.LogEventTagEntryList = v
	return s
}

func (s *TraceSpansLogEventList) Validate() error {
	if s.LogEventTagEntryList != nil {
		for _, item := range s.LogEventTagEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
