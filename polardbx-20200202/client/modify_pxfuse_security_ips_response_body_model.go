// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfuseSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) *ModifyPxfuseSecurityIpsResponseBody
	GetAccessDeniedDetail() *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail
	SetData(v *ModifyPxfuseSecurityIpsResponseBodyData) *ModifyPxfuseSecurityIpsResponseBody
	GetData() *ModifyPxfuseSecurityIpsResponseBodyData
	SetRequestId(v string) *ModifyPxfuseSecurityIpsResponseBody
	GetRequestId() *string
}

type ModifyPxfuseSecurityIpsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *ModifyPxfuseSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPxfuseSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfuseSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPxfuseSecurityIpsResponseBody) GetAccessDeniedDetail() *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyPxfuseSecurityIpsResponseBody) GetData() *ModifyPxfuseSecurityIpsResponseBodyData {
	return s.Data
}

func (s *ModifyPxfuseSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPxfuseSecurityIpsResponseBody) SetAccessDeniedDetail(v *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) *ModifyPxfuseSecurityIpsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBody) SetData(v *ModifyPxfuseSecurityIpsResponseBodyData) *ModifyPxfuseSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBody) SetRequestId(v string) *ModifyPxfuseSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBody) Validate() error {
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

type ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail struct {
	// The name of the operation.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authenticated principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authenticated principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the identity used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user.
	//
	// - AssumedRoleUser: RAM role.
	//
	// - Federated: SSO federated identity.
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
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyPxfuseSecurityIpsResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyPxfuseSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfuseSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyPxfuseSecurityIpsResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyPxfuseSecurityIpsResponseBodyData) SetTaskId(v int32) *ModifyPxfuseSecurityIpsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
