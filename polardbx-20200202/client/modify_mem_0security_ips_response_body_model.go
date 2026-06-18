// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMem0SecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) *ModifyMem0SecurityIpsResponseBody
	GetAccessDeniedDetail() *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *ModifyMem0SecurityIpsResponseBodyData) *ModifyMem0SecurityIpsResponseBody
	GetData() *ModifyMem0SecurityIpsResponseBodyData
	SetRequestId(v string) *ModifyMem0SecurityIpsResponseBody
	GetRequestId() *string
}

type ModifyMem0SecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *ModifyMem0SecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMem0SecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMem0SecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMem0SecurityIpsResponseBody) GetAccessDeniedDetail() *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyMem0SecurityIpsResponseBody) GetData() *ModifyMem0SecurityIpsResponseBodyData {
	return s.Data
}

func (s *ModifyMem0SecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMem0SecurityIpsResponseBody) SetAccessDeniedDetail(v *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) *ModifyMem0SecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBody) SetData(v *ModifyMem0SecurityIpsResponseBodyData) *ModifyMem0SecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBody) SetRequestId(v string) *ModifyMem0SecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBody) Validate() error {
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

type ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail struct {
	// The operation name.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the identity used for authentication in the request.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account to which the authenticated principal belongs.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the identity used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user
	//
	// - AssumedRoleUser: RAM role
	//
	// - Federated: SSO federated identity.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted complete diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of denial by the access policy. Valid values:
	//
	// - **ImplicitDeny**: The resource owner has not configured a relevant permission policy for the current user. Unauthorized operations are denied by default.
	//
	// - **ExplicitDeny**: The RAM policy configured by the resource owner explicitly denies the current user access to the corresponding resource.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyMem0SecurityIpsResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyMem0SecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyMem0SecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyMem0SecurityIpsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyMem0SecurityIpsResponseBodyData) SetTaskId(v int32) *ModifyMem0SecurityIpsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyMem0SecurityIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
