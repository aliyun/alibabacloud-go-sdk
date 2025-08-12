// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitorGroupDynamicRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PutMonitorGroupDynamicRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *PutMonitorGroupDynamicRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutMonitorGroupDynamicRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutMonitorGroupDynamicRuleResponseBody
	GetSuccess() *bool
}

type PutMonitorGroupDynamicRuleResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E73F1AB-D195-438A-BCA7-2F4355789C58
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
}

func (s PutMonitorGroupDynamicRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PutMonitorGroupDynamicRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutMonitorGroupDynamicRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutMonitorGroupDynamicRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetCode(v int32) *PutMonitorGroupDynamicRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetMessage(v string) *PutMonitorGroupDynamicRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetRequestId(v string) *PutMonitorGroupDynamicRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetSuccess(v bool) *PutMonitorGroupDynamicRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
