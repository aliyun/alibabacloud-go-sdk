// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateQualityRuleResponseBody
	GetCode() *string
	SetMessage(v string) *CreateQualityRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateQualityRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateQualityRuleResponseBody
	GetSuccess() *bool
}

type CreateQualityRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateQualityRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQualityRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateQualityRuleResponseBody) SetCode(v string) *CreateQualityRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetMessage(v string) *CreateQualityRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetRequestId(v string) *CreateQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityRuleResponseBody) SetSuccess(v bool) *CreateQualityRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
