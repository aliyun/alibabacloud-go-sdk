// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateServiceAccountResponseBodyAccessDeniedDetail) *CreateServiceAccountResponseBody
	GetAccessDeniedDetail() *CreateServiceAccountResponseBodyAccessDeniedDetail
	SetData(v *CreateServiceAccountResponseBodyData) *CreateServiceAccountResponseBody
	GetData() *CreateServiceAccountResponseBodyData
	SetMessage(v string) *CreateServiceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceAccountResponseBody
	GetSuccess() *bool
}

type CreateServiceAccountResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateServiceAccountResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned policy details.
	Data *CreateServiceAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information returned. If the request is successful, **success*	- is returned. If the request fails, the corresponding error code is returned.
	//
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAccountResponseBody) GetAccessDeniedDetail() *CreateServiceAccountResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateServiceAccountResponseBody) GetData() *CreateServiceAccountResponseBodyData {
	return s.Data
}

func (s *CreateServiceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceAccountResponseBody) SetAccessDeniedDetail(v *CreateServiceAccountResponseBodyAccessDeniedDetail) *CreateServiceAccountResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateServiceAccountResponseBody) SetData(v *CreateServiceAccountResponseBodyData) *CreateServiceAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceAccountResponseBody) SetMessage(v string) *CreateServiceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceAccountResponseBody) SetRequestId(v string) *CreateServiceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceAccountResponseBody) SetSuccess(v bool) *CreateServiceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceAccountResponseBody) Validate() error {
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

type CreateServiceAccountResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The authentication principal type.
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
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateServiceAccountResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAccountResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateServiceAccountResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateServiceAccountResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateServiceAccountResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateServiceAccountResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateServiceAccountResponseBodyData struct {
	// The account name.
	//
	// example:
	//
	// polardbx_meta_ro
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the account already exists.
	//
	// example:
	//
	// False
	AlreadyExists *bool `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E5DCFFC-A00D-****-836E-73318F8CA577
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the service account.
	//
	// example:
	//
	// 服务账号类型
	ServiceAccountType *string `json:"ServiceAccountType,omitempty" xml:"ServiceAccountType,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateServiceAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateServiceAccountResponseBodyData) GetAlreadyExists() *bool {
	return s.AlreadyExists
}

func (s *CreateServiceAccountResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceAccountResponseBodyData) GetServiceAccountType() *string {
	return s.ServiceAccountType
}

func (s *CreateServiceAccountResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceAccountResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateServiceAccountResponseBodyData) SetAccountName(v string) *CreateServiceAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) SetAlreadyExists(v bool) *CreateServiceAccountResponseBodyData {
	s.AlreadyExists = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) SetRequestId(v string) *CreateServiceAccountResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) SetServiceAccountType(v string) *CreateServiceAccountResponseBodyData {
	s.ServiceAccountType = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) SetStatus(v string) *CreateServiceAccountResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) SetTaskId(v int32) *CreateServiceAccountResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateServiceAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
