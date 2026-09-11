// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *DescribeSubscriptionMetaShrinkRequest
	GetDtsInstanceId() *string
	SetRegionId(v string) *DescribeSubscriptionMetaShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSubscriptionMetaShrinkRequest
	GetResourceGroupId() *string
	SetSid(v string) *DescribeSubscriptionMetaShrinkRequest
	GetSid() *string
	SetSubMigrationJobIdsShrink(v string) *DescribeSubscriptionMetaShrinkRequest
	GetSubMigrationJobIdsShrink() *string
	SetTopicsShrink(v string) *DescribeSubscriptionMetaShrinkRequest
	GetTopicsShrink() *string
}

type DescribeSubscriptionMetaShrinkRequest struct {
	// The instance ID of the distributed change tracking task.
	//
	// > This parameter is required.
	//
	// example:
	//
	// dtsbr4m9luv2******
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The region in which the change tracking instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The consumer group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// z38m91gg2******
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The IDs of all change tracking subtasks in the distributed change tracking task. Separate multiple IDs with commas (,).
	//
	// > You must specify at least one of this parameter and **Topics**. We recommend that you specify this parameter.
	//
	// example:
	//
	// ["zsls58agp6f****"]
	SubMigrationJobIdsShrink *string `json:"SubMigrationJobIds,omitempty" xml:"SubMigrationJobIds,omitempty"`
	// All topics of the distributed change tracking task. Separate multiple topics with commas (,).
	//
	// > You must specify at least one of this parameter and **SubMigrationJobIds**. We recommend that you specify **SubMigrationJobIds**.
	//
	// example:
	//
	// ["rm_bp15jj3qi1p8f****"]
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
}

func (s DescribeSubscriptionMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetSid() *string {
	return s.Sid
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetSubMigrationJobIdsShrink() *string {
	return s.SubMigrationJobIdsShrink
}

func (s *DescribeSubscriptionMetaShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetDtsInstanceId(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetRegionId(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetResourceGroupId(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetSid(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetSubMigrationJobIdsShrink(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.SubMigrationJobIdsShrink = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetTopicsShrink(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
