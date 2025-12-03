// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaiJupyterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *StopEaiJupyterResponseBodyAccessDeniedDetail) *StopEaiJupyterResponseBody
	GetAccessDeniedDetail() *StopEaiJupyterResponseBodyAccessDeniedDetail
	SetRequestId(v string) *StopEaiJupyterResponseBody
	GetRequestId() *string
}

type StopEaiJupyterResponseBody struct {
	AccessDeniedDetail *StopEaiJupyterResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopEaiJupyterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponseBody) GetAccessDeniedDetail() *StopEaiJupyterResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *StopEaiJupyterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopEaiJupyterResponseBody) SetAccessDeniedDetail(v *StopEaiJupyterResponseBodyAccessDeniedDetail) *StopEaiJupyterResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *StopEaiJupyterResponseBody) SetRequestId(v string) *StopEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopEaiJupyterResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopEaiJupyterResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// eais:StopEaiJupyter
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
	// AQFmj0FOZo9BTjMyQTFDRkIzLUE5MTItNUIwNC1BQzkxLTcyMUFFQTUyQjhGQQ==
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

func (s StopEaiJupyterResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s StopEaiJupyterResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthAction(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) SetPolicyType(v string) *StopEaiJupyterResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *StopEaiJupyterResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
