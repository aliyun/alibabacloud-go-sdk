// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupNotifyPolicyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetMessage() *string
	SetNotifyPolicyList(v *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetNotifyPolicyList() *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList
	SetRequestId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetSuccess() *string
	SetTotal(v int32) *DescribeMonitorGroupNotifyPolicyListResponseBody
	GetTotal() *int32
}

type DescribeMonitorGroupNotifyPolicyListResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned policies.
	NotifyPolicyList *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList `json:"NotifyPolicyList,omitempty" xml:"NotifyPolicyList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6072F026-C441-41A6-B114-35A1E8F8FDD3
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetNotifyPolicyList() *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList {
	return s.NotifyPolicyList
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetCode(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetMessage(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetNotifyPolicyList(v *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.NotifyPolicyList = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetRequestId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetSuccess(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetTotal(v int32) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList struct {
	NotifyPolicy []*DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy `json:"NotifyPolicy,omitempty" xml:"NotifyPolicy,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) GetNotifyPolicy() []*DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	return s.NotifyPolicy
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) SetNotifyPolicy(v []*DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList {
	s.NotifyPolicy = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy struct {
	// The end of the time range to query.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1551761781273
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 6780****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 123****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The beginning of the time range to query.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1551761781273
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PauseNotify
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GetId() *string {
	return s.Id
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GetType() *string {
	return s.Type
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetEndTime(v int64) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetGroupId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.Id = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetStartTime(v int64) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.StartTime = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetType(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) Validate() error {
	return dara.Validate(s)
}
