// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAITaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAITaskStatusRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAITaskStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAITaskStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAITaskStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAITaskStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAITaskStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeAITaskStatusRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAITaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAITaskStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAITaskStatusRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAITaskStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAITaskStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAITaskStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAITaskStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAITaskStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAITaskStatusRequest) SetDBClusterId(v string) *DescribeAITaskStatusRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAITaskStatusRequest) SetOwnerAccount(v string) *DescribeAITaskStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAITaskStatusRequest) SetOwnerId(v int64) *DescribeAITaskStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAITaskStatusRequest) SetRegionId(v string) *DescribeAITaskStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAITaskStatusRequest) SetResourceOwnerAccount(v string) *DescribeAITaskStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAITaskStatusRequest) SetResourceOwnerId(v int64) *DescribeAITaskStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAITaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
