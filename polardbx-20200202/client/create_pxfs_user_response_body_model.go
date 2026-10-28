// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePxfsUserResponseBodyAccessDeniedDetail) *CreatePxfsUserResponseBody
	GetAccessDeniedDetail() *CreatePxfsUserResponseBodyAccessDeniedDetail
	SetData(v *CreatePxfsUserResponseBodyData) *CreatePxfsUserResponseBody
	GetData() *CreatePxfsUserResponseBodyData
	SetRequestId(v string) *CreatePxfsUserResponseBody
	GetRequestId() *string
}

type CreatePxfsUserResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePxfsUserResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *CreatePxfsUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePxfsUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePxfsUserResponseBody) GetAccessDeniedDetail() *CreatePxfsUserResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePxfsUserResponseBody) GetData() *CreatePxfsUserResponseBodyData {
	return s.Data
}

func (s *CreatePxfsUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePxfsUserResponseBody) SetAccessDeniedDetail(v *CreatePxfsUserResponseBodyAccessDeniedDetail) *CreatePxfsUserResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePxfsUserResponseBody) SetData(v *CreatePxfsUserResponseBodyData) *CreatePxfsUserResponseBody {
	s.Data = v
	return s
}

func (s *CreatePxfsUserResponseBody) SetRequestId(v string) *CreatePxfsUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePxfsUserResponseBody) Validate() error {
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

type CreatePxfsUserResponseBodyAccessDeniedDetail struct {
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

func (s CreatePxfsUserResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsUserResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePxfsUserResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePxfsUserResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePxfsUserResponseBodyData struct {
	// The identity name.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// POSIX GID
	//
	// example:
	//
	// 1000
	PosixGid *int64 `json:"PosixGid,omitempty" xml:"PosixGid,omitempty"`
	// POSIX UID
	//
	// example:
	//
	// 1000
	PosixUid *int64 `json:"PosixUid,omitempty" xml:"PosixUid,omitempty"`
	// The identity role.
	//
	// example:
	//
	// MEMBER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreatePxfsUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePxfsUserResponseBodyData) GetIdentityName() *string {
	return s.IdentityName
}

func (s *CreatePxfsUserResponseBodyData) GetPosixGid() *int64 {
	return s.PosixGid
}

func (s *CreatePxfsUserResponseBodyData) GetPosixUid() *int64 {
	return s.PosixUid
}

func (s *CreatePxfsUserResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *CreatePxfsUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreatePxfsUserResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsUserResponseBodyData) SetIdentityName(v string) *CreatePxfsUserResponseBodyData {
	s.IdentityName = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) SetPosixGid(v int64) *CreatePxfsUserResponseBodyData {
	s.PosixGid = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) SetPosixUid(v int64) *CreatePxfsUserResponseBodyData {
	s.PosixUid = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) SetRole(v string) *CreatePxfsUserResponseBodyData {
	s.Role = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) SetStatus(v string) *CreatePxfsUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) SetTenantId(v string) *CreatePxfsUserResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
