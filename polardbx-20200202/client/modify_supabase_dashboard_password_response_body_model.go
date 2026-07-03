// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseDashboardPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) *ModifySupabaseDashboardPasswordResponseBody
	GetAccessDeniedDetail() *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail
	SetData(v *ModifySupabaseDashboardPasswordResponseBodyData) *ModifySupabaseDashboardPasswordResponseBody
	GetData() *ModifySupabaseDashboardPasswordResponseBodyData
	SetRequestId(v string) *ModifySupabaseDashboardPasswordResponseBody
	GetRequestId() *string
}

type ModifySupabaseDashboardPasswordResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *ModifySupabaseDashboardPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseDashboardPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseDashboardPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseDashboardPasswordResponseBody) GetAccessDeniedDetail() *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifySupabaseDashboardPasswordResponseBody) GetData() *ModifySupabaseDashboardPasswordResponseBodyData {
	return s.Data
}

func (s *ModifySupabaseDashboardPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseDashboardPasswordResponseBody) SetAccessDeniedDetail(v *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) *ModifySupabaseDashboardPasswordResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBody) SetData(v *ModifySupabaseDashboardPasswordResponseBodyData) *ModifySupabaseDashboardPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBody) SetRequestId(v string) *ModifySupabaseDashboardPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBody) Validate() error {
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

type ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifySupabaseDashboardPasswordResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s ModifySupabaseDashboardPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseDashboardPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySupabaseDashboardPasswordResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySupabaseDashboardPasswordResponseBodyData) SetDBInstanceName(v string) *ModifySupabaseDashboardPasswordResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySupabaseDashboardPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
