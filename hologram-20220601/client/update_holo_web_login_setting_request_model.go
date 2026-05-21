// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoloWebLoginSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateHoloWebLoginSettingRequest
	GetRegionId() *string
	SetAllowExternalAccountsLogin(v bool) *UpdateHoloWebLoginSettingRequest
	GetAllowExternalAccountsLogin() *bool
}

type UpdateHoloWebLoginSettingRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	AllowExternalAccountsLogin *bool `json:"allowExternalAccountsLogin,omitempty" xml:"allowExternalAccountsLogin,omitempty"`
}

func (s UpdateHoloWebLoginSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoloWebLoginSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoloWebLoginSettingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateHoloWebLoginSettingRequest) GetAllowExternalAccountsLogin() *bool {
	return s.AllowExternalAccountsLogin
}

func (s *UpdateHoloWebLoginSettingRequest) SetRegionId(v string) *UpdateHoloWebLoginSettingRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateHoloWebLoginSettingRequest) SetAllowExternalAccountsLogin(v bool) *UpdateHoloWebLoginSettingRequest {
	s.AllowExternalAccountsLogin = &v
	return s
}

func (s *UpdateHoloWebLoginSettingRequest) Validate() error {
	return dara.Validate(s)
}
