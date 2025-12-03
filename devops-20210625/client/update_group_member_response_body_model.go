// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateGroupMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateGroupMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v *UpdateGroupMemberResponseBodyResult) *UpdateGroupMemberResponseBody
	GetResult() *UpdateGroupMemberResponseBodyResult
	SetSuccess(v bool) *UpdateGroupMemberResponseBody
	GetSuccess() *bool
}

type UpdateGroupMemberResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateGroupMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGroupMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupMemberResponseBody) GetResult() *UpdateGroupMemberResponseBodyResult {
	return s.Result
}

func (s *UpdateGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGroupMemberResponseBody) SetErrorCode(v string) *UpdateGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetErrorMessage(v string) *UpdateGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetRequestId(v string) *UpdateGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetResult(v *UpdateGroupMemberResponseBodyResult) *UpdateGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetSuccess(v bool) *UpdateGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGroupMemberResponseBodyResult struct {
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 24661
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// USERS
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 223241
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateGroupMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *UpdateGroupMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdateGroupMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *UpdateGroupMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateGroupMemberResponseBodyResult) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateGroupMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateGroupMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *UpdateGroupMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *UpdateGroupMemberResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *UpdateGroupMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetEmail(v string) *UpdateGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetId(v int64) *UpdateGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetMemberType(v string) *UpdateGroupMemberResponseBodyResult {
	s.MemberType = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetName(v string) *UpdateGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetSourceId(v int64) *UpdateGroupMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetState(v string) *UpdateGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetUsername(v string) *UpdateGroupMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
