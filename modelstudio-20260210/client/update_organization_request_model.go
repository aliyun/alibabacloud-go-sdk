// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateOrganizationRequest
	GetDescription() *string
	SetName(v string) *UpdateOrganizationRequest
	GetName() *string
}

type UpdateOrganizationRequest struct {
	// The organization description.
	//
	// example:
	//
	// 新的组织描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The organization name.
	//
	// example:
	//
	// 新的组织名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateOrganizationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOrganizationRequest) GetName() *string {
	return s.Name
}

func (s *UpdateOrganizationRequest) SetDescription(v string) *UpdateOrganizationRequest {
	s.Description = &v
	return s
}

func (s *UpdateOrganizationRequest) SetName(v string) *UpdateOrganizationRequest {
	s.Name = &v
	return s
}

func (s *UpdateOrganizationRequest) Validate() error {
	return dara.Validate(s)
}
