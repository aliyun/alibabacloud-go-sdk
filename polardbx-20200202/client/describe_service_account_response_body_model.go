// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeServiceAccountResponseBodyAccessDeniedDetail) *DescribeServiceAccountResponseBody
	GetAccessDeniedDetail() *DescribeServiceAccountResponseBodyAccessDeniedDetail
	SetData(v *DescribeServiceAccountResponseBodyData) *DescribeServiceAccountResponseBody
	GetData() *DescribeServiceAccountResponseBodyData
	SetMessage(v string) *DescribeServiceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeServiceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeServiceAccountResponseBody
	GetSuccess() *bool
}

type DescribeServiceAccountResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeServiceAccountResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The instance details.
	Data *DescribeServiceAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. "success" is returned if the request was successful. Otherwise, the corresponding error code is returned.
	//
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountResponseBody) GetAccessDeniedDetail() *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeServiceAccountResponseBody) GetData() *DescribeServiceAccountResponseBodyData {
	return s.Data
}

func (s *DescribeServiceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeServiceAccountResponseBody) SetAccessDeniedDetail(v *DescribeServiceAccountResponseBodyAccessDeniedDetail) *DescribeServiceAccountResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeServiceAccountResponseBody) SetData(v *DescribeServiceAccountResponseBodyData) *DescribeServiceAccountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeServiceAccountResponseBody) SetMessage(v string) *DescribeServiceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceAccountResponseBody) SetRequestId(v string) *DescribeServiceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAccountResponseBody) SetSuccess(v bool) *DescribeServiceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeServiceAccountResponseBody) Validate() error {
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

type DescribeServiceAccountResponseBodyAccessDeniedDetail struct {
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
	// The type of the permission denial.
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

func (s DescribeServiceAccountResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeServiceAccountResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceAccountResponseBodyData struct {
	// The service account in the list.
	ServiceAccounts []*DescribeServiceAccountResponseBodyDataServiceAccounts `json:"ServiceAccounts,omitempty" xml:"ServiceAccounts,omitempty" type:"Repeated"`
}

func (s DescribeServiceAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountResponseBodyData) GetServiceAccounts() []*DescribeServiceAccountResponseBodyDataServiceAccounts {
	return s.ServiceAccounts
}

func (s *DescribeServiceAccountResponseBodyData) SetServiceAccounts(v []*DescribeServiceAccountResponseBodyDataServiceAccounts) *DescribeServiceAccountResponseBodyData {
	s.ServiceAccounts = v
	return s
}

func (s *DescribeServiceAccountResponseBodyData) Validate() error {
	if s.ServiceAccounts != nil {
		for _, item := range s.ServiceAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceAccountResponseBodyDataServiceAccounts struct {
	// The account name.
	//
	// example:
	//
	// polardbx_meta_ro
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The time when the account was created.
	//
	// example:
	//
	// 2023-07-21T14:17:25+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The service account type.
	//
	// example:
	//
	// METADATA_READONLY
	ServiceAccountType *string `json:"ServiceAccountType,omitempty" xml:"ServiceAccountType,omitempty"`
	// The status of the backup set. Valid values:
	//
	// - **0**: Backing up.
	//
	// - **1**: Backup succeeded.
	//
	// - **2**: Backup failed.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceAccountResponseBodyDataServiceAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountResponseBodyDataServiceAccounts) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) GetServiceAccountType() *string {
	return s.ServiceAccountType
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) GetStatus() *string {
	return s.Status
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) SetAccountName(v string) *DescribeServiceAccountResponseBodyDataServiceAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) SetCreatedTime(v string) *DescribeServiceAccountResponseBodyDataServiceAccounts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) SetServiceAccountType(v string) *DescribeServiceAccountResponseBodyDataServiceAccounts {
	s.ServiceAccountType = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) SetStatus(v string) *DescribeServiceAccountResponseBodyDataServiceAccounts {
	s.Status = &v
	return s
}

func (s *DescribeServiceAccountResponseBodyDataServiceAccounts) Validate() error {
	return dara.Validate(s)
}
