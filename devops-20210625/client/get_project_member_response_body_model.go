// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetProjectMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetProjectMemberResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetProjectMemberResponseBody
	GetRequestId() *string
	SetResult(v *GetProjectMemberResponseBodyResult) *GetProjectMemberResponseBody
	GetResult() *GetProjectMemberResponseBodyResult
	SetSuccess(v bool) *GetProjectMemberResponseBody
	GetSuccess() *bool
}

type GetProjectMemberResponseBody struct {
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
	// 30C99C69-A340-5E2E-ACE4-8888FF50CF52
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetProjectMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProjectMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProjectMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetProjectMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectMemberResponseBody) GetResult() *GetProjectMemberResponseBodyResult {
	return s.Result
}

func (s *GetProjectMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectMemberResponseBody) SetErrorCode(v string) *GetProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetErrorMessage(v string) *GetProjectMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetResult(v *GetProjectMemberResponseBodyResult) *GetProjectMemberResponseBody {
	s.Result = v
	return s
}

func (s *GetProjectMemberResponseBody) SetSuccess(v bool) *GetProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectMemberResponseBodyResult struct {
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c4ef67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 2959
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetProjectMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyResult) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *GetProjectMemberResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetProjectMemberResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *GetProjectMemberResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetProjectMemberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetProjectMemberResponseBodyResult) SetAccessLevel(v int32) *GetProjectMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetAvatarUrl(v string) *GetProjectMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetEmail(v string) *GetProjectMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetId(v int64) *GetProjectMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetName(v string) *GetProjectMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
