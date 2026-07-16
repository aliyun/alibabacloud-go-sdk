// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteServiceAccountResponseBodyAccessDeniedDetail) *DeleteServiceAccountResponseBody
	GetAccessDeniedDetail() *DeleteServiceAccountResponseBodyAccessDeniedDetail
	SetData(v *DeleteServiceAccountResponseBodyData) *DeleteServiceAccountResponseBody
	GetData() *DeleteServiceAccountResponseBodyData
	SetMessage(v string) *DeleteServiceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteServiceAccountResponseBody
	GetSuccess() *bool
}

type DeleteServiceAccountResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeleteServiceAccountResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DeleteServiceAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is empty if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServiceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceAccountResponseBody) GetAccessDeniedDetail() *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteServiceAccountResponseBody) GetData() *DeleteServiceAccountResponseBodyData {
	return s.Data
}

func (s *DeleteServiceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteServiceAccountResponseBody) SetAccessDeniedDetail(v *DeleteServiceAccountResponseBodyAccessDeniedDetail) *DeleteServiceAccountResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteServiceAccountResponseBody) SetData(v *DeleteServiceAccountResponseBodyData) *DeleteServiceAccountResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServiceAccountResponseBody) SetMessage(v string) *DeleteServiceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceAccountResponseBody) SetRequestId(v string) *DeleteServiceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceAccountResponseBody) SetSuccess(v bool) *DeleteServiceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteServiceAccountResponseBody) Validate() error {
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

type DeleteServiceAccountResponseBodyAccessDeniedDetail struct {
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
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of missing permission.
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

func (s DeleteServiceAccountResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAccountResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteServiceAccountResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteServiceAccountResponseBodyData struct {
	// The account name.
	//
	// example:
	//
	// polardbx_meta_ro
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the import task is successful.
	//
	// example:
	//
	// fase
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E5DCFFC-A00D-****-836E-73318F8CA577
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service account type.
	//
	// example:
	//
	// METADATA_READONLY
	ServiceAccountType *string `json:"ServiceAccountType,omitempty" xml:"ServiceAccountType,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteServiceAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteServiceAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteServiceAccountResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteServiceAccountResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceAccountResponseBodyData) GetServiceAccountType() *string {
	return s.ServiceAccountType
}

func (s *DeleteServiceAccountResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteServiceAccountResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteServiceAccountResponseBodyData) SetAccountName(v string) *DeleteServiceAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) SetDeleted(v bool) *DeleteServiceAccountResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) SetRequestId(v string) *DeleteServiceAccountResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) SetServiceAccountType(v string) *DeleteServiceAccountResponseBodyData {
	s.ServiceAccountType = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) SetStatus(v string) *DeleteServiceAccountResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) SetTaskId(v int32) *DeleteServiceAccountResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteServiceAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
