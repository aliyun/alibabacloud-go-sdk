// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) *DeleteOpenSearchAccountResponseBody
	GetAccessDeniedDetail() *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail
	SetData(v *DeleteOpenSearchAccountResponseBodyData) *DeleteOpenSearchAccountResponseBody
	GetData() *DeleteOpenSearchAccountResponseBodyData
	SetRequestId(v string) *DeleteOpenSearchAccountResponseBody
	GetRequestId() *string
}

type DeleteOpenSearchAccountResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DeleteOpenSearchAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOpenSearchAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchAccountResponseBody) GetAccessDeniedDetail() *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteOpenSearchAccountResponseBody) GetData() *DeleteOpenSearchAccountResponseBodyData {
	return s.Data
}

func (s *DeleteOpenSearchAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpenSearchAccountResponseBody) SetAccessDeniedDetail(v *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) *DeleteOpenSearchAccountResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteOpenSearchAccountResponseBody) SetData(v *DeleteOpenSearchAccountResponseBodyData) *DeleteOpenSearchAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteOpenSearchAccountResponseBody) SetRequestId(v string) *DeleteOpenSearchAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteOpenSearchAccountResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteOpenSearchAccountResponseBodyData struct {
	// The account name.
	//
	// example:
	//
	// polardbx_meta_ro
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteOpenSearchAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteOpenSearchAccountResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteOpenSearchAccountResponseBodyData) SetAccountName(v string) *DeleteOpenSearchAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyData) SetTaskId(v int32) *DeleteOpenSearchAccountResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteOpenSearchAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
