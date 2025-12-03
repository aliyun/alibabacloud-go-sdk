// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaiJupyterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *StartEaiJupyterResponseBodyAccessDeniedDetail) *StartEaiJupyterResponseBody
	GetAccessDeniedDetail() *StartEaiJupyterResponseBodyAccessDeniedDetail
	SetRequestId(v string) *StartEaiJupyterResponseBody
	GetRequestId() *string
}

type StartEaiJupyterResponseBody struct {
	AccessDeniedDetail *StartEaiJupyterResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEaiJupyterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponseBody) GetAccessDeniedDetail() *StartEaiJupyterResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *StartEaiJupyterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEaiJupyterResponseBody) SetAccessDeniedDetail(v *StartEaiJupyterResponseBodyAccessDeniedDetail) *StartEaiJupyterResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *StartEaiJupyterResponseBody) SetRequestId(v string) *StartEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEaiJupyterResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartEaiJupyterResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// eais:StartEaiJupyter
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 20560152949032****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 170718266783****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFmfh3BZn4dwUQyNzY4MDVELTgzQkUtNTBEOC04QjQyLTNGM0U1QUI5MjhBRA==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ExplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s StartEaiJupyterResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s StartEaiJupyterResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthAction(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) SetPolicyType(v string) *StartEaiJupyterResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *StartEaiJupyterResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
