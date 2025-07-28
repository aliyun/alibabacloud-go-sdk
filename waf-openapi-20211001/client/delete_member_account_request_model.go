// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteMemberAccountRequest
	GetInstanceId() *string
	SetMemberAccountId(v string) *DeleteMemberAccountRequest
	GetMemberAccountId() *string
	SetRegionId(v string) *DeleteMemberAccountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteMemberAccountRequest
	GetResourceManagerResourceGroupId() *string
	SetSourceIp(v string) *DeleteMemberAccountRequest
	GetSourceIp() *string
}

type DeleteMemberAccountRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_esasdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID of the managed member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131***********39
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
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

func (s DeleteMemberAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteMemberAccountRequest) GetMemberAccountId() *string {
	return s.MemberAccountId
}

func (s *DeleteMemberAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMemberAccountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteMemberAccountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteMemberAccountRequest) SetInstanceId(v string) *DeleteMemberAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetMemberAccountId(v string) *DeleteMemberAccountRequest {
	s.MemberAccountId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetRegionId(v string) *DeleteMemberAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetResourceManagerResourceGroupId(v string) *DeleteMemberAccountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetSourceIp(v string) *DeleteMemberAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteMemberAccountRequest) Validate() error {
	return dara.Validate(s)
}
