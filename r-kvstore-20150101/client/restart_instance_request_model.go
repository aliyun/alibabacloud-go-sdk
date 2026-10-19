// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *RestartInstanceRequest
	GetEffectiveTime() *string
	SetInstanceId(v string) *RestartInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RestartInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RestartInstanceRequest
	GetSecurityToken() *string
	SetUpgradeMinorVersion(v bool) *RestartInstanceRequest
	GetUpgradeMinorVersion() *bool
}

type RestartInstanceRequest struct {
	// The time when the instance is restarted. Valid values:
	//
	// 	- **Immediately**: The instance is restarted immediately. This is the default value.
	//
	// 	- **MaintainTime**: The instance is restarted during the maintenance window.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to upgrade the minor version to the latest version when the instance is restarted. Valid values:
	//
	// 	- **true**: Upgrades the minor version.
	//
	// 	- **false**: Does not upgrade the minor version.
	//
	// > Default value: **true**.
	//
	// example:
	//
	// true
	UpgradeMinorVersion *bool `json:"UpgradeMinorVersion,omitempty" xml:"UpgradeMinorVersion,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *RestartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestartInstanceRequest) GetUpgradeMinorVersion() *bool {
	return s.UpgradeMinorVersion
}

func (s *RestartInstanceRequest) SetEffectiveTime(v string) *RestartInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *RestartInstanceRequest) SetInstanceId(v string) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerAccount(v string) *RestartInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetOwnerId(v int64) *RestartInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerAccount(v string) *RestartInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceRequest) SetResourceOwnerId(v int64) *RestartInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartInstanceRequest) SetSecurityToken(v string) *RestartInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartInstanceRequest) SetUpgradeMinorVersion(v bool) *RestartInstanceRequest {
	s.UpgradeMinorVersion = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
