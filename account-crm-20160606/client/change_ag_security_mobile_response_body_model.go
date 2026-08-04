// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityMobileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeAgSecurityMobileResponseBody
	GetCode() *string
	SetMessage(v string) *ChangeAgSecurityMobileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeAgSecurityMobileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeAgSecurityMobileResponseBody
	GetSuccess() *bool
}

type ChangeAgSecurityMobileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAgSecurityMobileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityMobileResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityMobileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeAgSecurityMobileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeAgSecurityMobileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAgSecurityMobileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeAgSecurityMobileResponseBody) SetCode(v string) *ChangeAgSecurityMobileResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeAgSecurityMobileResponseBody) SetMessage(v string) *ChangeAgSecurityMobileResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeAgSecurityMobileResponseBody) SetRequestId(v string) *ChangeAgSecurityMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAgSecurityMobileResponseBody) SetSuccess(v bool) *ChangeAgSecurityMobileResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeAgSecurityMobileResponseBody) Validate() error {
	return dara.Validate(s)
}
