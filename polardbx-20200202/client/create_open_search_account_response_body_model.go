// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) *CreateOpenSearchAccountResponseBody
	GetAccessDeniedDetail() *CreateOpenSearchAccountResponseBodyAccessDeniedDetail
	SetData(v *CreateOpenSearchAccountResponseBodyData) *CreateOpenSearchAccountResponseBody
	GetData() *CreateOpenSearchAccountResponseBodyData
	SetRequestId(v string) *CreateOpenSearchAccountResponseBody
	GetRequestId() *string
}

type CreateOpenSearchAccountResponseBody struct {
	AccessDeniedDetail *CreateOpenSearchAccountResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *CreateOpenSearchAccountResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 9B2F3840-****-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpenSearchAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchAccountResponseBody) GetAccessDeniedDetail() *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateOpenSearchAccountResponseBody) GetData() *CreateOpenSearchAccountResponseBodyData {
	return s.Data
}

func (s *CreateOpenSearchAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSearchAccountResponseBody) SetAccessDeniedDetail(v *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) *CreateOpenSearchAccountResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateOpenSearchAccountResponseBody) SetData(v *CreateOpenSearchAccountResponseBodyData) *CreateOpenSearchAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateOpenSearchAccountResponseBody) SetRequestId(v string) *CreateOpenSearchAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBody) Validate() error {
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

type CreateOpenSearchAccountResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateOpenSearchAccountResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateOpenSearchAccountResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateOpenSearchAccountResponseBodyData struct {
	// example:
	//
	// polardbx_meta_ro
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateOpenSearchAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateOpenSearchAccountResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateOpenSearchAccountResponseBodyData) SetAccountName(v string) *CreateOpenSearchAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyData) SetTaskId(v int32) *CreateOpenSearchAccountResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateOpenSearchAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
