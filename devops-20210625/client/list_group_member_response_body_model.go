// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListGroupMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListGroupMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v []*ListGroupMemberResponseBodyResult) *ListGroupMemberResponseBody
	GetResult() []*ListGroupMemberResponseBodyResult
	SetSuccess(v bool) *ListGroupMemberResponseBody
	GetSuccess() *bool
}

type ListGroupMemberResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// C8F8B434-B006-59FB-8B9C-0382CF3D5680
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListGroupMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGroupMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupMemberResponseBody) GetResult() []*ListGroupMemberResponseBodyResult {
	return s.Result
}

func (s *ListGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGroupMemberResponseBody) SetErrorCode(v string) *ListGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetErrorMessage(v string) *ListGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetRequestId(v string) *ListGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetResult(v []*ListGroupMemberResponseBodyResult) *ListGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *ListGroupMemberResponseBody) SetSuccess(v bool) *ListGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupMemberResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupMemberResponseBodyResult struct {
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
	// 123456
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// USERS
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// test-codeup
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
	// test-codeup-nickname
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListGroupMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListGroupMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListGroupMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *ListGroupMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListGroupMemberResponseBodyResult) GetMemberType() *string {
	return s.MemberType
}

func (s *ListGroupMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListGroupMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListGroupMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListGroupMemberResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *ListGroupMemberResponseBodyResult) SetAccessLevel(v int32) *ListGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetAvatarUrl(v string) *ListGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetEmail(v string) *ListGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetId(v int64) *ListGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetMemberType(v string) *ListGroupMemberResponseBodyResult {
	s.MemberType = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetName(v string) *ListGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetSourceId(v int64) *ListGroupMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetState(v string) *ListGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetUsername(v string) *ListGroupMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
