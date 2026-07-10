// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseOrgMembershipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyLangfuseOrgMembershipResponseBodyData) *ModifyLangfuseOrgMembershipResponseBody
	GetData() *ModifyLangfuseOrgMembershipResponseBodyData
	SetRequestId(v string) *ModifyLangfuseOrgMembershipResponseBody
	GetRequestId() *string
}

type ModifyLangfuseOrgMembershipResponseBody struct {
	// The response data.
	Data *ModifyLangfuseOrgMembershipResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLangfuseOrgMembershipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseOrgMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseOrgMembershipResponseBody) GetData() *ModifyLangfuseOrgMembershipResponseBodyData {
	return s.Data
}

func (s *ModifyLangfuseOrgMembershipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLangfuseOrgMembershipResponseBody) SetData(v *ModifyLangfuseOrgMembershipResponseBodyData) *ModifyLangfuseOrgMembershipResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponseBody) SetRequestId(v string) *ModifyLangfuseOrgMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLangfuseOrgMembershipResponseBodyData struct {
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
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ModifyLangfuseOrgMembershipResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseOrgMembershipResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseOrgMembershipResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseOrgMembershipResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *ModifyLangfuseOrgMembershipResponseBodyData) SetEmail(v string) *ModifyLangfuseOrgMembershipResponseBodyData {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponseBodyData) SetRole(v string) *ModifyLangfuseOrgMembershipResponseBodyData {
	s.Role = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponseBodyData) Validate() error {
	return dara.Validate(s)
}
