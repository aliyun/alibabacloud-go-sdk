// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgSecurityMobileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAgSecurityMobileResponseBody
	GetCode() *string
	SetMessage(v string) *QueryAgSecurityMobileResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAgSecurityMobileResponseBody
	GetRequestId() *string
	SetSecurityMobile(v string) *QueryAgSecurityMobileResponseBody
	GetSecurityMobile() *string
	SetSuccess(v bool) *QueryAgSecurityMobileResponseBody
	GetSuccess() *bool
}

type QueryAgSecurityMobileResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityMobile *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAgSecurityMobileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAgSecurityMobileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAgSecurityMobileResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAgSecurityMobileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAgSecurityMobileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAgSecurityMobileResponseBody) GetSecurityMobile() *string {
	return s.SecurityMobile
}

func (s *QueryAgSecurityMobileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAgSecurityMobileResponseBody) SetCode(v string) *QueryAgSecurityMobileResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAgSecurityMobileResponseBody) SetMessage(v string) *QueryAgSecurityMobileResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAgSecurityMobileResponseBody) SetRequestId(v string) *QueryAgSecurityMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAgSecurityMobileResponseBody) SetSecurityMobile(v string) *QueryAgSecurityMobileResponseBody {
	s.SecurityMobile = &v
	return s
}

func (s *QueryAgSecurityMobileResponseBody) SetSuccess(v bool) *QueryAgSecurityMobileResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAgSecurityMobileResponseBody) Validate() error {
	return dara.Validate(s)
}
