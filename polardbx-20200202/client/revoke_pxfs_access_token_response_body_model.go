// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePxfsAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) *RevokePxfsAccessTokenResponseBody
	GetAccessDeniedDetail() *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail
	SetData(v *RevokePxfsAccessTokenResponseBodyData) *RevokePxfsAccessTokenResponseBody
	GetData() *RevokePxfsAccessTokenResponseBodyData
	SetRequestId(v string) *RevokePxfsAccessTokenResponseBody
	GetRequestId() *string
}

type RevokePxfsAccessTokenResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *RevokePxfsAccessTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokePxfsAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokePxfsAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePxfsAccessTokenResponseBody) GetAccessDeniedDetail() *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RevokePxfsAccessTokenResponseBody) GetData() *RevokePxfsAccessTokenResponseBodyData {
	return s.Data
}

func (s *RevokePxfsAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokePxfsAccessTokenResponseBody) SetAccessDeniedDetail(v *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) *RevokePxfsAccessTokenResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RevokePxfsAccessTokenResponseBody) SetData(v *RevokePxfsAccessTokenResponseBodyData) *RevokePxfsAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *RevokePxfsAccessTokenResponseBody) SetRequestId(v string) *RevokePxfsAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBody) Validate() error {
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

type RevokePxfsAccessTokenResponseBodyAccessDeniedDetail struct {
	// The authorization action. See the description above.
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
	// The owner ID of the authorization principal.
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

func (s RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RevokePxfsAccessTokenResponseBodyData struct {
	// Indicates whether the token has been revoked.
	//
	// example:
	//
	// true
	Revoked *bool `json:"Revoked,omitempty" xml:"Revoked,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Token ID
	//
	// example:
	//
	// token-001
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s RevokePxfsAccessTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RevokePxfsAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *RevokePxfsAccessTokenResponseBodyData) GetRevoked() *bool {
	return s.Revoked
}

func (s *RevokePxfsAccessTokenResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokePxfsAccessTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *RevokePxfsAccessTokenResponseBodyData) SetRevoked(v bool) *RevokePxfsAccessTokenResponseBodyData {
	s.Revoked = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyData) SetTenantId(v string) *RevokePxfsAccessTokenResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyData) SetTokenId(v string) *RevokePxfsAccessTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *RevokePxfsAccessTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
