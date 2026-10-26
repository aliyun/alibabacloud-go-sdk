// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableRuleRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DisableRuleRequest
	GetProjectId() *int64
	SetRegionId(v string) *DisableRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DisableRuleRequest
	GetRuleId() *string
}

type DisableRuleRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule that you want to disable.
	//
	// > You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DisableRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DisableRuleRequest) SetInstanceId(v string) *DisableRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableRuleRequest) SetProjectId(v int64) *DisableRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *DisableRuleRequest) SetRegionId(v string) *DisableRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DisableRuleRequest) SetRuleId(v string) *DisableRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DisableRuleRequest) Validate() error {
	return dara.Validate(s)
}
