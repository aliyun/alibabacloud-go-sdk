// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificationsForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *GetSpecificationsForLicenseRequest
	GetParamStr() *string
}

type GetSpecificationsForLicenseRequest struct {
	ParamStr *string `json:"ParamStr,omitempty" xml:"ParamStr,omitempty"`
}

func (s GetSpecificationsForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificationsForLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetSpecificationsForLicenseRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *GetSpecificationsForLicenseRequest) SetParamStr(v string) *GetSpecificationsForLicenseRequest {
	s.ParamStr = &v
	return s
}

func (s *GetSpecificationsForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
