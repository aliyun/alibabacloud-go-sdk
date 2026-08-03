// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *LookupEventsResponseBody
	GetEndTime() *string
	SetEvents(v []map[string]interface{}) *LookupEventsResponseBody
	GetEvents() []map[string]interface{}
	SetNextToken(v string) *LookupEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *LookupEventsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *LookupEventsResponseBody
	GetStartTime() *string
}

type LookupEventsResponseBody struct {
	// The end of the time range of the retrieved events.
	//
	// example:
	//
	// 2020-07-22T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of retrieved events.
	Events []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// > If NextToken is empty, no next page exists.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FD79665A-CE8B-49D4-82E6-5EE2E0E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range of the retrieved events.
	//
	// example:
	//
	// 2020-07-15T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LookupEventsResponseBody) GoString() string {
	return s.String()
}

func (s *LookupEventsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *LookupEventsResponseBody) GetEvents() []map[string]interface{} {
	return s.Events
}

func (s *LookupEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LookupEventsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *LookupEventsResponseBody) SetEndTime(v string) *LookupEventsResponseBody {
	s.EndTime = &v
	return s
}

func (s *LookupEventsResponseBody) SetEvents(v []map[string]interface{}) *LookupEventsResponseBody {
	s.Events = v
	return s
}

func (s *LookupEventsResponseBody) SetNextToken(v string) *LookupEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *LookupEventsResponseBody) SetRequestId(v string) *LookupEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupEventsResponseBody) SetStartTime(v string) *LookupEventsResponseBody {
	s.StartTime = &v
	return s
}

func (s *LookupEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
