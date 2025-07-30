// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *DescribeSubscriptionMetaRequest
	GetDtsInstanceId() *string
	SetRegionId(v string) *DescribeSubscriptionMetaRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSubscriptionMetaRequest
	GetResourceGroupId() *string
	SetSid(v string) *DescribeSubscriptionMetaRequest
	GetSid() *string
	SetSubMigrationJobIds(v map[string]interface{}) *DescribeSubscriptionMetaRequest
	GetSubMigrationJobIds() map[string]interface{}
	SetTopics(v map[string]interface{}) *DescribeSubscriptionMetaRequest
	GetTopics() map[string]interface{}
}

type DescribeSubscriptionMetaRequest struct {
	// The ID of the distributed change tracking instance.
	//
	// example:
	//
	// dtsbr4m9luv2******
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the region in which the change tracking instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// z38m91gg2******
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The IDs of all subtasks in the distributed change tracking task. Separate multiple subtask IDs with commas (,).
	//
	// >  You must specify at least one of the SubMigrationJobIds and **Topics*	- parameters. We recommend that you specify the SubMigrationJobIds parameter.
	//
	// example:
	//
	// z38m91gg2******
	SubMigrationJobIds map[string]interface{} `json:"SubMigrationJobIds,omitempty" xml:"SubMigrationJobIds,omitempty"`
	// The topics of all subtasks in the distributed change tracking task. Separate multiple topics with commas (,).
	//
	// >  You must specify at least one of the **SubMigrationJobIds*	- and Topics parameters. We recommend that you specify the **SubMigrationJobIds*	- parameter.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1n0x0x5tz******_dtstestdata_version2
	Topics map[string]interface{} `json:"Topics,omitempty" xml:"Topics,omitempty"`
}

func (s DescribeSubscriptionMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeSubscriptionMetaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionMetaRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSubscriptionMetaRequest) GetSid() *string {
	return s.Sid
}

func (s *DescribeSubscriptionMetaRequest) GetSubMigrationJobIds() map[string]interface{} {
	return s.SubMigrationJobIds
}

func (s *DescribeSubscriptionMetaRequest) GetTopics() map[string]interface{} {
	return s.Topics
}

func (s *DescribeSubscriptionMetaRequest) SetDtsInstanceId(v string) *DescribeSubscriptionMetaRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetRegionId(v string) *DescribeSubscriptionMetaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetResourceGroupId(v string) *DescribeSubscriptionMetaRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetSid(v string) *DescribeSubscriptionMetaRequest {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetSubMigrationJobIds(v map[string]interface{}) *DescribeSubscriptionMetaRequest {
	s.SubMigrationJobIds = v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetTopics(v map[string]interface{}) *DescribeSubscriptionMetaRequest {
	s.Topics = v
	return s
}

func (s *DescribeSubscriptionMetaRequest) Validate() error {
	return dara.Validate(s)
}
