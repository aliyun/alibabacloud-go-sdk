// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSupportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeDownloadSupportRequest
	GetClusterName() *string
	SetInstanceName(v string) *DescribeDownloadSupportRequest
	GetInstanceName() *string
	SetRegionCode(v string) *DescribeDownloadSupportRequest
	GetRegionCode() *string
}

type DescribeDownloadSupportRequest struct {
	// example:
	//
	// dds-example
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1a48p922r4b****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/26231.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDownloadSupportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSupportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDownloadSupportRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDownloadSupportRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDownloadSupportRequest) SetClusterName(v string) *DescribeDownloadSupportRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDownloadSupportRequest) SetInstanceName(v string) *DescribeDownloadSupportRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadSupportRequest) SetRegionCode(v string) *DescribeDownloadSupportRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadSupportRequest) Validate() error {
	return dara.Validate(s)
}
