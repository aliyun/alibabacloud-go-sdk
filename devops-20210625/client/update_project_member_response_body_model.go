// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateProjectMemberResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateProjectMemberResponseBody
	GetErrorMsg() *string
	SetMember(v *UpdateProjectMemberResponseBodyMember) *UpdateProjectMemberResponseBody
	GetMember() *UpdateProjectMemberResponseBodyMember
	SetRequestId(v string) *UpdateProjectMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectMemberResponseBody
	GetSuccess() *bool
}

type UpdateProjectMemberResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Member   *UpdateProjectMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProjectMemberResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateProjectMemberResponseBody) GetMember() *UpdateProjectMemberResponseBodyMember {
	return s.Member
}

func (s *UpdateProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectMemberResponseBody) SetErrorCode(v string) *UpdateProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetErrorMsg(v string) *UpdateProjectMemberResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetMember(v *UpdateProjectMemberResponseBodyMember) *UpdateProjectMemberResponseBody {
	s.Member = v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetRequestId(v string) *UpdateProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) SetSuccess(v bool) *UpdateProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectMemberResponseBody) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectMemberResponseBodyMember struct {
	// example:
	//
	// 1623916393000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1623916393000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1124382
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// project.admin
	RoleIdentifier *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	// example:
	//
	// 5e70xxxxxxcd000xxxxe96
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// example:
	//
	// Space
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	UserIdentifier *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	// example:
	//
	// user
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s UpdateProjectMemberResponseBodyMember) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponseBodyMember) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *UpdateProjectMemberResponseBodyMember) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *UpdateProjectMemberResponseBodyMember) GetId() *string {
	return s.Id
}

func (s *UpdateProjectMemberResponseBodyMember) GetRoleIdentifier() *string {
	return s.RoleIdentifier
}

func (s *UpdateProjectMemberResponseBodyMember) GetTargetIdentifier() *string {
	return s.TargetIdentifier
}

func (s *UpdateProjectMemberResponseBodyMember) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateProjectMemberResponseBodyMember) GetUserIdentifier() *string {
	return s.UserIdentifier
}

func (s *UpdateProjectMemberResponseBodyMember) GetUserType() *string {
	return s.UserType
}

func (s *UpdateProjectMemberResponseBodyMember) SetGmtCreate(v int64) *UpdateProjectMemberResponseBodyMember {
	s.GmtCreate = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetGmtModified(v int64) *UpdateProjectMemberResponseBodyMember {
	s.GmtModified = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetId(v string) *UpdateProjectMemberResponseBodyMember {
	s.Id = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetRoleIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.RoleIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetTargetIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetTargetType(v string) *UpdateProjectMemberResponseBodyMember {
	s.TargetType = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetUserIdentifier(v string) *UpdateProjectMemberResponseBodyMember {
	s.UserIdentifier = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) SetUserType(v string) *UpdateProjectMemberResponseBodyMember {
	s.UserType = &v
	return s
}

func (s *UpdateProjectMemberResponseBodyMember) Validate() error {
	return dara.Validate(s)
}
