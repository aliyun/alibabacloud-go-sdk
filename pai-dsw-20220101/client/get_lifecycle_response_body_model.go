// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLifecycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLifecycleResponseBody
	GetCode() *string
	SetLifecycle(v [][]*GetLifecycleResponseBodyLifecycle) *GetLifecycleResponseBody
	GetLifecycle() [][]*GetLifecycleResponseBodyLifecycle
	SetMessage(v string) *GetLifecycleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLifecycleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLifecycleResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetLifecycleResponseBody
	GetTotalCount() *int32
}

type GetLifecycleResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: All errors, except for parameter validation errors, are internal errors.
	//
	// 	- ValidationError: A parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The lifecycle details.
	//
	// example:
	//
	// [[{"Status":"Creating","GmtCreateTime":"2022-09-19T22:38:00Z","Reason":"","ReasonCode":""}]]
	Lifecycle [][]*GetLifecycleResponseBodyLifecycle `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of queried sessions.
	//
	// example:
	//
	// 35
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLifecycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLifecycleResponseBody) GetLifecycle() [][]*GetLifecycleResponseBodyLifecycle {
	return s.Lifecycle
}

func (s *GetLifecycleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLifecycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLifecycleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLifecycleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetLifecycleResponseBody) SetCode(v string) *GetLifecycleResponseBody {
	s.Code = &v
	return s
}

func (s *GetLifecycleResponseBody) SetLifecycle(v [][]*GetLifecycleResponseBodyLifecycle) *GetLifecycleResponseBody {
	s.Lifecycle = v
	return s
}

func (s *GetLifecycleResponseBody) SetMessage(v string) *GetLifecycleResponseBody {
	s.Message = &v
	return s
}

func (s *GetLifecycleResponseBody) SetRequestId(v string) *GetLifecycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLifecycleResponseBody) SetSuccess(v bool) *GetLifecycleResponseBody {
	s.Success = &v
	return s
}

func (s *GetLifecycleResponseBody) SetTotalCount(v int32) *GetLifecycleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetLifecycleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLifecycleResponseBodyLifecycle struct {
	// The status of the instance. Valid values:
	//
	// 	- Creating
	//
	// 	- SaveFailed: The instance image failed to be saved.
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- ResourceAllocating
	//
	// 	- Stopping
	//
	// 	- Updating
	//
	// 	- Saving
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Saved
	//
	// 	- EnvPreparing: Preparing environment.
	//
	// 	- ArrearStopping: The service is being stopped due to overdue payments.
	//
	// 	- Arrearge: The service is stopped due to overdue payments.
	//
	// 	- Queuing
	//
	// 	- Recovering
	//
	// example:
	//
	// Starting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason code that corresponds to an event.
	//
	// example:
	//
	// “”
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The reason message that corresponds to an event.
	//
	// example:
	//
	// “”
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The time the status was created, specifically the time the instance transitioned to this status (in GMT).
	//
	// example:
	//
	// 2022-10-21T07:27:44Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtEndTime    *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	LifecycleId   *string `json:"LifecycleId,omitempty" xml:"LifecycleId,omitempty"`
}

func (s GetLifecycleResponseBodyLifecycle) String() string {
	return dara.Prettify(s)
}

func (s GetLifecycleResponseBodyLifecycle) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponseBodyLifecycle) GetStatus() *string {
	return s.Status
}

func (s *GetLifecycleResponseBodyLifecycle) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetLifecycleResponseBodyLifecycle) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetLifecycleResponseBodyLifecycle) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetLifecycleResponseBodyLifecycle) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *GetLifecycleResponseBodyLifecycle) GetLifecycleId() *string {
	return s.LifecycleId
}

func (s *GetLifecycleResponseBodyLifecycle) SetStatus(v string) *GetLifecycleResponseBodyLifecycle {
	s.Status = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetReasonCode(v string) *GetLifecycleResponseBodyLifecycle {
	s.ReasonCode = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetReasonMessage(v string) *GetLifecycleResponseBodyLifecycle {
	s.ReasonMessage = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetGmtCreateTime(v string) *GetLifecycleResponseBodyLifecycle {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetGmtEndTime(v string) *GetLifecycleResponseBodyLifecycle {
	s.GmtEndTime = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) SetLifecycleId(v string) *GetLifecycleResponseBodyLifecycle {
	s.LifecycleId = &v
	return s
}

func (s *GetLifecycleResponseBodyLifecycle) Validate() error {
	return dara.Validate(s)
}
