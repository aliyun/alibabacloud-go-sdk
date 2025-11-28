// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddPrivacyRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddPrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddPrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddPrivacyRuleResponseBody
	GetSuccess() *bool
}

type AddPrivacyRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddPrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddPrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddPrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddPrivacyRuleResponseBody) SetCode(v string) *AddPrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddPrivacyRuleResponseBody) SetHttpStatusCode(v int32) *AddPrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPrivacyRuleResponseBody) SetMessage(v string) *AddPrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddPrivacyRuleResponseBody) SetRequestId(v string) *AddPrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrivacyRuleResponseBody) SetSuccess(v bool) *AddPrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *AddPrivacyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
