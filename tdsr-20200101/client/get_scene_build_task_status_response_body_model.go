// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneBuildTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) *GetSceneBuildTaskStatusResponseBody
	GetAccessDeniedDetail() *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetSceneBuildTaskStatusResponseBody
	GetCode() *int64
	SetErrorCode(v string) *GetSceneBuildTaskStatusResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetSceneBuildTaskStatusResponseBody
	GetErrorMsg() *string
	SetMessage(v string) *GetSceneBuildTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSceneBuildTaskStatusResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetSceneBuildTaskStatusResponseBody
	GetSceneId() *string
	SetStatus(v string) *GetSceneBuildTaskStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetSceneBuildTaskStatusResponseBody
	GetSuccess() *bool
}

type GetSceneBuildTaskStatusResponseBody struct {
	AccessDeniedDetail *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// m+0cmndEGjg9pv/hy4jh****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSceneBuildTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponseBody) GetAccessDeniedDetail() *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetSceneBuildTaskStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSceneBuildTaskStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSceneBuildTaskStatusResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetSceneBuildTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSceneBuildTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSceneBuildTaskStatusResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSceneBuildTaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSceneBuildTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSceneBuildTaskStatusResponseBody) SetAccessDeniedDetail(v *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) *GetSceneBuildTaskStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetCode(v int64) *GetSceneBuildTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetErrorCode(v string) *GetSceneBuildTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetErrorMsg(v string) *GetSceneBuildTaskStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetMessage(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetRequestId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetSceneId(v string) *GetSceneBuildTaskStatusResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetStatus(v string) *GetSceneBuildTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) SetSuccess(v bool) *GetSceneBuildTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetSceneBuildTaskStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
