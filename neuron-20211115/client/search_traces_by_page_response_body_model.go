// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchTracesByPageResponseBody
	GetRequestId() *string
	SetTraceInfos(v []*TraceInfo) *SearchTracesByPageResponseBody
	GetTraceInfos() []*TraceInfo
}

type SearchTracesByPageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// sdadawqewe
	RequestId  *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TraceInfos []*TraceInfo `json:"traceInfos,omitempty" xml:"traceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTracesByPageResponseBody) GetTraceInfos() []*TraceInfo {
	return s.TraceInfos
}

func (s *SearchTracesByPageResponseBody) SetRequestId(v string) *SearchTracesByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesByPageResponseBody) SetTraceInfos(v []*TraceInfo) *SearchTracesByPageResponseBody {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesByPageResponseBody) Validate() error {
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
