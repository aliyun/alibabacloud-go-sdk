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
	SetRegionId(v string) *DisableRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DisableRuleRequest
	GetRuleId() *string
}

type DisableRuleRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule to be disabled.
	//
	// >  You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the authorization rule ID to be disabled.
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
