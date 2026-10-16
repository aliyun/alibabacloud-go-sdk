// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicySetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePolicySetRequest
	GetDescription() *string
	SetPolicySetName(v string) *CreatePolicySetRequest
	GetPolicySetName() *string
	SetSourcePlatform(v string) *CreatePolicySetRequest
	GetSourcePlatform() *string
}

type CreatePolicySetRequest struct {
	// The description.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy set. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
	// The source platform.
	//
	// example:
	//
	// Manuel
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
}

func (s CreatePolicySetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicySetRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicySetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicySetRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicySetRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreatePolicySetRequest) SetDescription(v string) *CreatePolicySetRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicySetRequest) SetPolicySetName(v string) *CreatePolicySetRequest {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicySetRequest) SetSourcePlatform(v string) *CreatePolicySetRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CreatePolicySetRequest) Validate() error {
	return dara.Validate(s)
}
