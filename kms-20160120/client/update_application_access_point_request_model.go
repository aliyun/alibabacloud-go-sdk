// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateApplicationAccessPointRequest
	GetDescription() *string
	SetName(v string) *UpdateApplicationAccessPointRequest
	GetName() *string
	SetPolicies(v string) *UpdateApplicationAccessPointRequest
	GetPolicies() *string
}

type UpdateApplicationAccessPointRequest struct {
	// The description.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy that you want to update.
	//
	// > You can associate up to three permission policies with each AAP.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
}

func (s UpdateApplicationAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAccessPointRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationAccessPointRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationAccessPointRequest) GetPolicies() *string {
	return s.Policies
}

func (s *UpdateApplicationAccessPointRequest) SetDescription(v string) *UpdateApplicationAccessPointRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationAccessPointRequest) SetName(v string) *UpdateApplicationAccessPointRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationAccessPointRequest) SetPolicies(v string) *UpdateApplicationAccessPointRequest {
	s.Policies = &v
	return s
}

func (s *UpdateApplicationAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
