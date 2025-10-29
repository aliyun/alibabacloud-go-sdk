// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGDeploymentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DescribeSDGDeploymentStatusRequest
	GetDeploymentType() *string
	SetDiskIds(v []*string) *DescribeSDGDeploymentStatusRequest
	GetDiskIds() []*string
	SetInstanceIds(v []*string) *DescribeSDGDeploymentStatusRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeSDGDeploymentStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGDeploymentStatusRequest
	GetPageSize() *int32
	SetRegionIds(v []*string) *DescribeSDGDeploymentStatusRequest
	GetRegionIds() []*string
	SetSDGId(v string) *DescribeSDGDeploymentStatusRequest
	GetSDGId() *string
	SetStatus(v string) *DescribeSDGDeploymentStatusRequest
	GetStatus() *string
}

type DescribeSDGDeploymentStatusRequest struct {
	// The deployment type.
	//
	// example:
	//
	// shared
	DeploymentType *string   `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	DiskIds        []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// IDs of Android in Container (AIC) instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the nodes.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
	// The deployment status of the SDG.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSDGDeploymentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGDeploymentStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGDeploymentStatusRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DescribeSDGDeploymentStatusRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *DescribeSDGDeploymentStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeSDGDeploymentStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGDeploymentStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGDeploymentStatusRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *DescribeSDGDeploymentStatusRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeSDGDeploymentStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGDeploymentStatusRequest) SetDeploymentType(v string) *DescribeSDGDeploymentStatusRequest {
	s.DeploymentType = &v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetDiskIds(v []*string) *DescribeSDGDeploymentStatusRequest {
	s.DiskIds = v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetInstanceIds(v []*string) *DescribeSDGDeploymentStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetPageNumber(v int32) *DescribeSDGDeploymentStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetPageSize(v int32) *DescribeSDGDeploymentStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetRegionIds(v []*string) *DescribeSDGDeploymentStatusRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetSDGId(v string) *DescribeSDGDeploymentStatusRequest {
	s.SDGId = &v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) SetStatus(v string) *DescribeSDGDeploymentStatusRequest {
	s.Status = &v
	return s
}

func (s *DescribeSDGDeploymentStatusRequest) Validate() error {
	return dara.Validate(s)
}
