// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclPolicyList(v []*CreateAclPolicyRequestAclPolicyList) *CreateAclPolicyRequest
	GetAclPolicyList() []*CreateAclPolicyRequestAclPolicyList
	SetVpcId(v string) *CreateAclPolicyRequest
	GetVpcId() *string
}

type CreateAclPolicyRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	//
	// This parameter is required.
	AclPolicyList []*CreateAclPolicyRequestAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyRequest) GetAclPolicyList() []*CreateAclPolicyRequestAclPolicyList {
	return s.AclPolicyList
}

func (s *CreateAclPolicyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAclPolicyRequest) SetAclPolicyList(v []*CreateAclPolicyRequestAclPolicyList) *CreateAclPolicyRequest {
	s.AclPolicyList = v
	return s
}

func (s *CreateAclPolicyRequest) SetVpcId(v string) *CreateAclPolicyRequest {
	s.VpcId = &v
	return s
}

func (s *CreateAclPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAclPolicyRequestAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s CreateAclPolicyRequestAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s CreateAclPolicyRequestAclPolicyList) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyRequestAclPolicyList) GetComment() *string {
	return s.Comment
}

func (s *CreateAclPolicyRequestAclPolicyList) GetEntry() *string {
	return s.Entry
}

func (s *CreateAclPolicyRequestAclPolicyList) SetComment(v string) *CreateAclPolicyRequestAclPolicyList {
	s.Comment = &v
	return s
}

func (s *CreateAclPolicyRequestAclPolicyList) SetEntry(v string) *CreateAclPolicyRequestAclPolicyList {
	s.Entry = &v
	return s
}

func (s *CreateAclPolicyRequestAclPolicyList) Validate() error {
	return dara.Validate(s)
}
