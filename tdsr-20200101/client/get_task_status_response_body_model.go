// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetTaskStatusResponseBodyAccessDeniedDetail) *GetTaskStatusResponseBody
	GetAccessDeniedDetail() *GetTaskStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetTaskStatusResponseBody
	GetCode() *int64
	SetErrorCode(v string) *GetTaskStatusResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetTaskStatusResponseBody
	GetErrorMsg() *string
	SetMessage(v string) *GetTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetTaskStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetTaskStatusResponseBody
	GetSuccess() *bool
	SetType(v string) *GetTaskStatusResponseBody
	GetType() *string
}

type GetTaskStatusResponseBody struct {
	AccessDeniedDetail *GetTaskStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// xxx
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
	// init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// wallline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) GetAccessDeniedDetail() *GetTaskStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetTaskStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetTaskStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskStatusResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskStatusResponseBody) GetType() *string {
	return s.Type
}

func (s *GetTaskStatusResponseBody) SetAccessDeniedDetail(v *GetTaskStatusResponseBodyAccessDeniedDetail) *GetTaskStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetTaskStatusResponseBody) SetCode(v int64) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrorCode(v string) *GetTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrorMsg(v string) *GetTaskStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetStatus(v string) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetType(v string) *GetTaskStatusResponseBody {
	s.Type = &v
	return s
}

func (s *GetTaskStatusResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetTaskStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetTaskStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetTaskStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
