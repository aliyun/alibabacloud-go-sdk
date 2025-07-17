// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSchedulerTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerSchedulerTaskInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TriggerSchedulerTaskInstanceResponseBody
	GetSuccess() *bool
}

type TriggerSchedulerTaskInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerSchedulerTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerSchedulerTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerSchedulerTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerSchedulerTaskInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerSchedulerTaskInstanceResponseBody) SetRequestId(v string) *TriggerSchedulerTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceResponseBody) SetSuccess(v bool) *TriggerSchedulerTaskInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
