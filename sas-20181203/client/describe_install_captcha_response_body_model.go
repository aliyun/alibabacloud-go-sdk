// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCaptchaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaptchaCode(v string) *DescribeInstallCaptchaResponseBody
	GetCaptchaCode() *string
	SetDeadline(v string) *DescribeInstallCaptchaResponseBody
	GetDeadline() *string
	SetRequestId(v string) *DescribeInstallCaptchaResponseBody
	GetRequestId() *string
}

type DescribeInstallCaptchaResponseBody struct {
	// The installation verification code for manually installing the Security Center Agent.
	//
	// example:
	//
	// M1HH**
	CaptchaCode *string `json:"CaptchaCode,omitempty" xml:"CaptchaCode,omitempty"`
	// The expiration date of the installation verification code.
	//
	// > The installation verification code can be used only within its validity period. An expired installation verification code cannot be used.
	//
	// example:
	//
	// 2020-10-10 16:06:38
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The unique identifier that Alibaba Cloud generated for the request.
	//
	// example:
	//
	// 4E5BFDCF-B9DD-430D-9DA4-151BCB581C9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstallCaptchaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaResponseBody) GetCaptchaCode() *string {
	return s.CaptchaCode
}

func (s *DescribeInstallCaptchaResponseBody) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeInstallCaptchaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstallCaptchaResponseBody) SetCaptchaCode(v string) *DescribeInstallCaptchaResponseBody {
	s.CaptchaCode = &v
	return s
}

func (s *DescribeInstallCaptchaResponseBody) SetDeadline(v string) *DescribeInstallCaptchaResponseBody {
	s.Deadline = &v
	return s
}

func (s *DescribeInstallCaptchaResponseBody) SetRequestId(v string) *DescribeInstallCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstallCaptchaResponseBody) Validate() error {
	return dara.Validate(s)
}
