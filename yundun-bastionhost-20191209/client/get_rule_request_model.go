// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *GetRuleRequest
	GetRuleId() *string
}

type GetRuleRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-wwo35essw07
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the authorization rule to query.
	//
	// >  You can call the [ListRules](https://help.aliyun.com/document_detail/2758868.html) operation to query the authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetRuleRequest) SetInstanceId(v string) *GetRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRuleRequest) SetRegionId(v string) *GetRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetRuleRequest) SetRuleId(v string) *GetRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetRuleRequest) Validate() error {
	return dara.Validate(s)
}
