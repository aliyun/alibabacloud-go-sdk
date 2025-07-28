// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMemberAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyMemberAccountRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyMemberAccountRequest
	GetInstanceId() *string
	SetMemberAccountId(v string) *ModifyMemberAccountRequest
	GetMemberAccountId() *string
	SetRegionId(v string) *ModifyMemberAccountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyMemberAccountRequest
	GetResourceManagerResourceGroupId() *string
	SetSourceIp(v string) *ModifyMemberAccountRequest
	GetSourceIp() *string
}

type ModifyMemberAccountRequest struct {
	// The description of the member. The description must be 1 to 256 characters in length, and can contain letters, digits, periods (.), underscores (_), hyphens (-), and asterisks (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity**-*******021
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID of the managed member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131**********39
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system automatically obtains the value of this parameter.
	//
	// example:
	//
	// 0.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyMemberAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMemberAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyMemberAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyMemberAccountRequest) GetMemberAccountId() *string {
	return s.MemberAccountId
}

func (s *ModifyMemberAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMemberAccountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyMemberAccountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyMemberAccountRequest) SetDescription(v string) *ModifyMemberAccountRequest {
	s.Description = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetInstanceId(v string) *ModifyMemberAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetMemberAccountId(v string) *ModifyMemberAccountRequest {
	s.MemberAccountId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetRegionId(v string) *ModifyMemberAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetResourceManagerResourceGroupId(v string) *ModifyMemberAccountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetSourceIp(v string) *ModifyMemberAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyMemberAccountRequest) Validate() error {
	return dara.Validate(s)
}
