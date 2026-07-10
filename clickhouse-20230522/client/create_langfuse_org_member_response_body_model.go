// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateLangfuseOrgMemberResponseBodyData) *CreateLangfuseOrgMemberResponseBody
	GetData() *CreateLangfuseOrgMemberResponseBodyData
	SetRequestId(v string) *CreateLangfuseOrgMemberResponseBody
	GetRequestId() *string
}

type CreateLangfuseOrgMemberResponseBody struct {
	// The information about the user added to the Langfuse organization.
	Data *CreateLangfuseOrgMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLangfuseOrgMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgMemberResponseBody) GetData() *CreateLangfuseOrgMemberResponseBodyData {
	return s.Data
}

func (s *CreateLangfuseOrgMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLangfuseOrgMemberResponseBody) SetData(v *CreateLangfuseOrgMemberResponseBodyData) *CreateLangfuseOrgMemberResponseBody {
	s.Data = v
	return s
}

func (s *CreateLangfuseOrgMemberResponseBody) SetRequestId(v string) *CreateLangfuseOrgMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLangfuseOrgMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLangfuseOrgMemberResponseBodyData struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The role of the user in the organization.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateLangfuseOrgMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgMemberResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *CreateLangfuseOrgMemberResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *CreateLangfuseOrgMemberResponseBodyData) SetEmail(v string) *CreateLangfuseOrgMemberResponseBodyData {
	s.Email = &v
	return s
}

func (s *CreateLangfuseOrgMemberResponseBodyData) SetRole(v string) *CreateLangfuseOrgMemberResponseBodyData {
	s.Role = &v
	return s
}

func (s *CreateLangfuseOrgMemberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
