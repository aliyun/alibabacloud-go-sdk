// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchAccountInfoResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchAccountInfoResponseBodyData) *DescribeOpenSearchAccountInfoResponseBody
	GetData() *DescribeOpenSearchAccountInfoResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchAccountInfoResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchAccountInfoResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DescribeOpenSearchAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchAccountInfoResponseBody) GetData() *DescribeOpenSearchAccountInfoResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchAccountInfoResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchAccountInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBody) SetData(v *DescribeOpenSearchAccountInfoResponseBodyData) *DescribeOpenSearchAccountInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBody) SetRequestId(v string) *DescribeOpenSearchAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBody) Validate() error {
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

type DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail struct {
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

func (s DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchAccountInfoResponseBodyData struct {
	// The list of accounts.
	Accounts []*DescribeOpenSearchAccountInfoResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The time when the password was last modified.
	//
	// example:
	//
	// 2026-08-21T12:00:00Z
	PasswordLastModified *string `json:"PasswordLastModified,omitempty" xml:"PasswordLastModified,omitempty"`
	// The account name of the OpenSearch instance.
	//
	// example:
	//
	// elastic
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeOpenSearchAccountInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) GetAccounts() []*DescribeOpenSearchAccountInfoResponseBodyDataAccounts {
	return s.Accounts
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) GetPasswordLastModified() *string {
	return s.PasswordLastModified
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) SetAccounts(v []*DescribeOpenSearchAccountInfoResponseBodyDataAccounts) *DescribeOpenSearchAccountInfoResponseBodyData {
	s.Accounts = v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) SetPasswordLastModified(v string) *DescribeOpenSearchAccountInfoResponseBodyData {
	s.PasswordLastModified = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) SetUsername(v string) *DescribeOpenSearchAccountInfoResponseBodyData {
	s.Username = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyData) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOpenSearchAccountInfoResponseBodyDataAccounts struct {
	// The account status. Valid values:
	//
	// - **Creating**: The account is being created.
	//
	// - **Available**: The account is available.
	//
	// - **Deleting**: The account is being deleted.
	//
	// example:
	//
	// 1
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The account type.
	//
	// - Before three-role mode is enabled: 0 indicates a standard account, and 1 indicates a privileged account.
	//
	// - After three-role mode is enabled: 0 indicates a standard account, 2 indicates a system administrator account, 3 indicates a security administrator account, and 4 indicates an audit administrator account.
	//
	// example:
	//
	// 2,3,4
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The account name of the OpenSearch instance.
	//
	// example:
	//
	// elastic
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeOpenSearchAccountInfoResponseBodyDataAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) GetUsername() *string {
	return s.Username
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) SetAccountStatus(v string) *DescribeOpenSearchAccountInfoResponseBodyDataAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) SetAccountType(v string) *DescribeOpenSearchAccountInfoResponseBodyDataAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) SetUsername(v string) *DescribeOpenSearchAccountInfoResponseBodyDataAccounts {
	s.Username = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponseBodyDataAccounts) Validate() error {
	return dara.Validate(s)
}
