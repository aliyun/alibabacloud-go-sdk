// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInsightsEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInsightsEvents(v []*ListInsightsEventsResponseBodyInsightsEvents) *ListInsightsEventsResponseBody
	GetInsightsEvents() []*ListInsightsEventsResponseBodyInsightsEvents
	SetRequestId(v string) *ListInsightsEventsResponseBody
	GetRequestId() *string
}

type ListInsightsEventsResponseBody struct {
	// The details of the event.
	InsightsEvents []*ListInsightsEventsResponseBodyInsightsEvents `json:"InsightsEvents,omitempty" xml:"InsightsEvents,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6F1174DC-6085-5353-AAE7-D4ADCD******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInsightsEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInsightsEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponseBody) GetInsightsEvents() []*ListInsightsEventsResponseBodyInsightsEvents {
	return s.InsightsEvents
}

func (s *ListInsightsEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInsightsEventsResponseBody) SetInsightsEvents(v []*ListInsightsEventsResponseBodyInsightsEvents) *ListInsightsEventsResponseBody {
	s.InsightsEvents = v
	return s
}

func (s *ListInsightsEventsResponseBody) SetRequestId(v string) *ListInsightsEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInsightsEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInsightsEventsResponseBodyInsightsEvents struct {
	// The time when the event occurred. The value is a timestamp.
	//
	// example:
	//
	// 1658890560
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The description of the alert event.
	//
	// example:
	//
	// The overall response time of the [HTTP] service of the application [sd] spikes at [2022-07-27 10:57:00]
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The severity of the event.
	//
	// example:
	//
	// P3
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the application associated with the event.
	//
	// example:
	//
	// dsv9zcel92@7da413b******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The problem identifier.
	//
	// example:
	//
	// erep3o9zue@01ebe697ab70566|@1499161100890550|@cn-hangzhou|@1701841800000|@1701842040000|@daa6c51a-3c44-4d57-9548-4e212c******
	ProblemId *string `json:"ProblemId,omitempty" xml:"ProblemId,omitempty"`
	// The title of the event.
	//
	// example:
	//
	// Average response-time spikes of application services
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The type of the event.
	//
	// example:
	//
	// rtIncrease
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInsightsEventsResponseBodyInsightsEvents) String() string {
	return dara.Prettify(s)
}

func (s ListInsightsEventsResponseBodyInsightsEvents) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetDate() *int64 {
	return s.Date
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetDesc() *string {
	return s.Desc
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetLevel() *string {
	return s.Level
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetPid() *string {
	return s.Pid
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetProblemId() *string {
	return s.ProblemId
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetTitle() *string {
	return s.Title
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) GetType() *string {
	return s.Type
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetDate(v int64) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Date = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetDesc(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Desc = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetLevel(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Level = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetPid(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Pid = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetProblemId(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.ProblemId = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetTitle(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Title = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetType(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Type = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) Validate() error {
	return dara.Validate(s)
}
