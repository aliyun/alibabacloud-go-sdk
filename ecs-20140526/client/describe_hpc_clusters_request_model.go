// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHpcClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeHpcClustersRequest
	GetClientToken() *string
	SetHpcClusterIds(v string) *DescribeHpcClustersRequest
	GetHpcClusterIds() *string
	SetOwnerAccount(v string) *DescribeHpcClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHpcClustersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHpcClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHpcClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHpcClustersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHpcClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHpcClustersRequest
	GetResourceOwnerId() *int64
}

type DescribeHpcClustersRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of HPC clusters. The value can be a JSON array that consists of up to 100 HPC cluster IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["hpc-xxxxxxxxx", "hpc-yyyyyyyyy", … "hpc-zzzzzzzzz"]
	HpcClusterIds *string `json:"HpcClusterIds,omitempty" xml:"HpcClusterIds,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the HPC cluster list.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHpcClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHpcClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHpcClustersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeHpcClustersRequest) GetHpcClusterIds() *string {
	return s.HpcClusterIds
}

func (s *DescribeHpcClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHpcClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHpcClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHpcClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHpcClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHpcClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHpcClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHpcClustersRequest) SetClientToken(v string) *DescribeHpcClustersRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetHpcClusterIds(v string) *DescribeHpcClustersRequest {
	s.HpcClusterIds = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetOwnerAccount(v string) *DescribeHpcClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetOwnerId(v int64) *DescribeHpcClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetPageNumber(v int32) *DescribeHpcClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetPageSize(v int32) *DescribeHpcClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetRegionId(v string) *DescribeHpcClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetResourceOwnerAccount(v string) *DescribeHpcClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHpcClustersRequest) SetResourceOwnerId(v int64) *DescribeHpcClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHpcClustersRequest) Validate() error {
	return dara.Validate(s)
}
