// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePrivacyRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdatePrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePrivacyRuleResponseBody
	GetSuccess() *bool
}

type UpdatePrivacyRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePrivacyRuleResponseBody) SetCode(v string) *UpdatePrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrivacyRuleResponseBody) SetHttpStatusCode(v int32) *UpdatePrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePrivacyRuleResponseBody) SetMessage(v string) *UpdatePrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrivacyRuleResponseBody) SetRequestId(v string) *UpdatePrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivacyRuleResponseBody) SetSuccess(v bool) *UpdatePrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePrivacyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
