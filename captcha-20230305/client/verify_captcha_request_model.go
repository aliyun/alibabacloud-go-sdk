// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCaptchaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptchaVerifyParam(v string) *VerifyCaptchaRequest
	GetCaptchaVerifyParam() *string
}

type VerifyCaptchaRequest struct {
	// example:
	//
	// dsjidsjidsjkds*djsjdiskds
	CaptchaVerifyParam *string `json:"CaptchaVerifyParam,omitempty" xml:"CaptchaVerifyParam,omitempty"`
}

func (s VerifyCaptchaRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCaptchaRequest) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaRequest) GetCaptchaVerifyParam() *string {
	return s.CaptchaVerifyParam
}

func (s *VerifyCaptchaRequest) SetCaptchaVerifyParam(v string) *VerifyCaptchaRequest {
	s.CaptchaVerifyParam = &v
	return s
}

func (s *VerifyCaptchaRequest) Validate() error {
	return dara.Validate(s)
}
