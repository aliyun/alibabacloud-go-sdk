// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetUserInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUserInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetUserInfoResponseBody
	GetRequestId() *string
	SetResult(v *GetUserInfoResponseBodyResult) *GetUserInfoResponseBody
	GetResult() *GetUserInfoResponseBodyResult
	SetSuccess(v bool) *GetUserInfoResponseBody
	GetSuccess() *bool
}

type GetUserInfoResponseBody struct {
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
	// F7B85D1B-D1C2-140F-A039-341859F130B9
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetUserInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserInfoResponseBody) GetResult() *GetUserInfoResponseBodyResult {
	return s.Result
}

func (s *GetUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserInfoResponseBody) SetErrorCode(v string) *GetUserInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserInfoResponseBody) SetErrorMessage(v string) *GetUserInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetResult(v *GetUserInfoResponseBodyResult) *GetUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetUserInfoResponseBody) SetSuccess(v bool) *GetUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserInfoResponseBodyResult struct {
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
	// 4205
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetUserInfoResponseBodyResult) GetEmail() *string {
	return s.Email
}

func (s *GetUserInfoResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetUserInfoResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetUserInfoResponseBodyResult) GetUsername() *string {
	return s.Username
}

func (s *GetUserInfoResponseBodyResult) SetAvatarUrl(v string) *GetUserInfoResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetEmail(v string) *GetUserInfoResponseBodyResult {
	s.Email = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetId(v int64) *GetUserInfoResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetName(v string) *GetUserInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetUsername(v string) *GetUserInfoResponseBodyResult {
	s.Username = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
