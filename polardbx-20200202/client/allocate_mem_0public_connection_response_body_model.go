// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateMem0PublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) *AllocateMem0PublicConnectionResponseBody
	GetAccessDeniedDetail() *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail
	SetRequestId(v string) *AllocateMem0PublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateMem0PublicConnectionResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateMem0PublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateMem0PublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateMem0PublicConnectionResponseBody) GetAccessDeniedDetail() *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AllocateMem0PublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateMem0PublicConnectionResponseBody) SetAccessDeniedDetail(v *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) *AllocateMem0PublicConnectionResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBody) SetRequestId(v string) *AllocateMem0PublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail struct {
	// As described above.
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
	// As described above.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// As described above.
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
	// PolicyType
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AllocateMem0PublicConnectionResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
