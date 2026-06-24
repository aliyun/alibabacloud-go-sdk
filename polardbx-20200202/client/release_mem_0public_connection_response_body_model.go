// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseMem0PublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) *ReleaseMem0PublicConnectionResponseBody
	GetAccessDeniedDetail() *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail
	SetRequestId(v string) *ReleaseMem0PublicConnectionResponseBody
	GetRequestId() *string
}

type ReleaseMem0PublicConnectionResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseMem0PublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseMem0PublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseMem0PublicConnectionResponseBody) GetAccessDeniedDetail() *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReleaseMem0PublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseMem0PublicConnectionResponseBody) SetAccessDeniedDetail(v *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) *ReleaseMem0PublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBody) SetRequestId(v string) *ReleaseMem0PublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail struct {
	// Same as above.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// Same as above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// Same as above.
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
	// The type of missing permission.
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

func (s ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReleaseMem0PublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
