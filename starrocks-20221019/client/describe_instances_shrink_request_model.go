// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceStatus(v string) *DescribeInstancesShrinkRequest
	GetInstanceStatus() *string
	SetPageNumber(v int32) *DescribeInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *DescribeInstancesShrinkRequest
	GetTagShrink() *string
}

type DescribeInstancesShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-a0cb1c8ad6d35XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// starrocks_1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status to filter by. You can specify multiple statuses, separated by commas. Valid values:
	//
	// - unpaid
	//
	// - paid
	//
	// - creating
	//
	// - running
	//
	// - updating
	//
	// - disable
	//
	// - deleting
	//
	// - scaling_out
	//
	// - scaling_in
	//
	// - scaling_up
	//
	// - scaling_down
	//
	// - upgrading
	//
	// - modifying_config
	//
	// - enable_public_network
	//
	// - disable_public_network
	//
	// - convert_from_trial_to_official
	//
	// - restarting
	//
	// - migration_cluster_to_serverless
	//
	// - actively_disabled
	//
	// - enabling
	//
	// - agent_creating
	//
	// - agent_scaling_up
	//
	// example:
	//
	// running,creating
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmygmtrcenXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesShrinkRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *DescribeInstancesShrinkRequest) SetInstanceId(v string) *DescribeInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceName(v string) *DescribeInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceStatus(v string) *DescribeInstancesShrinkRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageNumber(v int32) *DescribeInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageSize(v int32) *DescribeInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetRegionId(v string) *DescribeInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetTagShrink(v string) *DescribeInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
