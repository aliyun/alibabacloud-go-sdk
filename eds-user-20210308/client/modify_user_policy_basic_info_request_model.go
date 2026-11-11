// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPolicyBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *ModifyUserPolicyBasicInfoRequest
	GetBusinessChannel() *string
	SetName(v string) *ModifyUserPolicyBasicInfoRequest
	GetName() *string
	SetPolicyId(v string) *ModifyUserPolicyBasicInfoRequest
	GetPolicyId() *string
}

type ModifyUserPolicyBasicInfoRequest struct {
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
	// deveoper-policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// up-89unlabdl****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ModifyUserPolicyBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPolicyBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPolicyBasicInfoRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *ModifyUserPolicyBasicInfoRequest) GetName() *string {
	return s.Name
}

func (s *ModifyUserPolicyBasicInfoRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyUserPolicyBasicInfoRequest) SetBusinessChannel(v string) *ModifyUserPolicyBasicInfoRequest {
	s.BusinessChannel = &v
	return s
}

func (s *ModifyUserPolicyBasicInfoRequest) SetName(v string) *ModifyUserPolicyBasicInfoRequest {
	s.Name = &v
	return s
}

func (s *ModifyUserPolicyBasicInfoRequest) SetPolicyId(v string) *ModifyUserPolicyBasicInfoRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyUserPolicyBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
