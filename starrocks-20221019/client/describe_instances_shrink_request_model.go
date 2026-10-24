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
	// The instance status. You can specify multiple statuses separated by commas. Valid values:
	//
	// <ul>
	//
	// <li>unpaid: Pending payment.</li>
	//
	// <li>paid: Paid.</li>
	//
	// <li>creating: Starting.</li>
	//
	// <li>running: Running.</li>
	//
	// <li>updating: Updating.</li>
	//
	// <li>disable: Unavailable.</li>
	//
	// <li>deleting: Releasing.</li>
	//
	// <li>scaling_out: Scaling out.</li>
	//
	// <li>scaling_in: Scaling in.</li>
	//
	// <li>scaling_up: Scaling up.</li>
	//
	// <li>scaling_down: Scaling down.</li>
	//
	// <li>upgrading: Upgrading.</li>
	//
	// <li>modifying_config: Updating configurations.</li>
	//
	// <li>enable_public_network: Enabling public network access.</li>
	//
	// <li>disable_public_network: Disabling public network access.</li>
	//
	// <li>convert_from_trial_to_official: Upgrading the instance edition.</li>
	//
	// <li>restarting: Restarting the cluster.</li>
	//
	// <li>migration_cluster_to_serverless: Migrating the cluster.</li>
	//
	// <li>actively_disabled: Stopped.</li>
	//
	// <li>enabling: Resuming.</li>
	//
	// <li>agent_creating: Creating the agent.</li>
	//
	// <li>agent_scaling_up: Scaling up the agent.</li>
	//
	// </ul>
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
	// The number of entries per page in a paging query.
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
