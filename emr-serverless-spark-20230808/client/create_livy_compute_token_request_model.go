// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoExpireConfiguration(v *CreateLivyComputeTokenRequestAutoExpireConfiguration) *CreateLivyComputeTokenRequest
	GetAutoExpireConfiguration() *CreateLivyComputeTokenRequestAutoExpireConfiguration
	SetName(v string) *CreateLivyComputeTokenRequest
	GetName() *string
	SetToken(v string) *CreateLivyComputeTokenRequest
	GetToken() *string
	SetRegionId(v string) *CreateLivyComputeTokenRequest
	GetRegionId() *string
}

type CreateLivyComputeTokenRequest struct {
	AutoExpireConfiguration *CreateLivyComputeTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// mytoken
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// gs3fy75w4o7hqe5s
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateLivyComputeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenRequest) GetAutoExpireConfiguration() *CreateLivyComputeTokenRequestAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *CreateLivyComputeTokenRequest) GetName() *string {
	return s.Name
}

func (s *CreateLivyComputeTokenRequest) GetToken() *string {
	return s.Token
}

func (s *CreateLivyComputeTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLivyComputeTokenRequest) SetAutoExpireConfiguration(v *CreateLivyComputeTokenRequestAutoExpireConfiguration) *CreateLivyComputeTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetName(v string) *CreateLivyComputeTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetToken(v string) *CreateLivyComputeTokenRequest {
	s.Token = &v
	return s
}

func (s *CreateLivyComputeTokenRequest) SetRegionId(v string) *CreateLivyComputeTokenRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLivyComputeTokenRequest) Validate() error {
	return dara.Validate(s)
}

type CreateLivyComputeTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s CreateLivyComputeTokenRequestAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) SetEnable(v bool) *CreateLivyComputeTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *CreateLivyComputeTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *CreateLivyComputeTokenRequestAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}
