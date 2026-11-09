// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSniConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SetSniConfigResponseBodyAccessDeniedDetail) *SetSniConfigResponseBody
	GetAccessDeniedDetail() *SetSniConfigResponseBodyAccessDeniedDetail
	SetRequestId(v string) *SetSniConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SetSniConfigResponseBody
	GetSuccess() *string
}

type SetSniConfigResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *SetSniConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E4581CCF-62AF-44D9-B5B4-D1DQDC0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of \\`true\\` indicates that the call was successful. A value of \\`false\\` indicates that the call failed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSniConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSniConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetSniConfigResponseBody) GetAccessDeniedDetail() *SetSniConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SetSniConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSniConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SetSniConfigResponseBody) SetAccessDeniedDetail(v *SetSniConfigResponseBodyAccessDeniedDetail) *SetSniConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SetSniConfigResponseBody) SetRequestId(v string) *SetSniConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSniConfigResponseBody) SetSuccess(v string) *SetSniConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetSniConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSniConfigResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// SetSniConfig
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization entity.
	//
	// example:
	//
	// 2063xxxxxxxx533
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the Alibaba Cloud account that owns the authorization entity.
	//
	// example:
	//
	// 1245xxxxx34343
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authorization entity.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic information.
	//
	// example:
	//
	// AQFma6gWZmuoFkMxQ0M3MUVCLTRFQzQtNTI1OS1BQ0Q0LTlBMTFGQkVDOTA3Qw==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of permission denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SetSniConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SetSniConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SetSniConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SetSniConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
