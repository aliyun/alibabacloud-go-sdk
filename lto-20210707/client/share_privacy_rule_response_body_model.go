// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSharePrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SharePrivacyRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SharePrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SharePrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *SharePrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SharePrivacyRuleResponseBody
	GetSuccess() *bool
}

type SharePrivacyRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SharePrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SharePrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SharePrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *SharePrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SharePrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SharePrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SharePrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SharePrivacyRuleResponseBody) SetCode(v string) *SharePrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *SharePrivacyRuleResponseBody) SetHttpStatusCode(v int32) *SharePrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SharePrivacyRuleResponseBody) SetMessage(v string) *SharePrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *SharePrivacyRuleResponseBody) SetRequestId(v string) *SharePrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SharePrivacyRuleResponseBody) SetSuccess(v bool) *SharePrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *SharePrivacyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
