// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *DescribeHybridCloudGroupsRequest
	GetClusterId() *int64
	SetClusterProxyType(v string) *DescribeHybridCloudGroupsRequest
	GetClusterProxyType() *string
	SetGroupName(v int32) *DescribeHybridCloudGroupsRequest
	GetGroupName() *int32
	SetGroupType(v string) *DescribeHybridCloudGroupsRequest
	GetGroupType() *string
	SetInstanceId(v string) *DescribeHybridCloudGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeHybridCloudGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridCloudGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridCloudGroupsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudGroupsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudGroupsRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 428
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of proxy cluster that is used. Valid values:
	//
	// 	- **service**: service-based traffic mirroring.
	//
	// 	- **cname**: reverse proxy.
	//
	// example:
	//
	// cname
	ClusterProxyType *string `json:"ClusterProxyType,omitempty" xml:"ClusterProxyType,omitempty"`
	// The name of the node group that you want to query.
	//
	// example:
	//
	// groupName1
	GroupName *int32 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// 	- **protect**
	//
	// 	- **control**
	//
	// 	- **storage**
	//
	// 	- **controlStorage**
	//
	// example:
	//
	// protect
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-********w0b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DescribeHybridCloudGroupsRequest) GetClusterProxyType() *string {
	return s.ClusterProxyType
}

func (s *DescribeHybridCloudGroupsRequest) GetGroupName() *int32 {
	return s.GroupName
}

func (s *DescribeHybridCloudGroupsRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeHybridCloudGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridCloudGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridCloudGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudGroupsRequest) SetClusterId(v int64) *DescribeHybridCloudGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetClusterProxyType(v string) *DescribeHybridCloudGroupsRequest {
	s.ClusterProxyType = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetGroupName(v int32) *DescribeHybridCloudGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetGroupType(v string) *DescribeHybridCloudGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetInstanceId(v string) *DescribeHybridCloudGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetPageNumber(v int32) *DescribeHybridCloudGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetPageSize(v int32) *DescribeHybridCloudGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetRegionId(v string) *DescribeHybridCloudGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) Validate() error {
	return dara.Validate(s)
}
