// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseSecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) *ModifySupabaseSecurityIPListResponseBody
	GetAccessDeniedDetail() *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail
	SetData(v *ModifySupabaseSecurityIPListResponseBodyData) *ModifySupabaseSecurityIPListResponseBody
	GetData() *ModifySupabaseSecurityIPListResponseBodyData
	SetRequestId(v string) *ModifySupabaseSecurityIPListResponseBody
	GetRequestId() *string
}

type ModifySupabaseSecurityIPListResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The operation result.
	Data *ModifySupabaseSecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseSecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseSecurityIPListResponseBody) GetAccessDeniedDetail() *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifySupabaseSecurityIPListResponseBody) GetData() *ModifySupabaseSecurityIPListResponseBodyData {
	return s.Data
}

func (s *ModifySupabaseSecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseSecurityIPListResponseBody) SetAccessDeniedDetail(v *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) *ModifySupabaseSecurityIPListResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBody) SetData(v *ModifySupabaseSecurityIPListResponseBodyData) *ModifySupabaseSecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBody) SetRequestId(v string) *ModifySupabaseSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBody) Validate() error {
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

type ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail struct {
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

func (s ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifySupabaseSecurityIPListResponseBodyData struct {
	// The instance name.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The whitelist group name.
	//
	// example:
	//
	// special
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP whitelist.
	//
	// example:
	//
	// 100.104.187.192/26,100.104.225.192/26,100.104.145.64/26
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySupabaseSecurityIPListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseSecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) SetDBInstanceName(v string) *ModifySupabaseSecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) SetGroupName(v string) *ModifySupabaseSecurityIPListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) SetSecurityIPList(v string) *ModifySupabaseSecurityIPListResponseBodyData {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
