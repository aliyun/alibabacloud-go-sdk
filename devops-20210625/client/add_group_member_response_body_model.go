// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddGroupMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddGroupMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v []*AddGroupMemberResponseBodyResult) *AddGroupMemberResponseBody
	GetResult() []*AddGroupMemberResponseBodyResult
	SetSuccess(v bool) *AddGroupMemberResponseBody
	GetSuccess() *bool
}

type AddGroupMemberResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*AddGroupMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddGroupMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGroupMemberResponseBody) GetResult() []*AddGroupMemberResponseBodyResult {
	return s.Result
}

func (s *AddGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGroupMemberResponseBody) SetErrorCode(v string) *AddGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetErrorMessage(v string) *AddGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetRequestId(v string) *AddGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetResult(v []*AddGroupMemberResponseBodyResult) *AddGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *AddGroupMemberResponseBody) SetSuccess(v bool) *AddGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddGroupMemberResponseBody) Validate() error {
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

type AddGroupMemberResponseBodyResult struct {
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

func (s AddGroupMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *AddGroupMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *AddGroupMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *AddGroupMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *AddGroupMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *AddGroupMemberResponseBodyResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *AddGroupMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *AddGroupMemberResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *AddGroupMemberResponseBodyResult) SetAccessLevel(v int32) *AddGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetAvatarUrl(v string) *AddGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetEmail(v string) *AddGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetId(v int64) *AddGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetName(v string) *AddGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetSourceId(v int64) *AddGroupMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetState(v string) *AddGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetUsername(v string) *AddGroupMemberResponseBodyResult {
	s.Username = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
