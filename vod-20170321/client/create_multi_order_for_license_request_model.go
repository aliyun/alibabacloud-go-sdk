// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiOrderForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *CreateMultiOrderForLicenseRequest
	GetParamStr() *string
}

type CreateMultiOrderForLicenseRequest struct {
	ParamStr *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
}

func (s CreateMultiOrderForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiOrderForLicenseRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiOrderForLicenseRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *CreateMultiOrderForLicenseRequest) SetParamStr(v string) *CreateMultiOrderForLicenseRequest {
	s.ParamStr = &v
	return s
}

func (s *CreateMultiOrderForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
