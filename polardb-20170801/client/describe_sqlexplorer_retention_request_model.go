// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLExplorerRetentionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSQLExplorerRetentionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLExplorerRetentionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSQLExplorerRetentionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSQLExplorerRetentionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLExplorerRetentionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeSQLExplorerRetentionRequest
	GetSecurityToken() *string
}

type DescribeSQLExplorerRetentionRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters under your account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions of all clusters under the target account.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSQLExplorerRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerRetentionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerRetentionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLExplorerRetentionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLExplorerRetentionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLExplorerRetentionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSQLExplorerRetentionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLExplorerRetentionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLExplorerRetentionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSQLExplorerRetentionRequest) SetDBInstanceId(v string) *DescribeSQLExplorerRetentionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetOwnerAccount(v string) *DescribeSQLExplorerRetentionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetOwnerId(v int64) *DescribeSQLExplorerRetentionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetRegionId(v string) *DescribeSQLExplorerRetentionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetResourceOwnerAccount(v string) *DescribeSQLExplorerRetentionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetResourceOwnerId(v int64) *DescribeSQLExplorerRetentionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) SetSecurityToken(v string) *DescribeSQLExplorerRetentionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSQLExplorerRetentionRequest) Validate() error {
	return dara.Validate(s)
}
