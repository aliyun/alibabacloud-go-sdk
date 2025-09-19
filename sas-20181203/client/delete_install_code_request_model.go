// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstallCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaptchaCode(v string) *DeleteInstallCodeRequest
	GetCaptchaCode() *string
}

type DeleteInstallCodeRequest struct {
	// The installation command.
	//
	// >  You can call the [DescribeInstallCodes](~~DescribeInstallCodes~~) operation to query installation commands.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1q****
	CaptchaCode *string `json:"CaptchaCode,omitempty" xml:"CaptchaCode,omitempty"`
}

func (s DeleteInstallCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstallCodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstallCodeRequest) GetCaptchaCode() *string {
	return s.CaptchaCode
}

func (s *DeleteInstallCodeRequest) SetCaptchaCode(v string) *DeleteInstallCodeRequest {
	s.CaptchaCode = &v
	return s
}

func (s *DeleteInstallCodeRequest) Validate() error {
	return dara.Validate(s)
}
