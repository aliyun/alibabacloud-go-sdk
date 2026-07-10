// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseProjectMembershipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyLangfuseProjectMembershipResponseBodyData) *ModifyLangfuseProjectMembershipResponseBody
	GetData() *ModifyLangfuseProjectMembershipResponseBodyData
	SetRequestId(v string) *ModifyLangfuseProjectMembershipResponseBody
	GetRequestId() *string
}

type ModifyLangfuseProjectMembershipResponseBody struct {
	// The returned data.
	Data *ModifyLangfuseProjectMembershipResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLangfuseProjectMembershipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseProjectMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseProjectMembershipResponseBody) GetData() *ModifyLangfuseProjectMembershipResponseBodyData {
	return s.Data
}

func (s *ModifyLangfuseProjectMembershipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLangfuseProjectMembershipResponseBody) SetData(v *ModifyLangfuseProjectMembershipResponseBodyData) *ModifyLangfuseProjectMembershipResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponseBody) SetRequestId(v string) *ModifyLangfuseProjectMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLangfuseProjectMembershipResponseBodyData struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The role of the user in the project.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ModifyLangfuseProjectMembershipResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseProjectMembershipResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseProjectMembershipResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseProjectMembershipResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *ModifyLangfuseProjectMembershipResponseBodyData) SetEmail(v string) *ModifyLangfuseProjectMembershipResponseBodyData {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponseBodyData) SetRole(v string) *ModifyLangfuseProjectMembershipResponseBodyData {
	s.Role = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponseBodyData) Validate() error {
	return dara.Validate(s)
}
