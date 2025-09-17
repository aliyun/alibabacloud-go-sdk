// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSdkServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeHybridCloudSdkServersRequest
	GetClusterName() *string
	SetHostName(v string) *DescribeHybridCloudSdkServersRequest
	GetHostName() *string
	SetInstanceId(v string) *DescribeHybridCloudSdkServersRequest
	GetInstanceId() *string
	SetIp(v string) *DescribeHybridCloudSdkServersRequest
	GetIp() *string
	SetPageNumber(v int32) *DescribeHybridCloudSdkServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridCloudSdkServersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridCloudSdkServersRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudSdkServersRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudSdkServersRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-i7m2***0b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
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

func (s DescribeHybridCloudSdkServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSdkServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSdkServersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridCloudSdkServersRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudSdkServersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudSdkServersRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudSdkServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridCloudSdkServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridCloudSdkServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudSdkServersRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudSdkServersRequest) SetClusterName(v string) *DescribeHybridCloudSdkServersRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetHostName(v string) *DescribeHybridCloudSdkServersRequest {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetInstanceId(v string) *DescribeHybridCloudSdkServersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetIp(v string) *DescribeHybridCloudSdkServersRequest {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetPageNumber(v int32) *DescribeHybridCloudSdkServersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetPageSize(v int32) *DescribeHybridCloudSdkServersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetRegionId(v string) *DescribeHybridCloudSdkServersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudSdkServersRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudSdkServersRequest) Validate() error {
	return dara.Validate(s)
}
