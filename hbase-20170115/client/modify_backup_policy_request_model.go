// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyBackupPolicyRequest
	GetClientToken() *string
	SetClusterId(v string) *ModifyBackupPolicyRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupEndTimeUTC() *string
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupStartTimeUTC() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetRegionId(v string) *ModifyBackupPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyBackupPolicyRequest
	GetZoneId() *string
}

type ModifyBackupPolicyRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId                 *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupEndTimeUTC *string `json:"PreferredBackupEndTimeUTC,omitempty" xml:"PreferredBackupEndTimeUTC,omitempty"`
	// This parameter is required.
	PreferredBackupPeriod       *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupStartTimeUTC *string `json:"PreferredBackupStartTimeUTC,omitempty" xml:"PreferredBackupStartTimeUTC,omitempty"`
	// This parameter is required.
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBackupPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupEndTimeUTC() *string {
	return s.PreferredBackupEndTimeUTC
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupStartTimeUTC() *string {
	return s.PreferredBackupStartTimeUTC
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyBackupPolicyRequest) SetClientToken(v string) *ModifyBackupPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetClusterId(v string) *ModifyBackupPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupEndTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupEndTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupStartTimeUTC(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupStartTimeUTC = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRegionId(v string) *ModifyBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetZoneId(v string) *ModifyBackupPolicyRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
