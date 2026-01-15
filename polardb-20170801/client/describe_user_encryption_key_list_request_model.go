// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeUserEncryptionKeyListRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserEncryptionKeyListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeUserEncryptionKeyListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserEncryptionKeyListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeUserEncryptionKeyListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserEncryptionKeyListRequest
	GetResourceOwnerId() *int64
	SetTDERegion(v string) *DescribeUserEncryptionKeyListRequest
	GetTDERegion() *string
}

type DescribeUserEncryptionKeyListRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// example:
	//
	// pc-************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all regions that are available for your account, such as the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region where the TDE key resides.
	//
	// example:
	//
	// cn-beijing
	TDERegion *string `json:"TDERegion,omitempty" xml:"TDERegion,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeUserEncryptionKeyListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserEncryptionKeyListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserEncryptionKeyListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserEncryptionKeyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserEncryptionKeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserEncryptionKeyListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserEncryptionKeyListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserEncryptionKeyListRequest) GetTDERegion() *string {
	return s.TDERegion
}

func (s *DescribeUserEncryptionKeyListRequest) SetDBClusterId(v string) *DescribeUserEncryptionKeyListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetOwnerId(v int64) *DescribeUserEncryptionKeyListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v int32) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v int32) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetResourceOwnerAccount(v string) *DescribeUserEncryptionKeyListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetResourceOwnerId(v int64) *DescribeUserEncryptionKeyListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetTDERegion(v string) *DescribeUserEncryptionKeyListRequest {
	s.TDERegion = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) Validate() error {
	return dara.Validate(s)
}
