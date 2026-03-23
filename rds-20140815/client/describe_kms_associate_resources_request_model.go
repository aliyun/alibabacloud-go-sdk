// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsAssociateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeKmsAssociateResourcesRequest
	GetClientToken() *string
	SetKmsResourceId(v string) *DescribeKmsAssociateResourcesRequest
	GetKmsResourceId() *string
	SetKmsResourceRegionId(v string) *DescribeKmsAssociateResourcesRequest
	GetKmsResourceRegionId() *string
	SetKmsResourceType(v string) *DescribeKmsAssociateResourcesRequest
	GetKmsResourceType() *string
	SetKmsResourceUser(v string) *DescribeKmsAssociateResourcesRequest
	GetKmsResourceUser() *string
	SetOwnerAccount(v string) *DescribeKmsAssociateResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DescribeKmsAssociateResourcesRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeKmsAssociateResourcesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeKmsAssociateResourcesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeKmsAssociateResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeKmsAssociateResourcesRequest
	GetResourceOwnerId() *int64
}

type DescribeKmsAssociateResourcesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	KmsResourceId       *string `json:"KmsResourceId,omitempty" xml:"KmsResourceId,omitempty"`
	KmsResourceRegionId *string `json:"KmsResourceRegionId,omitempty" xml:"KmsResourceRegionId,omitempty"`
	// This parameter is required.
	KmsResourceType *string `json:"KmsResourceType,omitempty" xml:"KmsResourceType,omitempty"`
	// This parameter is required.
	KmsResourceUser      *string `json:"KmsResourceUser,omitempty" xml:"KmsResourceUser,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeKmsAssociateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsAssociateResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKmsAssociateResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeKmsAssociateResourcesRequest) GetKmsResourceId() *string {
	return s.KmsResourceId
}

func (s *DescribeKmsAssociateResourcesRequest) GetKmsResourceRegionId() *string {
	return s.KmsResourceRegionId
}

func (s *DescribeKmsAssociateResourcesRequest) GetKmsResourceType() *string {
	return s.KmsResourceType
}

func (s *DescribeKmsAssociateResourcesRequest) GetKmsResourceUser() *string {
	return s.KmsResourceUser
}

func (s *DescribeKmsAssociateResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeKmsAssociateResourcesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeKmsAssociateResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKmsAssociateResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeKmsAssociateResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeKmsAssociateResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeKmsAssociateResourcesRequest) SetClientToken(v string) *DescribeKmsAssociateResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetKmsResourceId(v string) *DescribeKmsAssociateResourcesRequest {
	s.KmsResourceId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetKmsResourceRegionId(v string) *DescribeKmsAssociateResourcesRequest {
	s.KmsResourceRegionId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetKmsResourceType(v string) *DescribeKmsAssociateResourcesRequest {
	s.KmsResourceType = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetKmsResourceUser(v string) *DescribeKmsAssociateResourcesRequest {
	s.KmsResourceUser = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetOwnerAccount(v string) *DescribeKmsAssociateResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetOwnerId(v string) *DescribeKmsAssociateResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetRegionId(v string) *DescribeKmsAssociateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetResourceGroupId(v string) *DescribeKmsAssociateResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetResourceOwnerAccount(v string) *DescribeKmsAssociateResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) SetResourceOwnerId(v int64) *DescribeKmsAssociateResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKmsAssociateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
