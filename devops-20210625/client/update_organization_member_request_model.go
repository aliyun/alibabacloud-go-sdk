// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationMemberName(v string) *UpdateOrganizationMemberRequest
	GetOrganizationMemberName() *string
}

type UpdateOrganizationMemberRequest struct {
	// This parameter is required.
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
}

func (s UpdateOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberRequest) GetOrganizationMemberName() *string {
	return s.OrganizationMemberName
}

func (s *UpdateOrganizationMemberRequest) SetOrganizationMemberName(v string) *UpdateOrganizationMemberRequest {
	s.OrganizationMemberName = &v
	return s
}

func (s *UpdateOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
