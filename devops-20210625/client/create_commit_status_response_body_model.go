// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateCommitStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateCommitStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateCommitStatusResponseBody
	GetRequestId() *string
	SetResult(v *CreateCommitStatusResponseBodyResult) *CreateCommitStatusResponseBody
	GetResult() *CreateCommitStatusResponseBodyResult
	SetSuccess(v bool) *CreateCommitStatusResponseBody
	GetSuccess() *bool
}

type CreateCommitStatusResponseBody struct {
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
	// 020F71F3-F063-5B8B-8978-2B01833216BB
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateCommitStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCommitStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommitStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCommitStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateCommitStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommitStatusResponseBody) GetResult() *CreateCommitStatusResponseBodyResult {
	return s.Result
}

func (s *CreateCommitStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCommitStatusResponseBody) SetErrorCode(v string) *CreateCommitStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCommitStatusResponseBody) SetErrorMessage(v string) *CreateCommitStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCommitStatusResponseBody) SetRequestId(v string) *CreateCommitStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommitStatusResponseBody) SetResult(v *CreateCommitStatusResponseBodyResult) *CreateCommitStatusResponseBody {
	s.Result = v
	return s
}

func (s *CreateCommitStatusResponseBody) SetSuccess(v bool) *CreateCommitStatusResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCommitStatusResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCommitStatusResponseBodyResult struct {
	// example:
	//
	// default
	Context     *string                                      `json:"context,omitempty" xml:"context,omitempty"`
	Creator     *CreateCommitStatusResponseBodyResultCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 30815
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// c0ca103441b9fa7f369be8d24a6f20db44dfddf7
	Sha *string `json:"sha,omitempty" xml:"sha,omitempty"`
	// example:
	//
	// success
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// xxx
	TargetUrl *string `json:"targetUrl,omitempty" xml:"targetUrl,omitempty"`
}

func (s CreateCommitStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCommitStatusResponseBodyResult) GetContext() *string {
	return s.Context
}

func (s *CreateCommitStatusResponseBodyResult) GetCreator() *CreateCommitStatusResponseBodyResultCreator {
	return s.Creator
}

func (s *CreateCommitStatusResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateCommitStatusResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateCommitStatusResponseBodyResult) GetSha() *string {
	return s.Sha
}

func (s *CreateCommitStatusResponseBodyResult) GetState() *string {
	return s.State
}

func (s *CreateCommitStatusResponseBodyResult) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateCommitStatusResponseBodyResult) SetContext(v string) *CreateCommitStatusResponseBodyResult {
	s.Context = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetCreator(v *CreateCommitStatusResponseBodyResultCreator) *CreateCommitStatusResponseBodyResult {
	s.Creator = v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetDescription(v string) *CreateCommitStatusResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetId(v int64) *CreateCommitStatusResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetSha(v string) *CreateCommitStatusResponseBodyResult {
	s.Sha = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetState(v string) *CreateCommitStatusResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) SetTargetUrl(v string) *CreateCommitStatusResponseBodyResult {
	s.TargetUrl = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResult) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCommitStatusResponseBodyResultCreator struct {
	// example:
	//
	// 204485087002425236
	AliyunPk *int64 `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// codeup
	Login *string `json:"login,omitempty" xml:"login,omitempty"`
	// example:
	//
	// User
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCommitStatusResponseBodyResultCreator) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitStatusResponseBodyResultCreator) GoString() string {
	return s.String()
}

func (s *CreateCommitStatusResponseBodyResultCreator) GetAliyunPk() *int64 {
	return s.AliyunPk
}

func (s *CreateCommitStatusResponseBodyResultCreator) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateCommitStatusResponseBodyResultCreator) GetLogin() *string {
	return s.Login
}

func (s *CreateCommitStatusResponseBodyResultCreator) GetType() *string {
	return s.Type
}

func (s *CreateCommitStatusResponseBodyResultCreator) SetAliyunPk(v int64) *CreateCommitStatusResponseBodyResultCreator {
	s.AliyunPk = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResultCreator) SetAvatarUrl(v string) *CreateCommitStatusResponseBodyResultCreator {
	s.AvatarUrl = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResultCreator) SetLogin(v string) *CreateCommitStatusResponseBodyResultCreator {
	s.Login = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResultCreator) SetType(v string) *CreateCommitStatusResponseBodyResultCreator {
	s.Type = &v
	return s
}

func (s *CreateCommitStatusResponseBodyResultCreator) Validate() error {
	return dara.Validate(s)
}
