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
	// The end of the time range when event details were queried.
	//
	// example:
	//
	// 2020-07-22T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The returned event details.
	//
	// For more information about the fields in an event log, see [ActionTrail event log reference](https://help.aliyun.com/document_detail/28819.html).
	Events []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The token used to return the next page of query results.
	//
	// > This parameter is not returned if no more results are to be returned.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FD79665A-CE8B-49D4-82E6-5EE2E0E791DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range when event details were queried.
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
