// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSettingValue(v string) *UpdatePrometheusUserSettingRequest
	GetSettingValue() *string
}

type UpdatePrometheusUserSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc
	SettingValue *string `json:"settingValue,omitempty" xml:"settingValue,omitempty"`
}

func (s UpdatePrometheusUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusUserSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusUserSettingRequest) GetSettingValue() *string {
	return s.SettingValue
}

func (s *UpdatePrometheusUserSettingRequest) SetSettingValue(v string) *UpdatePrometheusUserSettingRequest {
	s.SettingValue = &v
	return s
}

func (s *UpdatePrometheusUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
