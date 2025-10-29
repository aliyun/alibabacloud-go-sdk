// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGDeploymentStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentType(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetDeploymentType() *string
	SetDiskIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetDiskIdsShrink() *string
	SetInstanceIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *DescribeSDGDeploymentStatusShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGDeploymentStatusShrinkRequest
	GetPageSize() *int32
	SetRegionIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetRegionIdsShrink() *string
	SetSDGId(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetSDGId() *string
	SetStatus(v string) *DescribeSDGDeploymentStatusShrinkRequest
	GetStatus() *string
}

type DescribeSDGDeploymentStatusShrinkRequest struct {
	// The deployment type.
	//
	// example:
	//
	// shared
	DeploymentType *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	DiskIdsShrink  *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// IDs of Android in Container (AIC) instances.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
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
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
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

func (s DescribeSDGDeploymentStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGDeploymentStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetDiskIdsShrink() *string {
	return s.DiskIdsShrink
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetDeploymentType(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.DeploymentType = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetDiskIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.DiskIdsShrink = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetInstanceIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetPageNumber(v int32) *DescribeSDGDeploymentStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetPageSize(v int32) *DescribeSDGDeploymentStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetRegionIdsShrink(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetSDGId(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) SetStatus(v string) *DescribeSDGDeploymentStatusShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeSDGDeploymentStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
