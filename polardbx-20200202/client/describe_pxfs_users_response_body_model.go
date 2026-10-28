// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfsUsersResponseBodyAccessDeniedDetail) *DescribePxfsUsersResponseBody
	GetAccessDeniedDetail() *DescribePxfsUsersResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfsUsersResponseBodyData) *DescribePxfsUsersResponseBody
	GetData() *DescribePxfsUsersResponseBodyData
	SetRequestId(v string) *DescribePxfsUsersResponseBody
	GetRequestId() *string
}

type DescribePxfsUsersResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribePxfsUsersResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DescribePxfsUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfsUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersResponseBody) GetAccessDeniedDetail() *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfsUsersResponseBody) GetData() *DescribePxfsUsersResponseBodyData {
	return s.Data
}

func (s *DescribePxfsUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfsUsersResponseBody) SetAccessDeniedDetail(v *DescribePxfsUsersResponseBodyAccessDeniedDetail) *DescribePxfsUsersResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfsUsersResponseBody) SetData(v *DescribePxfsUsersResponseBodyData) *DescribePxfsUsersResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfsUsersResponseBody) SetRequestId(v string) *DescribePxfsUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfsUsersResponseBody) Validate() error {
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

type DescribePxfsUsersResponseBodyAccessDeniedDetail struct {
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
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribePxfsUsersResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfsUsersResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfsUsersResponseBodyData struct {
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The list of tenant identities.
	Users []*DescribePxfsUsersResponseBodyDataUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribePxfsUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribePxfsUsersResponseBodyData) GetUsers() []*DescribePxfsUsersResponseBodyDataUsers {
	return s.Users
}

func (s *DescribePxfsUsersResponseBodyData) SetTenantId(v string) *DescribePxfsUsersResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyData) SetUsers(v []*DescribePxfsUsersResponseBodyDataUsers) *DescribePxfsUsersResponseBodyData {
	s.Users = v
	return s
}

func (s *DescribePxfsUsersResponseBodyData) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePxfsUsersResponseBodyDataUsers struct {
	// The creation time.
	//
	// example:
	//
	// 2026-09-03T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
}

func (s DescribePxfsUsersResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetPosixGid() *int64 {
	return s.PosixGid
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetPosixUid() *int64 {
	return s.PosixUid
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetRole() *string {
	return s.Role
}

func (s *DescribePxfsUsersResponseBodyDataUsers) GetStatus() *string {
	return s.Status
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetCreateTime(v string) *DescribePxfsUsersResponseBodyDataUsers {
	s.CreateTime = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetIdentityName(v string) *DescribePxfsUsersResponseBodyDataUsers {
	s.IdentityName = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetPosixGid(v int64) *DescribePxfsUsersResponseBodyDataUsers {
	s.PosixGid = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetPosixUid(v int64) *DescribePxfsUsersResponseBodyDataUsers {
	s.PosixUid = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetRole(v string) *DescribePxfsUsersResponseBodyDataUsers {
	s.Role = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) SetStatus(v string) *DescribePxfsUsersResponseBodyDataUsers {
	s.Status = &v
	return s
}

func (s *DescribePxfsUsersResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
