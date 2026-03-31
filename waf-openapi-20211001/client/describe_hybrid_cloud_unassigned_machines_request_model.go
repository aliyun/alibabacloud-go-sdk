// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnassignedMachinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v int32) *DescribeHybridCloudUnassignedMachinesRequest
	GetClusterId() *int32
	SetHostName(v string) *DescribeHybridCloudUnassignedMachinesRequest
	GetHostName() *string
	SetInstanceId(v string) *DescribeHybridCloudUnassignedMachinesRequest
	GetInstanceId() *string
	SetIp(v string) *DescribeHybridCloudUnassignedMachinesRequest
	GetIp() *string
	SetPageNumber(v int32) *DescribeHybridCloudUnassignedMachinesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridCloudUnassignedMachinesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridCloudUnassignedMachinesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUnassignedMachinesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudUnassignedMachinesRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ClusterId *int32 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The host name.
	//
	// example:
	//
	// online-xagent1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-5yd****7009
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 1.X.X.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
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
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudUnassignedMachinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnassignedMachinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetClusterId() *int32 {
	return s.ClusterId
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetClusterId(v int32) *DescribeHybridCloudUnassignedMachinesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetHostName(v string) *DescribeHybridCloudUnassignedMachinesRequest {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetInstanceId(v string) *DescribeHybridCloudUnassignedMachinesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetIp(v string) *DescribeHybridCloudUnassignedMachinesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetPageNumber(v int32) *DescribeHybridCloudUnassignedMachinesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetPageSize(v int32) *DescribeHybridCloudUnassignedMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetRegionId(v string) *DescribeHybridCloudUnassignedMachinesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUnassignedMachinesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesRequest) Validate() error {
	return dara.Validate(s)
}
