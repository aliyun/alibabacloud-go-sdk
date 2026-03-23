// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDTCSecurityIpHostsForSQLServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetResourceOwnerId() *int64
	SetSecurityIpHosts(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetSecurityIpHosts() *string
	SetSecurityToken(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetSecurityToken() *string
	SetWhiteListGroupName(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest
	GetWhiteListGroupName() *string
}

type ModifyDTCSecurityIpHostsForSQLServerRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SecurityIpHosts *string `json:"SecurityIpHosts,omitempty" xml:"SecurityIpHosts,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	WhiteListGroupName *string `json:"WhiteListGroupName,omitempty" xml:"WhiteListGroupName,omitempty"`
}

func (s ModifyDTCSecurityIpHostsForSQLServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDTCSecurityIpHostsForSQLServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetSecurityIpHosts() *string {
	return s.SecurityIpHosts
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) GetWhiteListGroupName() *string {
	return s.WhiteListGroupName
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetDBInstanceId(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetOwnerAccount(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetOwnerId(v int64) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetRegionId(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetResourceOwnerAccount(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetResourceOwnerId(v int64) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetSecurityIpHosts(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.SecurityIpHosts = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetSecurityToken(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) SetWhiteListGroupName(v string) *ModifyDTCSecurityIpHostsForSQLServerRequest {
	s.WhiteListGroupName = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerRequest) Validate() error {
	return dara.Validate(s)
}
