// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserLoginPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *RemoveUserLoginPoliciesRequest
	GetBusinessChannel() *string
	SetPolicyIds(v []*string) *RemoveUserLoginPoliciesRequest
	GetPolicyIds() []*string
}

type RemoveUserLoginPoliciesRequest struct {
	// The business channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The list of policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s RemoveUserLoginPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserLoginPoliciesRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserLoginPoliciesRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *RemoveUserLoginPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *RemoveUserLoginPoliciesRequest) SetBusinessChannel(v string) *RemoveUserLoginPoliciesRequest {
	s.BusinessChannel = &v
	return s
}

func (s *RemoveUserLoginPoliciesRequest) SetPolicyIds(v []*string) *RemoveUserLoginPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *RemoveUserLoginPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
