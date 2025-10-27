// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*GetExecutionHistoryResponseBodyEvents) *GetExecutionHistoryResponseBody
	GetEvents() []*GetExecutionHistoryResponseBodyEvents
	SetNextToken(v string) *GetExecutionHistoryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetExecutionHistoryResponseBody
	GetRequestId() *string
}

type GetExecutionHistoryResponseBody struct {
	// The events.
	Events []*GetExecutionHistoryResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// You do not need to specify this parameter for the first request. The returned value of **ScheduleEventId*	- is used as the token for the next query. No value is returned for the last query.
	//
	// example:
	//
	// 3
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecutionHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBody) GetEvents() []*GetExecutionHistoryResponseBodyEvents {
	return s.Events
}

func (s *GetExecutionHistoryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetExecutionHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecutionHistoryResponseBody) SetEvents(v []*GetExecutionHistoryResponseBodyEvents) *GetExecutionHistoryResponseBody {
	s.Events = v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetNextToken(v string) *GetExecutionHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetRequestId(v string) *GetExecutionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutionHistoryResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExecutionHistoryResponseBodyEvents struct {
	// The details about the execution step.
	//
	// example:
	//
	// {}
	EventDetail *string `json:"EventDetail,omitempty" xml:"EventDetail,omitempty"`
	// The ID of the execution step.
	//
	// example:
	//
	// 2
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the scheduling step.
	//
	// example:
	//
	// 1
	ScheduleEventId *int64 `json:"ScheduleEventId,omitempty" xml:"ScheduleEventId,omitempty"`
	// The name of the execution step.
	//
	// example:
	//
	// passStep
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The time when the event was updated.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the execution step. Valid values:
	//
	// 	- **StepEntered**
	//
	// 	- **StepStarted**
	//
	// 	- **StepSucceeded**
	//
	// 	- **StepFailed**
	//
	// 	- **StepExited**
	//
	// 	- **BranchEntered**
	//
	// 	- **BranchExited**
	//
	// 	- **IterationEntered**
	//
	// 	- **IterationExited**
	//
	// 	- **TaskScheduled**
	//
	// 	- **TaskStarted**
	//
	// 	- **TaskSubmitted**
	//
	// 	- **TaskSubmitFailed**
	//
	// 	- **TaskSucceeded**
	//
	// 	- **TaskFailed**
	//
	// 	- **TaskTimedOut**
	//
	// 	- **ExecutionStarted**
	//
	// 	- **ExecutionStopped**
	//
	// 	- **ExecutionSucceeded**
	//
	// 	- **ExecutionFailed**
	//
	// 	- **ExecutionTimedOut**
	//
	// example:
	//
	// TaskSucceeded
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExecutionHistoryResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionHistoryResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBodyEvents) GetEventDetail() *string {
	return s.EventDetail
}

func (s *GetExecutionHistoryResponseBodyEvents) GetEventId() *int64 {
	return s.EventId
}

func (s *GetExecutionHistoryResponseBodyEvents) GetScheduleEventId() *int64 {
	return s.ScheduleEventId
}

func (s *GetExecutionHistoryResponseBodyEvents) GetStepName() *string {
	return s.StepName
}

func (s *GetExecutionHistoryResponseBodyEvents) GetTime() *string {
	return s.Time
}

func (s *GetExecutionHistoryResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventDetail(v string) *GetExecutionHistoryResponseBodyEvents {
	s.EventDetail = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetScheduleEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.ScheduleEventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetStepName(v string) *GetExecutionHistoryResponseBodyEvents {
	s.StepName = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetTime(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetType(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
