// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommitStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListCommitStatusesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListCommitStatusesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListCommitStatusesResponseBody
	GetRequestId() *string
	SetResult(v []*ListCommitStatusesResponseBodyResult) *ListCommitStatusesResponseBody
	GetResult() []*ListCommitStatusesResponseBodyResult
	SetSuccess(v bool) *ListCommitStatusesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListCommitStatusesResponseBody
	GetTotal() *int64
}

type ListCommitStatusesResponseBody struct {
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
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListCommitStatusesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 28
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCommitStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommitStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommitStatusesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCommitStatusesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCommitStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommitStatusesResponseBody) GetResult() []*ListCommitStatusesResponseBodyResult {
	return s.Result
}

func (s *ListCommitStatusesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCommitStatusesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCommitStatusesResponseBody) SetErrorCode(v string) *ListCommitStatusesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCommitStatusesResponseBody) SetErrorMessage(v string) *ListCommitStatusesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCommitStatusesResponseBody) SetRequestId(v string) *ListCommitStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommitStatusesResponseBody) SetResult(v []*ListCommitStatusesResponseBodyResult) *ListCommitStatusesResponseBody {
	s.Result = v
	return s
}

func (s *ListCommitStatusesResponseBody) SetSuccess(v bool) *ListCommitStatusesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCommitStatusesResponseBody) SetTotal(v int64) *ListCommitStatusesResponseBody {
	s.Total = &v
	return s
}

func (s *ListCommitStatusesResponseBody) Validate() error {
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

type ListCommitStatusesResponseBodyResult struct {
	// example:
	//
	// test-commit-status-context
	Context     *string                                      `json:"context,omitempty" xml:"context,omitempty"`
	Creator     *ListCommitStatusesResponseBodyResultCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 19285
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 61cc69557962d29f737a91730b3e86f497f083a3
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

func (s ListCommitStatusesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCommitStatusesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommitStatusesResponseBodyResult) GetContext() *string {
	return s.Context
}

func (s *ListCommitStatusesResponseBodyResult) GetCreator() *ListCommitStatusesResponseBodyResultCreator {
	return s.Creator
}

func (s *ListCommitStatusesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListCommitStatusesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListCommitStatusesResponseBodyResult) GetSha() *string {
	return s.Sha
}

func (s *ListCommitStatusesResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListCommitStatusesResponseBodyResult) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *ListCommitStatusesResponseBodyResult) SetContext(v string) *ListCommitStatusesResponseBodyResult {
	s.Context = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetCreator(v *ListCommitStatusesResponseBodyResultCreator) *ListCommitStatusesResponseBodyResult {
	s.Creator = v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetDescription(v string) *ListCommitStatusesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetId(v int64) *ListCommitStatusesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetSha(v string) *ListCommitStatusesResponseBodyResult {
	s.Sha = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetState(v string) *ListCommitStatusesResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) SetTargetUrl(v string) *ListCommitStatusesResponseBodyResult {
	s.TargetUrl = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResult) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCommitStatusesResponseBodyResultCreator struct {
	// example:
	//
	// 235671547828975455
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
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

func (s ListCommitStatusesResponseBodyResultCreator) String() string {
	return dara.Prettify(s)
}

func (s ListCommitStatusesResponseBodyResultCreator) GoString() string {
	return s.String()
}

func (s *ListCommitStatusesResponseBodyResultCreator) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListCommitStatusesResponseBodyResultCreator) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListCommitStatusesResponseBodyResultCreator) GetLogin() *string {
	return s.Login
}

func (s *ListCommitStatusesResponseBodyResultCreator) GetType() *string {
	return s.Type
}

func (s *ListCommitStatusesResponseBodyResultCreator) SetAliyunPk(v string) *ListCommitStatusesResponseBodyResultCreator {
	s.AliyunPk = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResultCreator) SetAvatarUrl(v string) *ListCommitStatusesResponseBodyResultCreator {
	s.AvatarUrl = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResultCreator) SetLogin(v string) *ListCommitStatusesResponseBodyResultCreator {
	s.Login = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResultCreator) SetType(v string) *ListCommitStatusesResponseBodyResultCreator {
	s.Type = &v
	return s
}

func (s *ListCommitStatusesResponseBodyResultCreator) Validate() error {
	return dara.Validate(s)
}
