// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeRCInstancesRequest
	GetDescription() *string
	SetHostIp(v string) *DescribeRCInstancesRequest
	GetHostIp() *string
	SetImageId(v string) *DescribeRCInstancesRequest
	GetImageId() *string
	SetInstanceId(v string) *DescribeRCInstancesRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *DescribeRCInstancesRequest
	GetInstanceIds() *string
	SetInstanceName(v string) *DescribeRCInstancesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *DescribeRCInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCInstancesRequest
	GetPageSize() *int32
	SetPublicIp(v string) *DescribeRCInstancesRequest
	GetPublicIp() *string
	SetRegionId(v string) *DescribeRCInstancesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeRCInstancesRequest
	GetStatus() *string
	SetTag(v string) *DescribeRCInstancesRequest
	GetTag() *string
	SetVpcId(v string) *DescribeRCInstancesRequest
	GetVpcId() *string
}

type DescribeRCInstancesRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostIp      *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze704f*****
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// k8s-node
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number.
	//
	// Page starts from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCInstancesRequest) GetHostIp() *string {
	return s.HostIp
}

func (s *DescribeRCInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeRCInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInstancesRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCInstancesRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeRCInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCInstancesRequest) SetDescription(v string) *DescribeRCInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetHostIp(v string) *DescribeRCInstancesRequest {
	s.HostIp = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetImageId(v string) *DescribeRCInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetInstanceId(v string) *DescribeRCInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetInstanceIds(v string) *DescribeRCInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetInstanceName(v string) *DescribeRCInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetPageNumber(v int32) *DescribeRCInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetPageSize(v int32) *DescribeRCInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetPublicIp(v string) *DescribeRCInstancesRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetRegionId(v string) *DescribeRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetStatus(v string) *DescribeRCInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetTag(v string) *DescribeRCInstancesRequest {
	s.Tag = &v
	return s
}

func (s *DescribeRCInstancesRequest) SetVpcId(v string) *DescribeRCInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
