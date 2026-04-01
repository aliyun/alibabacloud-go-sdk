// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustinsResourceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIds(v string) *DescribeCustinsResourceInfoRequest
	GetDBInstanceIds() *string
	SetOwnerId(v int64) *DescribeCustinsResourceInfoRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeCustinsResourceInfoRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCustinsResourceInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCustinsResourceInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeCustinsResourceInfoRequest struct {
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-wz9s06u4drmqj4aqv
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCustinsResourceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustinsResourceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustinsResourceInfoRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *DescribeCustinsResourceInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCustinsResourceInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCustinsResourceInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCustinsResourceInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCustinsResourceInfoRequest) SetDBInstanceIds(v string) *DescribeCustinsResourceInfoRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeCustinsResourceInfoRequest) SetOwnerId(v int64) *DescribeCustinsResourceInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCustinsResourceInfoRequest) SetResourceGroupId(v string) *DescribeCustinsResourceInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCustinsResourceInfoRequest) SetResourceOwnerAccount(v string) *DescribeCustinsResourceInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCustinsResourceInfoRequest) SetResourceOwnerId(v int64) *DescribeCustinsResourceInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCustinsResourceInfoRequest) Validate() error {
	return dara.Validate(s)
}
