// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *CreateOrderForLicenseRequest
	GetParamStr() *string
}

type CreateOrderForLicenseRequest struct {
	ParamStr *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
}

func (s CreateOrderForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForLicenseRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForLicenseRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *CreateOrderForLicenseRequest) SetParamStr(v string) *CreateOrderForLicenseRequest {
	s.ParamStr = &v
	return s
}

func (s *CreateOrderForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
