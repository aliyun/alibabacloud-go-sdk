// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLoginPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DescribeUserLoginPoliciesRequest
	GetBusinessChannel() *string
	SetName(v string) *DescribeUserLoginPoliciesRequest
	GetName() *string
	SetPolicyIds(v []*string) *DescribeUserLoginPoliciesRequest
	GetPolicyIds() []*string
}

type DescribeUserLoginPoliciesRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Security audit
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of policy IDs. This is a filter condition for the query.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s DescribeUserLoginPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLoginPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLoginPoliciesRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeUserLoginPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeUserLoginPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DescribeUserLoginPoliciesRequest) SetBusinessChannel(v string) *DescribeUserLoginPoliciesRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeUserLoginPoliciesRequest) SetName(v string) *DescribeUserLoginPoliciesRequest {
	s.Name = &v
	return s
}

func (s *DescribeUserLoginPoliciesRequest) SetPolicyIds(v []*string) *DescribeUserLoginPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DescribeUserLoginPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
