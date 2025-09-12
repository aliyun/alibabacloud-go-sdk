// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAutoScalingRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyAutoScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAutoScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAutoScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAutoScalingRuleResponseBody
	GetSuccess() *bool
}

type ModifyAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAutoScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAutoScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAutoScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAutoScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAutoScalingRuleResponseBody) SetCode(v string) *ModifyAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *ModifyAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetMessage(v string) *ModifyAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetRequestId(v string) *ModifyAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetSuccess(v bool) *ModifyAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
