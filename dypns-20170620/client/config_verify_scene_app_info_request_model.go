// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigVerifySceneAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCM(v bool) *ConfigVerifySceneAppInfoRequest
	GetCM() *bool
	SetCT(v bool) *ConfigVerifySceneAppInfoRequest
	GetCT() *bool
	SetCU(v bool) *ConfigVerifySceneAppInfoRequest
	GetCU() *bool
	SetEmail(v string) *ConfigVerifySceneAppInfoRequest
	GetEmail() *string
	SetIpWhitelist(v string) *ConfigVerifySceneAppInfoRequest
	GetIpWhitelist() *string
	SetOwnerId(v int64) *ConfigVerifySceneAppInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ConfigVerifySceneAppInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConfigVerifySceneAppInfoRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *ConfigVerifySceneAppInfoRequest
	GetSceneCode() *string
}

type ConfigVerifySceneAppInfoRequest struct {
	// example:
	//
	// false
	CM *bool `json:"CM,omitempty" xml:"CM,omitempty"`
	// example:
	//
	// false
	CT *bool `json:"CT,omitempty" xml:"CT,omitempty"`
	// example:
	//
	// false
	CU *bool `json:"CU,omitempty" xml:"CU,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx@xxx.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 112.124.18.99
	IpWhitelist          *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FC220000011285006
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s ConfigVerifySceneAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigVerifySceneAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ConfigVerifySceneAppInfoRequest) GetCM() *bool {
	return s.CM
}

func (s *ConfigVerifySceneAppInfoRequest) GetCT() *bool {
	return s.CT
}

func (s *ConfigVerifySceneAppInfoRequest) GetCU() *bool {
	return s.CU
}

func (s *ConfigVerifySceneAppInfoRequest) GetEmail() *string {
	return s.Email
}

func (s *ConfigVerifySceneAppInfoRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ConfigVerifySceneAppInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConfigVerifySceneAppInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConfigVerifySceneAppInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConfigVerifySceneAppInfoRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *ConfigVerifySceneAppInfoRequest) SetCM(v bool) *ConfigVerifySceneAppInfoRequest {
	s.CM = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetCT(v bool) *ConfigVerifySceneAppInfoRequest {
	s.CT = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetCU(v bool) *ConfigVerifySceneAppInfoRequest {
	s.CU = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetEmail(v string) *ConfigVerifySceneAppInfoRequest {
	s.Email = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetIpWhitelist(v string) *ConfigVerifySceneAppInfoRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetOwnerId(v int64) *ConfigVerifySceneAppInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetResourceOwnerAccount(v string) *ConfigVerifySceneAppInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetResourceOwnerId(v int64) *ConfigVerifySceneAppInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) SetSceneCode(v string) *ConfigVerifySceneAppInfoRequest {
	s.SceneCode = &v
	return s
}

func (s *ConfigVerifySceneAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
