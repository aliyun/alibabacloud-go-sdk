// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteScheduledScalingRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *DeleteScheduledScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScheduledScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScheduledScalingRuleResponseBody
	GetSuccess() *bool
}

type DeleteScheduledScalingRuleResponseBody struct {
	// The responses code. The value 200 indicates that the request was successful.
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
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScheduledScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteScheduledScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScheduledScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScheduledScalingRuleResponseBody) SetCode(v int64) *DeleteScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetMessage(v string) *DeleteScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetRequestId(v string) *DeleteScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetSuccess(v bool) *DeleteScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
