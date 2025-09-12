// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSingleZoneToMultiZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArbitraryVSwitchId(v string) *MigrateSingleZoneToMultiZoneRequest
	GetArbitraryVSwitchId() *string
	SetArbitraryZoneId(v string) *MigrateSingleZoneToMultiZoneRequest
	GetArbitraryZoneId() *string
	SetDryRun(v bool) *MigrateSingleZoneToMultiZoneRequest
	GetDryRun() *bool
	SetInstanceId(v string) *MigrateSingleZoneToMultiZoneRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *MigrateSingleZoneToMultiZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateSingleZoneToMultiZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateSingleZoneToMultiZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateSingleZoneToMultiZoneRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *MigrateSingleZoneToMultiZoneRequest
	GetSecurityToken() *string
	SetStandbyVSwitchId(v string) *MigrateSingleZoneToMultiZoneRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *MigrateSingleZoneToMultiZoneRequest
	GetStandbyZoneId() *string
}

type MigrateSingleZoneToMultiZoneRequest struct {
	ArbitraryVSwitchId *string `json:"ArbitraryVSwitchId,omitempty" xml:"ArbitraryVSwitchId,omitempty"`
	ArbitraryZoneId    *string `json:"ArbitraryZoneId,omitempty" xml:"ArbitraryZoneId,omitempty"`
	DryRun             *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	StandbyZoneId        *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
}

func (s MigrateSingleZoneToMultiZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateSingleZoneToMultiZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetArbitraryVSwitchId() *string {
	return s.ArbitraryVSwitchId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetArbitraryZoneId() *string {
	return s.ArbitraryZoneId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *MigrateSingleZoneToMultiZoneRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetArbitraryVSwitchId(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.ArbitraryVSwitchId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetArbitraryZoneId(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.ArbitraryZoneId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetDryRun(v bool) *MigrateSingleZoneToMultiZoneRequest {
	s.DryRun = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetInstanceId(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetOwnerAccount(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetOwnerId(v int64) *MigrateSingleZoneToMultiZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetResourceOwnerAccount(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetResourceOwnerId(v int64) *MigrateSingleZoneToMultiZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetSecurityToken(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.SecurityToken = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetStandbyVSwitchId(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) SetStandbyZoneId(v string) *MigrateSingleZoneToMultiZoneRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *MigrateSingleZoneToMultiZoneRequest) Validate() error {
	return dara.Validate(s)
}
