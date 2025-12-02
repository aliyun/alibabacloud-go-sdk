// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) *DescribeHistoryTasksStatResponseBody
	GetAccessDeniedDetail() *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail
	SetCode(v string) *DescribeHistoryTasksStatResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeHistoryTasksStatResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*DescribeHistoryTasksStatResponseBodyItems) *DescribeHistoryTasksStatResponseBody
	GetItems() []*DescribeHistoryTasksStatResponseBodyItems
	SetMessage(v string) *DescribeHistoryTasksStatResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHistoryTasksStatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHistoryTasksStatResponseBody
	GetSuccess() *bool
}

type DescribeHistoryTasksStatResponseBody struct {
	// The queried information about the request denial.
	AccessDeniedDetail *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// API status or POP error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The tasks.
	Items []*DescribeHistoryTasksStatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, an success message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHistoryTasksStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponseBody) GetAccessDeniedDetail() *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeHistoryTasksStatResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHistoryTasksStatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeHistoryTasksStatResponseBody) GetItems() []*DescribeHistoryTasksStatResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryTasksStatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHistoryTasksStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryTasksStatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHistoryTasksStatResponseBody) SetAccessDeniedDetail(v *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) *DescribeHistoryTasksStatResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetCode(v string) *DescribeHistoryTasksStatResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetHttpStatusCode(v int32) *DescribeHistoryTasksStatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetItems(v []*DescribeHistoryTasksStatResponseBodyItems) *DescribeHistoryTasksStatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetMessage(v string) *DescribeHistoryTasksStatResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetRequestId(v string) *DescribeHistoryTasksStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) SetSuccess(v bool) *DescribeHistoryTasksStatResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHistoryTasksStatResponseBodyAccessDeniedDetail struct {
	// Authorized action
	//
	// example:
	//
	// test
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the requester.
	//
	// example:
	//
	// test
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account to which the authentication principal belongs.
	//
	// example:
	//
	// 141345906006****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The requester type.
	//
	// example:
	//
	// test
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// test
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the policy denial. Valid values:
	//
	// 	- **ImplicitDeny**: The resource holder has not configured a policy for the current user. By default, unauthorized operations are denied.
	//
	// 	- **ExplicitDeny**: The RAM policy configured by the resource holder explicitly denies the current user access to the corresponding resources.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The type of the policy that causes the access denied error.
	//
	// 	- **ControlPolicy**: control policy
	//
	// 	- **SessionPolicy**: an additional policy that is added to the temporary token.
	//
	// 	- **AssumeRolePolicy**: the authorization policy of the RAM role.
	//
	// 	- **AccountLevelIdentityBasedPolicy**: the principal policies of the account authorization scope, including custom policies and system policies.
	//
	// 	- **ResourceGroupLevelIdentityBasedPolicy**: the principal policy of the resource group authorization scope, including custom policies and system policies.
	//
	// example:
	//
	// ControlPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeHistoryTasksStatResponseBodyItems struct {
	// The status of the APS job.
	//
	// 	- **Scheduled**
	//
	// 	- **Running**
	//
	// 	- **Succeed**
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Cancelling**
	//
	// 	- **Canceled**
	//
	// 	- **Waiting**
	//
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryTasksStatResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksStatResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryTasksStatResponseBodyItems) SetStatus(v string) *DescribeHistoryTasksStatResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyItems) SetTotalCount(v int32) *DescribeHistoryTasksStatResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryTasksStatResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
