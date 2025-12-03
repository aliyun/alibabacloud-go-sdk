// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRepositoryMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddRepositoryMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddRepositoryMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddRepositoryMemberResponseBody
	GetRequestId() *string
	SetResult(v []*AddRepositoryMemberResponseBodyResult) *AddRepositoryMemberResponseBody
	GetResult() []*AddRepositoryMemberResponseBodyResult
	SetSuccess(v bool) *AddRepositoryMemberResponseBody
	GetSuccess() *bool
}

type AddRepositoryMemberResponseBody struct {
	// example:
	//
	// Invalid.IdNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 4D6AF7CC-B43B-5454-86AB-023D25E44868
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*AddRepositoryMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddRepositoryMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddRepositoryMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddRepositoryMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRepositoryMemberResponseBody) GetResult() []*AddRepositoryMemberResponseBodyResult {
	return s.Result
}

func (s *AddRepositoryMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRepositoryMemberResponseBody) SetErrorCode(v string) *AddRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetErrorMessage(v string) *AddRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetRequestId(v string) *AddRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetResult(v []*AddRepositoryMemberResponseBodyResult) *AddRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetSuccess(v bool) *AddRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) Validate() error {
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

type AddRepositoryMemberResponseBodyResult struct {
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
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 123456
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s AddRepositoryMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *AddRepositoryMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *AddRepositoryMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *AddRepositoryMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *AddRepositoryMemberResponseBodyResult) GetState() *string {
	return s.State
}

func (s *AddRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *AddRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *AddRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetEmail(v string) *AddRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetId(v int64) *AddRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetState(v string) *AddRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
