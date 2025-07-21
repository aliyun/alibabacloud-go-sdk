// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationMethod(v string) *CreateApplicationAccessPointRequest
	GetAuthenticationMethod() *string
	SetDescription(v string) *CreateApplicationAccessPointRequest
	GetDescription() *string
	SetName(v string) *CreateApplicationAccessPointRequest
	GetName() *string
	SetPolicies(v string) *CreateApplicationAccessPointRequest
	GetPolicies() *string
}

type CreateApplicationAccessPointRequest struct {
	// The authentication method. Currently, only ClientKey is supported.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The description of the AAP.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy.
	//
	// > You can bind up to three permission policies to each AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
}

func (s CreateApplicationAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointRequest) GetAuthenticationMethod() *string {
	return s.AuthenticationMethod
}

func (s *CreateApplicationAccessPointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationAccessPointRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationAccessPointRequest) GetPolicies() *string {
	return s.Policies
}

func (s *CreateApplicationAccessPointRequest) SetAuthenticationMethod(v string) *CreateApplicationAccessPointRequest {
	s.AuthenticationMethod = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetDescription(v string) *CreateApplicationAccessPointRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetName(v string) *CreateApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) SetPolicies(v string) *CreateApplicationAccessPointRequest {
	s.Policies = &v
	return s
}

func (s *CreateApplicationAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
