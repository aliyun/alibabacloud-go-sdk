// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ModifyScheduledScalingRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *ModifyScheduledScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScheduledScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyScheduledScalingRuleResponseBody
	GetSuccess() *bool
}

type ModifyScheduledScalingRuleResponseBody struct {
	// The response code.
	//
	// The value **200*	- indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScheduledScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ModifyScheduledScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScheduledScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduledScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScheduledScalingRuleResponseBody) SetCode(v int64) *ModifyScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetMessage(v string) *ModifyScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetRequestId(v string) *ModifyScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetSuccess(v bool) *ModifyScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
