// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRuleRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteRuleRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DeleteRuleRequest
	GetRuleId() *string
}

type DeleteRuleRequest struct {
	// The instance ID of the bastion host to which the authorization rule to be deleted belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host instance from which you want to delete the authorization rule.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule to be deleted.
	//
	// > You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteRuleRequest) SetInstanceId(v string) *DeleteRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRuleRequest) SetProjectId(v int64) *DeleteRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteRuleRequest) SetRegionId(v string) *DeleteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
