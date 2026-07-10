// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateLangfuseUserResponseBodyData) *CreateLangfuseUserResponseBody
	GetData() *CreateLangfuseUserResponseBodyData
	SetRequestId(v string) *CreateLangfuseUserResponseBody
	GetRequestId() *string
}

type CreateLangfuseUserResponseBody struct {
	// The response data.
	Data *CreateLangfuseUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLangfuseUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLangfuseUserResponseBody) GetData() *CreateLangfuseUserResponseBodyData {
	return s.Data
}

func (s *CreateLangfuseUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLangfuseUserResponseBody) SetData(v *CreateLangfuseUserResponseBodyData) *CreateLangfuseUserResponseBody {
	s.Data = v
	return s
}

func (s *CreateLangfuseUserResponseBody) SetRequestId(v string) *CreateLangfuseUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLangfuseUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLangfuseUserResponseBodyData struct {
	// The time when the user was created.
	//
	// example:
	//
	// 2026-03-04T10:20:33Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The role of the user.
	Membership *CreateLangfuseUserResponseBodyDataMembership `json:"Membership,omitempty" xml:"Membership,omitempty" type:"Struct"`
	// The username.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateLangfuseUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLangfuseUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateLangfuseUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *CreateLangfuseUserResponseBodyData) GetMembership() *CreateLangfuseUserResponseBodyDataMembership {
	return s.Membership
}

func (s *CreateLangfuseUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseUserResponseBodyData) SetCreatedAt(v string) *CreateLangfuseUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateLangfuseUserResponseBodyData) SetEmail(v string) *CreateLangfuseUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *CreateLangfuseUserResponseBodyData) SetMembership(v *CreateLangfuseUserResponseBodyDataMembership) *CreateLangfuseUserResponseBodyData {
	s.Membership = v
	return s
}

func (s *CreateLangfuseUserResponseBodyData) SetName(v string) *CreateLangfuseUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateLangfuseUserResponseBodyData) Validate() error {
	if s.Membership != nil {
		if err := s.Membership.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLangfuseUserResponseBodyDataMembership struct {
	// The Langfuse organization ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The role of the user in the organization.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateLangfuseUserResponseBodyDataMembership) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseUserResponseBodyDataMembership) GoString() string {
	return s.String()
}

func (s *CreateLangfuseUserResponseBodyDataMembership) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseUserResponseBodyDataMembership) GetRole() *string {
	return s.Role
}

func (s *CreateLangfuseUserResponseBodyDataMembership) SetOrganizationId(v string) *CreateLangfuseUserResponseBodyDataMembership {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseUserResponseBodyDataMembership) SetRole(v string) *CreateLangfuseUserResponseBodyDataMembership {
	s.Role = &v
	return s
}

func (s *CreateLangfuseUserResponseBodyDataMembership) Validate() error {
	return dara.Validate(s)
}
