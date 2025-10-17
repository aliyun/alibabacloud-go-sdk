// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoExpireConfiguration(v *UpdateKyuubiTokenRequestAutoExpireConfiguration) *UpdateKyuubiTokenRequest
	GetAutoExpireConfiguration() *UpdateKyuubiTokenRequestAutoExpireConfiguration
	SetMemberArns(v []*string) *UpdateKyuubiTokenRequest
	GetMemberArns() []*string
	SetName(v string) *UpdateKyuubiTokenRequest
	GetName() *string
	SetToken(v string) *UpdateKyuubiTokenRequest
	GetToken() *string
	SetRegionId(v string) *UpdateKyuubiTokenRequest
	GetRegionId() *string
}

type UpdateKyuubiTokenRequest struct {
	AutoExpireConfiguration *UpdateKyuubiTokenRequestAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
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

func (s UpdateKyuubiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiTokenRequest) GetAutoExpireConfiguration() *UpdateKyuubiTokenRequestAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *UpdateKyuubiTokenRequest) GetMemberArns() []*string {
	return s.MemberArns
}

func (s *UpdateKyuubiTokenRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKyuubiTokenRequest) GetToken() *string {
	return s.Token
}

func (s *UpdateKyuubiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKyuubiTokenRequest) SetAutoExpireConfiguration(v *UpdateKyuubiTokenRequestAutoExpireConfiguration) *UpdateKyuubiTokenRequest {
	s.AutoExpireConfiguration = v
	return s
}

func (s *UpdateKyuubiTokenRequest) SetMemberArns(v []*string) *UpdateKyuubiTokenRequest {
	s.MemberArns = v
	return s
}

func (s *UpdateKyuubiTokenRequest) SetName(v string) *UpdateKyuubiTokenRequest {
	s.Name = &v
	return s
}

func (s *UpdateKyuubiTokenRequest) SetToken(v string) *UpdateKyuubiTokenRequest {
	s.Token = &v
	return s
}

func (s *UpdateKyuubiTokenRequest) SetRegionId(v string) *UpdateKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKyuubiTokenRequest) Validate() error {
	if s.AutoExpireConfiguration != nil {
		if err := s.AutoExpireConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKyuubiTokenRequestAutoExpireConfiguration struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 365
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s UpdateKyuubiTokenRequestAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiTokenRequestAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiTokenRequestAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateKyuubiTokenRequestAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *UpdateKyuubiTokenRequestAutoExpireConfiguration) SetEnable(v bool) *UpdateKyuubiTokenRequestAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateKyuubiTokenRequestAutoExpireConfiguration) SetExpireDays(v int32) *UpdateKyuubiTokenRequestAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *UpdateKyuubiTokenRequestAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}
