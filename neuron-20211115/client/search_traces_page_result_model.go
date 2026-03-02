// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetTraceInfos(v []*TraceInfo) *SearchTracesPageResult
	GetTraceInfos() []*TraceInfo
}

type SearchTracesPageResult struct {
	TraceInfos []*TraceInfo `json:"traceInfos,omitempty" xml:"traceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesPageResult) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesPageResult) GoString() string {
	return s.String()
}

func (s *SearchTracesPageResult) GetTraceInfos() []*TraceInfo {
	return s.TraceInfos
}

func (s *SearchTracesPageResult) SetTraceInfos(v []*TraceInfo) *SearchTracesPageResult {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesPageResult) Validate() error {
	if s.TraceInfos != nil {
		for _, item := range s.TraceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
