// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoExpireConfiguration(v *CreateKyuubiTokenRequestAutoExpireConfiguration) *CreateKyuubiTokenRequest
	GetAutoExpireConfiguration() *CreateKyuubiTokenRequestAutoExpireConfiguration
	SetMemberArns(v []*string) *CreateKyuubiTokenRequest
	GetMemberArns() []*string
	SetName(v string) *CreateKyuubiTokenRequest
	GetName() *string
	SetToken(v string) *CreateKyuubiTokenRequest
	GetToken() *string
	SetRegionId(v string) *CreateKyuubiTokenRequest
	GetRegionId() *string
}

type CreateKyuubiTokenRequest struct {
	AutoExpireConfiguration *CreateKyuubiTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	MemberArns              []*string                                        `json:"memberArns,omitempty" xml:"memberArns,omitempty" type:"Repeated"`
	// example:
	//
	// dev_serverless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rjy7ejhej9gkzjjuun49jnx2xk8if2cu
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateKyuubiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateKyuubiTokenRequest) GetAutoExpireConfiguration() *CreateKyuubiTokenRequestAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *CreateKyuubiTokenRequest) GetMemberArns() []*string {
	return s.MemberArns
}

func (s *CreateKyuubiTokenRequest) GetName() *string {
	return s.Name
}

func (s *CreateKyuubiTokenRequest) GetToken() *string {
	return s.Token
}

func (s *CreateKyuubiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKyuubiTokenRequest) SetAutoExpireConfiguration(v *CreateKyuubiTokenRequestAutoExpireConfiguration) *CreateKyuubiTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *CreateKyuubiTokenRequest) SetMemberArns(v []*string) *CreateKyuubiTokenRequest {
	s.MemberArns = v
	return s
}

func (s *CreateKyuubiTokenRequest) SetName(v string) *CreateKyuubiTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateKyuubiTokenRequest) SetToken(v string) *CreateKyuubiTokenRequest {
	s.Token = &v
	return s
}

func (s *CreateKyuubiTokenRequest) SetRegionId(v string) *CreateKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKyuubiTokenRequest) Validate() error {
	return dara.Validate(s)
}

type CreateKyuubiTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 365
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s CreateKyuubiTokenRequestAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *CreateKyuubiTokenRequestAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateKyuubiTokenRequestAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *CreateKyuubiTokenRequestAutoExpireConfiguration) SetEnable(v bool) *CreateKyuubiTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateKyuubiTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *CreateKyuubiTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *CreateKyuubiTokenRequestAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}
