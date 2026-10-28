// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) *CreatePxfsAccessTokenResponseBody
	GetAccessDeniedDetail() *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail
	SetData(v *CreatePxfsAccessTokenResponseBodyData) *CreatePxfsAccessTokenResponseBody
	GetData() *CreatePxfsAccessTokenResponseBodyData
	SetRequestId(v string) *CreatePxfsAccessTokenResponseBody
	GetRequestId() *string
}

type CreatePxfsAccessTokenResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *CreatePxfsAccessTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePxfsAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePxfsAccessTokenResponseBody) GetAccessDeniedDetail() *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePxfsAccessTokenResponseBody) GetData() *CreatePxfsAccessTokenResponseBodyData {
	return s.Data
}

func (s *CreatePxfsAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePxfsAccessTokenResponseBody) SetAccessDeniedDetail(v *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) *CreatePxfsAccessTokenResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePxfsAccessTokenResponseBody) SetData(v *CreatePxfsAccessTokenResponseBodyData) *CreatePxfsAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreatePxfsAccessTokenResponseBody) SetRequestId(v string) *CreatePxfsAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBody) Validate() error {
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

type CreatePxfsAccessTokenResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// polardbx:QueryColumnarLog
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
	// The diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the missing permission.
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

func (s CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePxfsAccessTokenResponseBodyData struct {
	// The access endpoint.
	//
	// example:
	//
	// agent-files.example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-12-31T23:59:59Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The name of the identity.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The label of the token.
	//
	// example:
	//
	// production-agent
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The port.
	//
	// example:
	//
	// 12049
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The plaintext of the token.
	//
	// example:
	//
	// pxfs_tenant-001.token-001.******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// Token ID
	//
	// example:
	//
	// token-001
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s CreatePxfsAccessTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetIdentityName() *string {
	return s.IdentityName
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreatePxfsAccessTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetEndpoint(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetExpiresAt(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetIdentityName(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.IdentityName = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetLabel(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.Label = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetPort(v int32) *CreatePxfsAccessTokenResponseBodyData {
	s.Port = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetTenantId(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetToken(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) SetTokenId(v string) *CreatePxfsAccessTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *CreatePxfsAccessTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
