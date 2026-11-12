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
	// The whitelist of IP address ranges that are allowed to access the VPC. This parameter is required. If this parameter is not specified, an InvalidParameter error is returned, indicating that AclPolicyList must be set.
	AclPolicyList []*DeleteAclPolicyRequestAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The ID of the VPC. For information about how to obtain the VPC ID, see DescribeVpcs.
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
	// The comment for the IP address range that is allowed to access the VPC.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP address range that is allowed to access the VPC.
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
