// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthInfo(v string) *CheckLicenseResponseBody
	GetAuthInfo() *string
	SetCode(v string) *CheckLicenseResponseBody
	GetCode() *string
	SetMessage(v string) *CheckLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckLicenseResponseBody
	GetRequestId() *string
}

type CheckLicenseResponseBody struct {
	AuthInfo  *string `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLicenseResponseBody) GetAuthInfo() *string {
	return s.AuthInfo
}

func (s *CheckLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckLicenseResponseBody) SetAuthInfo(v string) *CheckLicenseResponseBody {
	s.AuthInfo = &v
	return s
}

func (s *CheckLicenseResponseBody) SetCode(v string) *CheckLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *CheckLicenseResponseBody) SetMessage(v string) *CheckLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *CheckLicenseResponseBody) SetRequestId(v string) *CheckLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
