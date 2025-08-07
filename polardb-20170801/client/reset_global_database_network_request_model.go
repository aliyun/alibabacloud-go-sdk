// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ResetGlobalDatabaseNetworkRequest
	GetDBClusterId() *string
	SetGDNId(v string) *ResetGlobalDatabaseNetworkRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *ResetGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ResetGlobalDatabaseNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResetGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ResetGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type ResetGlobalDatabaseNetworkRequest struct {
	// The ID of the cluster in the GDN.
	//
	// >  You can call the [DescribeGlobalDatabaseNetwork](https://help.aliyun.com/document_detail/264580.html) operation to view the ID of the cluster in the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-wz9fb5nn44u1d****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// The ID of the region.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *ResetGlobalDatabaseNetworkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ResetGlobalDatabaseNetworkRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *ResetGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetGlobalDatabaseNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResetGlobalDatabaseNetworkRequest) SetDBClusterId(v string) *ResetGlobalDatabaseNetworkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetGDNId(v string) *ResetGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *ResetGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *ResetGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetRegionId(v string) *ResetGlobalDatabaseNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *ResetGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *ResetGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *ResetGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
