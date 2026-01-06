// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *QueryTracedEventsRequest
	GetEndTime() *int64
	SetEventBusName(v string) *QueryTracedEventsRequest
	GetEventBusName() *string
	SetEventSource(v string) *QueryTracedEventsRequest
	GetEventSource() *string
	SetEventType(v string) *QueryTracedEventsRequest
	GetEventType() *string
	SetLimit(v int32) *QueryTracedEventsRequest
	GetLimit() *int32
	SetMatchedRule(v string) *QueryTracedEventsRequest
	GetMatchedRule() *string
	SetNextToken(v string) *QueryTracedEventsRequest
	GetNextToken() *string
	SetStartTime(v int64) *QueryTracedEventsRequest
	GetStartTime() *int64
	SetSubject(v string) *QueryTracedEventsRequest
	GetSubject() *string
}

type QueryTracedEventsRequest struct {
	// The end of the time range when event traces are queried. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1661773509000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the event bus.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// mse
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The event type.
	//
	// example:
	//
	// eventbridge:Events:HTTPEvent
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The maximum number of entries to return in a request. You can use this parameter and NextToken to implement paging.
	//
	// >  A maximum of 100 entries can be returned in a request.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event rule that is matched.
	//
	// example:
	//
	// test-mnsrule
	MatchedRule *string `json:"MatchedRule,omitempty" xml:"MatchedRule,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 1000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query event traces. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1661773509000
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Subject   *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s QueryTracedEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventsRequest) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryTracedEventsRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryTracedEventsRequest) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryTracedEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *QueryTracedEventsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryTracedEventsRequest) GetMatchedRule() *string {
	return s.MatchedRule
}

func (s *QueryTracedEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryTracedEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryTracedEventsRequest) GetSubject() *string {
	return s.Subject
}

func (s *QueryTracedEventsRequest) SetEndTime(v int64) *QueryTracedEventsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventBusName(v string) *QueryTracedEventsRequest {
	s.EventBusName = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventSource(v string) *QueryTracedEventsRequest {
	s.EventSource = &v
	return s
}

func (s *QueryTracedEventsRequest) SetEventType(v string) *QueryTracedEventsRequest {
	s.EventType = &v
	return s
}

func (s *QueryTracedEventsRequest) SetLimit(v int32) *QueryTracedEventsRequest {
	s.Limit = &v
	return s
}

func (s *QueryTracedEventsRequest) SetMatchedRule(v string) *QueryTracedEventsRequest {
	s.MatchedRule = &v
	return s
}

func (s *QueryTracedEventsRequest) SetNextToken(v string) *QueryTracedEventsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryTracedEventsRequest) SetStartTime(v int64) *QueryTracedEventsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTracedEventsRequest) SetSubject(v string) *QueryTracedEventsRequest {
	s.Subject = &v
	return s
}

func (s *QueryTracedEventsRequest) Validate() error {
	return dara.Validate(s)
}
