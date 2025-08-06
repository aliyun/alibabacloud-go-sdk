// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCssOrderForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *QueryCssOrderForLicenseRequest
	GetParamStr() *string
}

type QueryCssOrderForLicenseRequest struct {
	ParamStr *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
}

func (s QueryCssOrderForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCssOrderForLicenseRequest) GoString() string {
	return s.String()
}

func (s *QueryCssOrderForLicenseRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *QueryCssOrderForLicenseRequest) SetParamStr(v string) *QueryCssOrderForLicenseRequest {
	s.ParamStr = &v
	return s
}

func (s *QueryCssOrderForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
