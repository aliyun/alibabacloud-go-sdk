// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventTracesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEventTracesResponseBody
	GetCode() *string
	SetData(v []*QueryEventTracesResponseBodyData) *QueryEventTracesResponseBody
	GetData() []*QueryEventTracesResponseBodyData
	SetMessage(v string) *QueryEventTracesResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEventTracesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEventTracesResponseBody
	GetSuccess() *bool
}

type QueryEventTracesResponseBody struct {
	// The response code. Valid values:
	//
	// 200: The request was successful.
	//
	// Other codes: The request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the event source.
	Data []*QueryEventTracesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB934571-1F5A-5E17-91DD-E2BC3E1BFBFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEventTracesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEventTracesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEventTracesResponseBody) GetData() []*QueryEventTracesResponseBodyData {
	return s.Data
}

func (s *QueryEventTracesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEventTracesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEventTracesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEventTracesResponseBody) SetCode(v string) *QueryEventTracesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetData(v []*QueryEventTracesResponseBodyData) *QueryEventTracesResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventTracesResponseBody) SetMessage(v string) *QueryEventTracesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetRequestId(v string) *QueryEventTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEventTracesResponseBody) SetSuccess(v bool) *QueryEventTracesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEventTracesResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryEventTracesResponseBodyData struct {
	// The type of the event trace. Valid values: PutEvent, FilterEvent, and PushEvent. The value PutEvent indicates that the event was delivered. The value FilterEvent indicates that the event was filtered. The value PushEvent indicates that the event was pushed.
	//
	// example:
	//
	// PutEvent
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The execution time of the event trace.
	//
	// example:
	//
	// 1659495343896
	ActionTime *int64 `json:"ActionTime,omitempty" xml:"ActionTime,omitempty"`
	// The endpoint of the event target. This parameter is returned only if Action is set to PushEvent.
	//
	// example:
	//
	// acs:mns:cn-zhangjiakou:123456789098****:queues/testQueue
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// demo
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The event ID.
	//
	// example:
	//
	// a5747e4f-2af2-40b6-b262-d0140e995bf7
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the event source.
	//
	// example:
	//
	// cert-api
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The delay period for which the event was delivered to the event target. This parameter is returned only if Action is set to PushEvent.
	//
	// example:
	//
	// 80
	NotifyLatency *string `json:"NotifyLatency,omitempty" xml:"NotifyLatency,omitempty"`
	// The delivery status.
	//
	// example:
	//
	// [200]Ok
	NotifyStatus *string `json:"NotifyStatus,omitempty" xml:"NotifyStatus,omitempty"`
	// The time when the event was delivered to the event target. This parameter is returned only if Action is set to PushEvent.
	//
	// example:
	//
	// 1659495343896
	NotifyTime *int64 `json:"NotifyTime,omitempty" xml:"NotifyTime,omitempty"`
	// The time when the event was delivered to the event bus. This parameter is returned only if Action is set to PutEvent.
	//
	// example:
	//
	// 1659495343896
	ReceivedTime *int64 `json:"ReceivedTime,omitempty" xml:"ReceivedTime,omitempty"`
	// The time when the event rule was matched. This parameter is returned only if Action is set to FilterEvent.
	//
	// example:
	//
	// 1659495343896
	RuleMatchingTime *string `json:"RuleMatchingTime,omitempty" xml:"RuleMatchingTime,omitempty"`
	// The name of the event rule.
	//
	// example:
	//
	// ramrolechange-mns
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s QueryEventTracesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEventTracesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *QueryEventTracesResponseBodyData) GetActionTime() *int64 {
	return s.ActionTime
}

func (s *QueryEventTracesResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *QueryEventTracesResponseBodyData) GetEventBusName() *string {
	return s.EventBusName
}

func (s *QueryEventTracesResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *QueryEventTracesResponseBodyData) GetEventSource() *string {
	return s.EventSource
}

func (s *QueryEventTracesResponseBodyData) GetNotifyLatency() *string {
	return s.NotifyLatency
}

func (s *QueryEventTracesResponseBodyData) GetNotifyStatus() *string {
	return s.NotifyStatus
}

func (s *QueryEventTracesResponseBodyData) GetNotifyTime() *int64 {
	return s.NotifyTime
}

func (s *QueryEventTracesResponseBodyData) GetReceivedTime() *int64 {
	return s.ReceivedTime
}

func (s *QueryEventTracesResponseBodyData) GetRuleMatchingTime() *string {
	return s.RuleMatchingTime
}

func (s *QueryEventTracesResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *QueryEventTracesResponseBodyData) SetAction(v string) *QueryEventTracesResponseBodyData {
	s.Action = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetActionTime(v int64) *QueryEventTracesResponseBodyData {
	s.ActionTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEndpoint(v string) *QueryEventTracesResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventBusName(v string) *QueryEventTracesResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventId(v string) *QueryEventTracesResponseBodyData {
	s.EventId = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetEventSource(v string) *QueryEventTracesResponseBodyData {
	s.EventSource = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyLatency(v string) *QueryEventTracesResponseBodyData {
	s.NotifyLatency = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyStatus(v string) *QueryEventTracesResponseBodyData {
	s.NotifyStatus = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetNotifyTime(v int64) *QueryEventTracesResponseBodyData {
	s.NotifyTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetReceivedTime(v int64) *QueryEventTracesResponseBodyData {
	s.ReceivedTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetRuleMatchingTime(v string) *QueryEventTracesResponseBodyData {
	s.RuleMatchingTime = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) SetRuleName(v string) *QueryEventTracesResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *QueryEventTracesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
