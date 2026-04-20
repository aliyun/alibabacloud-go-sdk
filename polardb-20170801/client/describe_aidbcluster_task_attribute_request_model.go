// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterTaskAttributeRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeAIDBClusterTaskAttributeRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeAIDBClusterTaskAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClusterTaskAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAIDBClusterTaskAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAIDBClusterTaskAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClusterTaskAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAIDBClusterTaskAttributeRequest struct {
	// example:
	//
	// pm-2ze99***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// RelatedAITask
	DescribeType *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAIDBClusterTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClusterTaskAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetDBClusterId(v string) *DescribeAIDBClusterTaskAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetDescribeType(v string) *DescribeAIDBClusterTaskAttributeRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetOwnerAccount(v string) *DescribeAIDBClusterTaskAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetOwnerId(v int64) *DescribeAIDBClusterTaskAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetRegionId(v string) *DescribeAIDBClusterTaskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClusterTaskAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) SetResourceOwnerId(v int64) *DescribeAIDBClusterTaskAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClusterTaskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
