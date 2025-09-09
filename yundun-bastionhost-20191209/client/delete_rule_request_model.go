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
	SetRegionId(v string) *DeleteRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DeleteRuleRequest
	GetRuleId() *string
}

type DeleteRuleRequest struct {
	// The ID of the bastion host from which you want to delete the authorization rule.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to delete the authorization rule.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule that you want to delete.
	//
	// >  You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the authorization rule ID.
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
