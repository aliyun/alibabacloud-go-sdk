// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsAnalysisResourceQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterCategory(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetClusterCategory() *string
	SetClusterMode(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetClusterMode() *string
	SetNodeClass(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetNodeClass() *string
	SetNodeCount(v int32) *DescribeRdsAnalysisResourceQuotasRequest
	GetNodeCount() *int32
	SetOwnerAccount(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRdsAnalysisResourceQuotasRequest
	GetOwnerId() *int64
	SetRdsInstanceId(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetRdsInstanceId() *string
	SetRegionId(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRdsAnalysisResourceQuotasRequest
	GetResourceOwnerId() *int64
	SetStorageType(v string) *DescribeRdsAnalysisResourceQuotasRequest
	GetStorageType() *string
}

type DescribeRdsAnalysisResourceQuotasRequest struct {
	// The edition of the MySQL analytic instance.
	//
	// example:
	//
	// mixed_storage
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The mode of the MySQL analytic instance.
	//
	// example:
	//
	// flexible
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The instance type of the MySQL analytic instance.
	//
	// example:
	//
	// E32
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The number of nodes in the MySQL analytic instance.
	//
	// example:
	//
	// 5
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraDB RDS instance from which data is synchronized to the MySQL analytic instance.
	//
	// example:
	//
	// rm-2ze09tofcv39h7165
	RdsInstanceId *string `json:"RdsInstanceId,omitempty" xml:"RdsInstanceId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage type of the MySQL analytic instance.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetRdsInstanceId() *string {
	return s.RdsInstanceId
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetClusterCategory(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.ClusterCategory = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetClusterMode(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.ClusterMode = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetNodeClass(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.NodeClass = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetNodeCount(v int32) *DescribeRdsAnalysisResourceQuotasRequest {
	s.NodeCount = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetOwnerAccount(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetOwnerId(v int64) *DescribeRdsAnalysisResourceQuotasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetRdsInstanceId(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.RdsInstanceId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetRegionId(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetResourceGroupId(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetResourceOwnerAccount(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetResourceOwnerId(v int64) *DescribeRdsAnalysisResourceQuotasRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) SetStorageType(v string) *DescribeRdsAnalysisResourceQuotasRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasRequest) Validate() error {
	return dara.Validate(s)
}
