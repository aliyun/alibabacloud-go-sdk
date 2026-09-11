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
	// Retrieves instances by instance status. Separate multiple instance statuses with commas. Valid values:
	//
	// <ul>
	//
	// <li>unpaid: Pending payment.</li>
	//
	// <li>paid: Paid.</li>
	//
	// <li>creating: Being created.</li>
	//
	// <li>running: Running.</li>
	//
	// <li>updating: Being upgraded.</li>
	//
	// <li>disable: Unavailable.</li>
	//
	// <li>deleting: Being deleted.</li>
	//
	// <li>scaling_out: Scaling out.</li>
	//
	// <li>scaling_in: Scaling in.</li>
	//
	// <li>scaling_up: Specifications are being upgraded.</li>
	//
	// <li>scaling_down: Specifications are being used to decrease the quota.</li>
	//
	// <li>upgrading: Version is being upgraded.</li>
	//
	// <li>modifying_config: Configuration is being updated.</li>
	//
	// <li>enable_public_network: Public network access is being enabled.</li>
	//
	// <li>disable_public_network: Public network access is being shutdown.</li>
	//
	// <li>convert_from_trial_to_official: The instance edition is being upgraded.</li>
	//
	// <li>restarting: The cluster is restarting.</li>
	//
	// <li>migration_cluster_to_serverless: The cluster is being migrated.</li>
	//
	// <li>actively_disabled: The instance is stopped.</li>
	//
	// <li>enabling: The instance is being recovered.</li>
	//
	// <li>agent_creating: The agent is being created.</li>
	//
	// <li>agent_scaling_up: The agent specifications are being upgraded.</li>
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
	// The number of entries per page for a paged query.
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
	// The tags.
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
