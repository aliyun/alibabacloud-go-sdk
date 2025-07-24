// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLivyComputeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoExpireConfiguration(v *RefreshLivyComputeTokenRequestAutoExpireConfiguration) *RefreshLivyComputeTokenRequest
	GetAutoExpireConfiguration() *RefreshLivyComputeTokenRequestAutoExpireConfiguration
	SetName(v string) *RefreshLivyComputeTokenRequest
	GetName() *string
	SetToken(v string) *RefreshLivyComputeTokenRequest
	GetToken() *string
	SetRegionId(v string) *RefreshLivyComputeTokenRequest
	GetRegionId() *string
}

type RefreshLivyComputeTokenRequest struct {
	AutoExpireConfiguration *RefreshLivyComputeTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// fe86812667f04v343
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s RefreshLivyComputeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenRequest) GetAutoExpireConfiguration() *RefreshLivyComputeTokenRequestAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *RefreshLivyComputeTokenRequest) GetName() *string {
	return s.Name
}

func (s *RefreshLivyComputeTokenRequest) GetToken() *string {
	return s.Token
}

func (s *RefreshLivyComputeTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RefreshLivyComputeTokenRequest) SetAutoExpireConfiguration(v *RefreshLivyComputeTokenRequestAutoExpireConfiguration) *RefreshLivyComputeTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetName(v string) *RefreshLivyComputeTokenRequest {
	s.Name = &v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetToken(v string) *RefreshLivyComputeTokenRequest {
	s.Token = &v
	return s
}

func (s *RefreshLivyComputeTokenRequest) SetRegionId(v string) *RefreshLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshLivyComputeTokenRequest) Validate() error {
	return dara.Validate(s)
}

type RefreshLivyComputeTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s RefreshLivyComputeTokenRequestAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RefreshLivyComputeTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) SetEnable(v bool) *RefreshLivyComputeTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *RefreshLivyComputeTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *RefreshLivyComputeTokenRequestAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}
