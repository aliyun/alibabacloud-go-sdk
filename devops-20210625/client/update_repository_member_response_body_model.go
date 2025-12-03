// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateRepositoryMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateRepositoryMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateRepositoryMemberResponseBody
	GetRequestId() *string
	SetResult(v *UpdateRepositoryMemberResponseBodyResult) *UpdateRepositoryMemberResponseBody
	GetResult() *UpdateRepositoryMemberResponseBodyResult
	SetSuccess(v bool) *UpdateRepositoryMemberResponseBody
	GetSuccess() *bool
}

type UpdateRepositoryMemberResponseBody struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F590C9D8-E908-5B6C-95AC-56B7E8011FFA
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateRepositoryMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRepositoryMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateRepositoryMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateRepositoryMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepositoryMemberResponseBody) GetResult() *UpdateRepositoryMemberResponseBodyResult {
	return s.Result
}

func (s *UpdateRepositoryMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorCode(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorMessage(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetRequestId(v string) *UpdateRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetResult(v *UpdateRepositoryMemberResponseBodyResult) *UpdateRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetSuccess(v bool) *UpdateRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRepositoryMemberResponseBodyResult struct {
	// example:
	//
	// 40
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// test@alibaba.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 2020-08-08 08:08:08
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test_memeber_name
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
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
	// Project
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// normal
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// codeup-test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// web url
	//
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s UpdateRepositoryMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetMemberName() *string {
	return s.MemberName
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *UpdateRepositoryMemberResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetEmail(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetExpireAt(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetMemberName(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.MemberName = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetMemberType(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.MemberType = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetName(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetSourceId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetSourceType(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetState(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetUsername(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetWebUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
