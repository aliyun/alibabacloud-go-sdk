// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetDBClusterId() *string
	SetForced(v bool) *SwitchOverGlobalDatabaseNetworkRequest
	GetForced() *bool
	SetGDNId(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchOverGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchOverGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SwitchOverGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type SwitchOverGlobalDatabaseNetworkRequest struct {
	// The ID of the cluster that will become the primary cluster in the GDN.
	//
	// You can call the [DescribeGlobalDatabaseNetwork](https://help.aliyun.com/document_detail/264580.html) operation to query the ID of the cluster in the GDN.
	//
	// example:
	//
	// pc-wz9fb5nn44u1d****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to forcibly switch over the primary and secondary clusters in the GDN. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Forced *bool `json:"Forced,omitempty" xml:"Forced,omitempty"`
	// The ID of the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId        *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchOverGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetForced() *bool {
	return s.Forced
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetDBClusterId(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.DBClusterId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetForced(v bool) *SwitchOverGlobalDatabaseNetworkRequest {
	s.Forced = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetGDNId(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *SwitchOverGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetRegionId(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetResourceGroupId(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *SwitchOverGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *SwitchOverGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchOverGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
