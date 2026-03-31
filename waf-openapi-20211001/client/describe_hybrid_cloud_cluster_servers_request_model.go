// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int64) *DescribeHybridCloudClusterServersRequest
	GetClusterId() *int64
	SetGroupName(v string) *DescribeHybridCloudClusterServersRequest
	GetGroupName() *string
	SetGroupType(v string) *DescribeHybridCloudClusterServersRequest
	GetGroupType() *string
	SetHostName(v string) *DescribeHybridCloudClusterServersRequest
	GetHostName() *string
	SetInstanceId(v string) *DescribeHybridCloudClusterServersRequest
	GetInstanceId() *string
	SetIp(v string) *DescribeHybridCloudClusterServersRequest
	GetIp() *string
	SetPageNumber(v int32) *DescribeHybridCloudClusterServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridCloudClusterServersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridCloudClusterServersRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterServersRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudClusterServersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// protect
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// example:
	//
	// online-***wwq
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-******nd07
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.X.X.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudClusterServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterServersRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *DescribeHybridCloudClusterServersRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHybridCloudClusterServersRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeHybridCloudClusterServersRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudClusterServersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudClusterServersRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudClusterServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridCloudClusterServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridCloudClusterServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudClusterServersRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudClusterServersRequest) SetClusterId(v int64) *DescribeHybridCloudClusterServersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetGroupName(v string) *DescribeHybridCloudClusterServersRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetGroupType(v string) *DescribeHybridCloudClusterServersRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetHostName(v string) *DescribeHybridCloudClusterServersRequest {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetInstanceId(v string) *DescribeHybridCloudClusterServersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetIp(v string) *DescribeHybridCloudClusterServersRequest {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetPageNumber(v int32) *DescribeHybridCloudClusterServersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetPageSize(v int32) *DescribeHybridCloudClusterServersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetRegionId(v string) *DescribeHybridCloudClusterServersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudClusterServersRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersRequest) Validate() error {
	return dara.Validate(s)
}
