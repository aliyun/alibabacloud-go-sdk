// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiOutboundTaskProgressResponseBody
	GetCode() *string
	SetData(v *GetAiOutboundTaskProgressResponseBodyData) *GetAiOutboundTaskProgressResponseBody
	GetData() *GetAiOutboundTaskProgressResponseBodyData
	SetMessage(v string) *GetAiOutboundTaskProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiOutboundTaskProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiOutboundTaskProgressResponseBody
	GetSuccess() *bool
}

type GetAiOutboundTaskProgressResponseBody struct {
	// Status code.
	//
	// example:
	//
	// ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Job progress.
	Data *GetAiOutboundTaskProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAiOutboundTaskProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiOutboundTaskProgressResponseBody) GetData() *GetAiOutboundTaskProgressResponseBodyData {
	return s.Data
}

func (s *GetAiOutboundTaskProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiOutboundTaskProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiOutboundTaskProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiOutboundTaskProgressResponseBody) SetCode(v string) *GetAiOutboundTaskProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBody) SetData(v *GetAiOutboundTaskProgressResponseBodyData) *GetAiOutboundTaskProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBody) SetMessage(v string) *GetAiOutboundTaskProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBody) SetRequestId(v string) *GetAiOutboundTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBody) SetSuccess(v bool) *GetAiOutboundTaskProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskProgressResponseBodyData struct {
	// Progress by outbound call dimension.
	CalloutProgress *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress `json:"CalloutProgress,omitempty" xml:"CalloutProgress,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Job dimension progress.
	TaskProgress *GetAiOutboundTaskProgressResponseBodyDataTaskProgress `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty" type:"Struct"`
	// Task Type. Valid values:
	//
	// - **2**: Predictive.
	//
	// - **3**: Automatic.
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiOutboundTaskProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressResponseBodyData) GetCalloutProgress() *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	return s.CalloutProgress
}

func (s *GetAiOutboundTaskProgressResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskProgressResponseBodyData) GetTaskProgress() *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	return s.TaskProgress
}

func (s *GetAiOutboundTaskProgressResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetAiOutboundTaskProgressResponseBodyData) SetCalloutProgress(v *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) *GetAiOutboundTaskProgressResponseBodyData {
	s.CalloutProgress = v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyData) SetTaskId(v int64) *GetAiOutboundTaskProgressResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyData) SetTaskProgress(v *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) *GetAiOutboundTaskProgressResponseBodyData {
	s.TaskProgress = v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyData) SetType(v int32) *GetAiOutboundTaskProgressResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyData) Validate() error {
	if s.CalloutProgress != nil {
		if err := s.CalloutProgress.Validate(); err != nil {
			return err
		}
	}
	if s.TaskProgress != nil {
		if err := s.TaskProgress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskProgressResponseBodyDataCalloutProgress struct {
	// Call loss count.
	//
	// > A parameter specific to predictive outbound calls.
	//
	// example:
	//
	// 1
	CallLossCount *int32 `json:"CallLossCount,omitempty" xml:"CallLossCount,omitempty"`
	// Call Loss Rate.
	//
	// > A parameter exclusive to predictive outbound calls.
	//
	// example:
	//
	// 0.01
	CallLossRate *float32 `json:"CallLossRate,omitempty" xml:"CallLossRate,omitempty"`
	// Number of successful outbound call connections.
	//
	// > This parameter is specific to automatic outbound calls.
	//
	// example:
	//
	// 5
	CallOutConnectCount *int32 `json:"CallOutConnectCount,omitempty" xml:"CallOutConnectCount,omitempty"`
	// Outbound Call Connection Rate.
	//
	// > A parameter exclusive to automated outbound calls.
	//
	// example:
	//
	// 0.5
	CallOutConnectRate *float32 `json:"CallOutConnectRate,omitempty" xml:"CallOutConnectRate,omitempty"`
	// Number of outbound calls.
	//
	// example:
	//
	// 10
	CallOutCount *int32 `json:"CallOutCount,omitempty" xml:"CallOutCount,omitempty"`
	// Agent Pickup Count.
	//
	// > A parameter exclusive to predictive outbound calls.
	//
	// example:
	//
	// 4
	CallOutServicerPickupCount *int32 `json:"CallOutServicerPickupCount,omitempty" xml:"CallOutServicerPickupCount,omitempty"`
	// Agent pickup rate.
	//
	// > A parameter specific to predictive outbound calls.
	//
	// example:
	//
	// 0.4
	CallOutServicerPickupRate *float32 `json:"CallOutServicerPickupRate,omitempty" xml:"CallOutServicerPickupRate,omitempty"`
	// Customer Pickup Count.
	//
	// > A parameter exclusive to predictive outbound calls.
	//
	// example:
	//
	// 5
	CallOutUserPickupCount *int32 `json:"CallOutUserPickupCount,omitempty" xml:"CallOutUserPickupCount,omitempty"`
	// Customer Pickup Rate.
	//
	// > A parameter exclusive to predictive outbound calls.
	//
	// example:
	//
	// 0.5
	CallOutUserPickupRate *float32 `json:"CallOutUserPickupRate,omitempty" xml:"CallOutUserPickupRate,omitempty"`
}

func (s GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallLossCount() *int32 {
	return s.CallLossCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallLossRate() *float32 {
	return s.CallLossRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutConnectCount() *int32 {
	return s.CallOutConnectCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutConnectRate() *float32 {
	return s.CallOutConnectRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutCount() *int32 {
	return s.CallOutCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutServicerPickupCount() *int32 {
	return s.CallOutServicerPickupCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutServicerPickupRate() *float32 {
	return s.CallOutServicerPickupRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutUserPickupCount() *int32 {
	return s.CallOutUserPickupCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) GetCallOutUserPickupRate() *float32 {
	return s.CallOutUserPickupRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallLossCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallLossCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallLossRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallLossRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutConnectCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutConnectCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutConnectRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutConnectRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutServicerPickupCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutServicerPickupCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutServicerPickupRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutServicerPickupRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutUserPickupCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutUserPickupCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) SetCallOutUserPickupRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress {
	s.CallOutUserPickupRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataCalloutProgress) Validate() error {
	return dara.Validate(s)
}

type GetAiOutboundTaskProgressResponseBodyDataTaskProgress struct {
	// Number of jobs in calling status.
	//
	// example:
	//
	// 2
	CallingCount *int32 `json:"CallingCount,omitempty" xml:"CallingCount,omitempty"`
	// Number of connected jobs.
	//
	// > Parameter specific to auto dialing.
	//
	// example:
	//
	// 2
	ConnectCount *int32 `json:"ConnectCount,omitempty" xml:"ConnectCount,omitempty"`
	// Job connection rate.
	//
	// > A parameter specific to automatic outbound calls.
	//
	// example:
	//
	// 0.2
	ConnectRate *float32 `json:"ConnectRate,omitempty" xml:"ConnectRate,omitempty"`
	// Number of completed jobs.
	//
	// example:
	//
	// 2
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// Job completion rate.
	//
	// example:
	//
	// 0.20
	FinishRate *float32 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// Number of agent pickups.
	//
	// > Parameter specific to predictive dialing.
	//
	// example:
	//
	// 2
	ServicerPickupCount *int32 `json:"ServicerPickupCount,omitempty" xml:"ServicerPickupCount,omitempty"`
	// Agent pickup rate.
	//
	// > Exclusive parameter for predictive outbound calls.
	//
	// example:
	//
	// 0.2
	ServicerPickupRate *float32 `json:"ServicerPickupRate,omitempty" xml:"ServicerPickupRate,omitempty"`
	// Number of stopped jobs.
	//
	// example:
	//
	// 2
	TerminateCount *int32 `json:"TerminateCount,omitempty" xml:"TerminateCount,omitempty"`
	// Total number of jobs.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of customer pickups.
	//
	// > Parameter specific to predictive dialing.
	//
	// example:
	//
	// 2
	UserPickupCount *int32 `json:"UserPickupCount,omitempty" xml:"UserPickupCount,omitempty"`
	// Customer pickup rate.
	//
	// > Exclusive parameter for predictive outbound calls.
	//
	// example:
	//
	// 0.2
	UserPickupRate *float32 `json:"UserPickupRate,omitempty" xml:"UserPickupRate,omitempty"`
	// Number of pending call jobs.
	//
	// example:
	//
	// 2
	WaitingCallCount *int32 `json:"WaitingCallCount,omitempty" xml:"WaitingCallCount,omitempty"`
	// Number of jobs pending redial.
	//
	// example:
	//
	// 2
	WaitingRecallCount *int32 `json:"WaitingRecallCount,omitempty" xml:"WaitingRecallCount,omitempty"`
}

func (s GetAiOutboundTaskProgressResponseBodyDataTaskProgress) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetCallingCount() *int32 {
	return s.CallingCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetConnectCount() *int32 {
	return s.ConnectCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetConnectRate() *float32 {
	return s.ConnectRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetFinishRate() *float32 {
	return s.FinishRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetServicerPickupCount() *int32 {
	return s.ServicerPickupCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetServicerPickupRate() *float32 {
	return s.ServicerPickupRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetTerminateCount() *int32 {
	return s.TerminateCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetUserPickupCount() *int32 {
	return s.UserPickupCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetUserPickupRate() *float32 {
	return s.UserPickupRate
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetWaitingCallCount() *int32 {
	return s.WaitingCallCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) GetWaitingRecallCount() *int32 {
	return s.WaitingRecallCount
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetCallingCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.CallingCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetConnectCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.ConnectCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetConnectRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.ConnectRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetFinishCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.FinishCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetFinishRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.FinishRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetServicerPickupCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.ServicerPickupCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetServicerPickupRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.ServicerPickupRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetTerminateCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.TerminateCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetTotalCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.TotalCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetUserPickupCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.UserPickupCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetUserPickupRate(v float32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.UserPickupRate = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetWaitingCallCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.WaitingCallCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) SetWaitingRecallCount(v int32) *GetAiOutboundTaskProgressResponseBodyDataTaskProgress {
	s.WaitingRecallCount = &v
	return s
}

func (s *GetAiOutboundTaskProgressResponseBodyDataTaskProgress) Validate() error {
	return dara.Validate(s)
}
