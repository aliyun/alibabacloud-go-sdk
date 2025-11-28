// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePrivacyRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePrivacyRuleResponseBody
	GetSuccess() *bool
}

type DeletePrivacyRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePrivacyRuleResponseBody) SetCode(v string) *DeletePrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrivacyRuleResponseBody) SetHttpStatusCode(v int32) *DeletePrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePrivacyRuleResponseBody) SetMessage(v string) *DeletePrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrivacyRuleResponseBody) SetRequestId(v string) *DeletePrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivacyRuleResponseBody) SetSuccess(v bool) *DeletePrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePrivacyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
