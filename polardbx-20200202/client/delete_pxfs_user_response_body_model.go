// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeletePxfsUserResponseBodyAccessDeniedDetail) *DeletePxfsUserResponseBody
	GetAccessDeniedDetail() *DeletePxfsUserResponseBodyAccessDeniedDetail
	SetData(v *DeletePxfsUserResponseBodyData) *DeletePxfsUserResponseBody
	GetData() *DeletePxfsUserResponseBodyData
	SetRequestId(v string) *DeletePxfsUserResponseBody
	GetRequestId() *string
}

type DeletePxfsUserResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DeletePxfsUserResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DeletePxfsUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePxfsUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePxfsUserResponseBody) GetAccessDeniedDetail() *DeletePxfsUserResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeletePxfsUserResponseBody) GetData() *DeletePxfsUserResponseBodyData {
	return s.Data
}

func (s *DeletePxfsUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePxfsUserResponseBody) SetAccessDeniedDetail(v *DeletePxfsUserResponseBodyAccessDeniedDetail) *DeletePxfsUserResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeletePxfsUserResponseBody) SetData(v *DeletePxfsUserResponseBodyData) *DeletePxfsUserResponseBody {
	s.Data = v
	return s
}

func (s *DeletePxfsUserResponseBody) SetRequestId(v string) *DeletePxfsUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePxfsUserResponseBody) Validate() error {
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

type DeletePxfsUserResponseBodyAccessDeniedDetail struct {
	// The authorization action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The description is the same as above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authorization principal.
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
	// The type of the no-permission error.
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

func (s DeletePxfsUserResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsUserResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeletePxfsUserResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeletePxfsUserResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeletePxfsUserResponseBodyData struct {
	// Indicates whether the identity is deleted.
	//
	// example:
	//
	// true
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The identity name.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeletePxfsUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePxfsUserResponseBodyData) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeletePxfsUserResponseBodyData) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DeletePxfsUserResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *DeletePxfsUserResponseBodyData) SetDeleted(v bool) *DeletePxfsUserResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DeletePxfsUserResponseBodyData) SetIdentityName(v string) *DeletePxfsUserResponseBodyData {
	s.IdentityName = &v
	return s
}

func (s *DeletePxfsUserResponseBodyData) SetTenantId(v string) *DeletePxfsUserResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DeletePxfsUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
