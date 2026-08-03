// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupInsightEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []map[string]interface{}) *LookupInsightEventsResponseBody
	GetEvents() []map[string]interface{}
	SetNextToken(v string) *LookupInsightEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *LookupInsightEventsResponseBody
	GetRequestId() *string
}

type LookupInsightEventsResponseBody struct {
	// The Insights events.
	Events []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// VjE6bHJlTGoxdm1M****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 851038F3-33AB-4C49-97D7-6AB37D35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupInsightEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LookupInsightEventsResponseBody) GoString() string {
	return s.String()
}

func (s *LookupInsightEventsResponseBody) GetEvents() []map[string]interface{} {
	return s.Events
}

func (s *LookupInsightEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *LookupInsightEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LookupInsightEventsResponseBody) SetEvents(v []map[string]interface{}) *LookupInsightEventsResponseBody {
	s.Events = v
	return s
}

func (s *LookupInsightEventsResponseBody) SetNextToken(v string) *LookupInsightEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *LookupInsightEventsResponseBody) SetRequestId(v string) *LookupInsightEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupInsightEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
