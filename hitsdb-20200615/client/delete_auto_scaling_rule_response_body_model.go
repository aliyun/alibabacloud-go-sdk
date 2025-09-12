// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAutoScalingRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAutoScalingRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAutoScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAutoScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAutoScalingRuleResponseBody
	GetSuccess() *bool
}

type DeleteAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutoScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAutoScalingRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAutoScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAutoScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAutoScalingRuleResponseBody) SetCode(v string) *DeleteAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *DeleteAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetMessage(v string) *DeleteAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetRequestId(v string) *DeleteAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetSuccess(v bool) *DeleteAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
