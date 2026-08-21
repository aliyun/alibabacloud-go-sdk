// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateMemberAccountsRequest
	GetInstanceId() *string
	SetMemberAccountIds(v []*string) *CreateMemberAccountsRequest
	GetMemberAccountIds() []*string
	SetRegionId(v string) *CreateMemberAccountsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateMemberAccountsRequest
	GetResourceManagerResourceGroupId() *string
	SetSourceIp(v string) *CreateMemberAccountsRequest
	GetSourceIp() *string
}

type CreateMemberAccountsRequest struct {
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of your current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of WAF member account IDs to manage. You can add up to 10 member accounts at a time.
	//
	// This parameter is required.
	MemberAccountIds []*string `json:"MemberAccountIds,omitempty" xml:"MemberAccountIds,omitempty" type:"Repeated"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The source IP address of the request. You do not need to specify this parameter. The system automatically obtains the value.
	//
	// example:
	//
	// 1.1.X.X
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateMemberAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberAccountsRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMemberAccountsRequest) GetMemberAccountIds() []*string {
	return s.MemberAccountIds
}

func (s *CreateMemberAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMemberAccountsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateMemberAccountsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateMemberAccountsRequest) SetInstanceId(v string) *CreateMemberAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetMemberAccountIds(v []*string) *CreateMemberAccountsRequest {
	s.MemberAccountIds = v
	return s
}

func (s *CreateMemberAccountsRequest) SetRegionId(v string) *CreateMemberAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetResourceManagerResourceGroupId(v string) *CreateMemberAccountsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetSourceIp(v string) *CreateMemberAccountsRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateMemberAccountsRequest) Validate() error {
	return dara.Validate(s)
}
