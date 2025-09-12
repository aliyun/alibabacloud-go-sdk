// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAutoScalingRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateAutoScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAutoScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAutoScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAutoScalingRuleResponseBody
	GetSuccess() *bool
}

type CreateAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAutoScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAutoScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAutoScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAutoScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAutoScalingRuleResponseBody) SetCode(v string) *CreateAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *CreateAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetMessage(v string) *CreateAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetRequestId(v string) *CreateAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetSuccess(v bool) *CreateAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
