// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateScheduledScalingRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *CreateScheduledScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduledScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScheduledScalingRuleResponseBody
	GetSuccess() *bool
}

type CreateScheduledScalingRuleResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScheduledScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateScheduledScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduledScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScheduledScalingRuleResponseBody) SetCode(v int64) *CreateScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetMessage(v string) *CreateScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetRequestId(v string) *CreateScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetSuccess(v bool) *CreateScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
