// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsAccessTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) *DescribePxfsAccessTokensResponseBody
	GetAccessDeniedDetail() *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfsAccessTokensResponseBodyData) *DescribePxfsAccessTokensResponseBody
	GetData() *DescribePxfsAccessTokensResponseBodyData
	SetRequestId(v string) *DescribePxfsAccessTokensResponseBody
	GetRequestId() *string
}

type DescribePxfsAccessTokensResponseBody struct {
	// The detailed reason why access was denied.
	AccessDeniedDetail *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DescribePxfsAccessTokensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfsAccessTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensResponseBody) GetAccessDeniedDetail() *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfsAccessTokensResponseBody) GetData() *DescribePxfsAccessTokensResponseBodyData {
	return s.Data
}

func (s *DescribePxfsAccessTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfsAccessTokensResponseBody) SetAccessDeniedDetail(v *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) *DescribePxfsAccessTokensResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfsAccessTokensResponseBody) SetData(v *DescribePxfsAccessTokensResponseBodyData) *DescribePxfsAccessTokensResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfsAccessTokensResponseBody) SetRequestId(v string) *DescribePxfsAccessTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBody) Validate() error {
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

type DescribePxfsAccessTokensResponseBodyAccessDeniedDetail struct {
	// The authorization action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authorization in the request.
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
	// The diagnostic information.
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

func (s DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfsAccessTokensResponseBodyData struct {
	// The list of access tokens.
	AccessTokens []*DescribePxfsAccessTokensResponseBodyDataAccessTokens `json:"AccessTokens,omitempty" xml:"AccessTokens,omitempty" type:"Repeated"`
}

func (s DescribePxfsAccessTokensResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensResponseBodyData) GetAccessTokens() []*DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	return s.AccessTokens
}

func (s *DescribePxfsAccessTokensResponseBodyData) SetAccessTokens(v []*DescribePxfsAccessTokensResponseBodyDataAccessTokens) *DescribePxfsAccessTokensResponseBodyData {
	s.AccessTokens = v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyData) Validate() error {
	if s.AccessTokens != nil {
		for _, item := range s.AccessTokens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePxfsAccessTokensResponseBodyDataAccessTokens struct {
	// The creation time.
	//
	// example:
	//
	// 2026-09-03T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-12-31T23:59:59Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The identity name.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The token label.
	//
	// example:
	//
	// production-agent
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Token ID
	//
	// example:
	//
	// token-001
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s DescribePxfsAccessTokensResponseBodyDataAccessTokens) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensResponseBodyDataAccessTokens) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetLabel() *string {
	return s.Label
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetStatus() *string {
	return s.Status
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) GetTokenId() *string {
	return s.TokenId
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetCreateTime(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.CreateTime = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetExpiresAt(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.ExpiresAt = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetIdentityName(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.IdentityName = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetLabel(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.Label = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetStatus(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.Status = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) SetTokenId(v string) *DescribePxfsAccessTokensResponseBodyDataAccessTokens {
	s.TokenId = &v
	return s
}

func (s *DescribePxfsAccessTokensResponseBodyDataAccessTokens) Validate() error {
	return dara.Validate(s)
}
