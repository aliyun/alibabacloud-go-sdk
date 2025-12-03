// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteGroupMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteGroupMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v *DeleteGroupMemberResponseBodyResult) *DeleteGroupMemberResponseBody
	GetResult() *DeleteGroupMemberResponseBodyResult
	SetSuccess(v bool) *DeleteGroupMemberResponseBody
	GetSuccess() *bool
}

type DeleteGroupMemberResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 4D6AF7CC-B43B-5454-86AB-023D25E44868
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteGroupMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteGroupMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupMemberResponseBody) GetResult() *DeleteGroupMemberResponseBodyResult {
	return s.Result
}

func (s *DeleteGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGroupMemberResponseBody) SetErrorCode(v string) *DeleteGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetErrorMessage(v string) *DeleteGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetRequestId(v string) *DeleteGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetResult(v *DeleteGroupMemberResponseBodyResult) *DeleteGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetSuccess(v bool) *DeleteGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGroupMemberResponseBodyResult struct {
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
	// 524836
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
	// 2080398
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

func (s DeleteGroupMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *DeleteGroupMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *DeleteGroupMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *DeleteGroupMemberResponseBodyResult) GetMemberType() *string {
	return s.MemberType
}

func (s *DeleteGroupMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DeleteGroupMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *DeleteGroupMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *DeleteGroupMemberResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *DeleteGroupMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetAvatarUrl(v string) *DeleteGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetId(v int64) *DeleteGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetMemberType(v string) *DeleteGroupMemberResponseBodyResult {
	s.MemberType = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetName(v string) *DeleteGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetSourceId(v int64) *DeleteGroupMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetState(v string) *DeleteGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetUsername(v string) *DeleteGroupMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
