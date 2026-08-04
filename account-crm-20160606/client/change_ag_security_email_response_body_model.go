// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeAgSecurityEmailResponseBody
	GetCode() *string
	SetMessage(v string) *ChangeAgSecurityEmailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeAgSecurityEmailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeAgSecurityEmailResponseBody
	GetSuccess() *bool
}

type ChangeAgSecurityEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAgSecurityEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityEmailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeAgSecurityEmailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeAgSecurityEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAgSecurityEmailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeAgSecurityEmailResponseBody) SetCode(v string) *ChangeAgSecurityEmailResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeAgSecurityEmailResponseBody) SetMessage(v string) *ChangeAgSecurityEmailResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeAgSecurityEmailResponseBody) SetRequestId(v string) *ChangeAgSecurityEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAgSecurityEmailResponseBody) SetSuccess(v bool) *ChangeAgSecurityEmailResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeAgSecurityEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
