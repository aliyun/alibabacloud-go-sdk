// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCaptchaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeadline(v string) *DescribeInstallCaptchaRequest
	GetDeadline() *string
	SetLang(v string) *DescribeInstallCaptchaRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInstallCaptchaRequest
	GetSourceIp() *string
}

type DescribeInstallCaptchaRequest struct {
	// The validity period of verification codes. If this parameter is not specified, only the valid verification codes are returned.
	//
	// >  An installation verification code can be used only within the validity period. An expired installation verification code cannot be used to install the Security Center agent.
	//
	// example:
	//
	// 2020-10-11 16:26:22
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstallCaptchaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCaptchaRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaRequest) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeInstallCaptchaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstallCaptchaRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstallCaptchaRequest) SetDeadline(v string) *DescribeInstallCaptchaRequest {
	s.Deadline = &v
	return s
}

func (s *DescribeInstallCaptchaRequest) SetLang(v string) *DescribeInstallCaptchaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstallCaptchaRequest) SetSourceIp(v string) *DescribeInstallCaptchaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstallCaptchaRequest) Validate() error {
	return dara.Validate(s)
}
