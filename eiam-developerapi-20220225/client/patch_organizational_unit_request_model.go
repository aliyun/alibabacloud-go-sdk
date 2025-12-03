// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PatchOrganizationalUnitRequest
	GetDescription() *string
	SetOrganizationalUnitName(v string) *PatchOrganizationalUnitRequest
	GetOrganizationalUnitName() *string
}

type PatchOrganizationalUnitRequest struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// test organizational unit
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
}

func (s PatchOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitRequest) GetDescription() *string {
	return s.Description
}

func (s *PatchOrganizationalUnitRequest) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *PatchOrganizationalUnitRequest) SetDescription(v string) *PatchOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *PatchOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *PatchOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *PatchOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
