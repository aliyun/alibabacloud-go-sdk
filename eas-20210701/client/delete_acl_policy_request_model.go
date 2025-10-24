// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclPolicyList(v []*DeleteAclPolicyRequestAclPolicyList) *DeleteAclPolicyRequest
	GetAclPolicyList() []*DeleteAclPolicyRequestAclPolicyList
	SetVpcId(v string) *DeleteAclPolicyRequest
	GetVpcId() *string
}

type DeleteAclPolicyRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	AclPolicyList []*DeleteAclPolicyRequestAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteAclPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyRequest) GetAclPolicyList() []*DeleteAclPolicyRequestAclPolicyList {
	return s.AclPolicyList
}

func (s *DeleteAclPolicyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteAclPolicyRequest) SetAclPolicyList(v []*DeleteAclPolicyRequestAclPolicyList) *DeleteAclPolicyRequest {
	s.AclPolicyList = v
	return s
}

func (s *DeleteAclPolicyRequest) SetVpcId(v string) *DeleteAclPolicyRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteAclPolicyRequest) Validate() error {
	if s.AclPolicyList != nil {
		for _, item := range s.AclPolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteAclPolicyRequestAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s DeleteAclPolicyRequestAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclPolicyRequestAclPolicyList) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyRequestAclPolicyList) GetComment() *string {
	return s.Comment
}

func (s *DeleteAclPolicyRequestAclPolicyList) GetEntry() *string {
	return s.Entry
}

func (s *DeleteAclPolicyRequestAclPolicyList) SetComment(v string) *DeleteAclPolicyRequestAclPolicyList {
	s.Comment = &v
	return s
}

func (s *DeleteAclPolicyRequestAclPolicyList) SetEntry(v string) *DeleteAclPolicyRequestAclPolicyList {
	s.Entry = &v
	return s
}

func (s *DeleteAclPolicyRequestAclPolicyList) Validate() error {
	return dara.Validate(s)
}
