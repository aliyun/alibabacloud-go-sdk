// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeKernelVersionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeKernelVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeKernelVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeKernelVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeKernelVersionRequest
	GetResourceOwnerId() *int64
}

type DescribeKernelVersionRequest struct {
	// The cluster ID.
	//
	// >
	//
	// 	- You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-2ze918p6qf6h9****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shenzhen
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeKernelVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeKernelVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeKernelVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeKernelVersionRequest) SetDBClusterId(v string) *DescribeKernelVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeKernelVersionRequest) SetOwnerAccount(v string) *DescribeKernelVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeKernelVersionRequest) SetOwnerId(v int64) *DescribeKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeKernelVersionRequest) SetRegionId(v string) *DescribeKernelVersionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKernelVersionRequest) SetResourceOwnerAccount(v string) *DescribeKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKernelVersionRequest) SetResourceOwnerId(v int64) *DescribeKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
