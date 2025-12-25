// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryOriginImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RecoveryOriginImageResponseBodyAccessDeniedDetail) *RecoveryOriginImageResponseBody
	GetAccessDeniedDetail() *RecoveryOriginImageResponseBodyAccessDeniedDetail
	SetCode(v int64) *RecoveryOriginImageResponseBody
	GetCode() *int64
	SetMessage(v string) *RecoveryOriginImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecoveryOriginImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecoveryOriginImageResponseBody
	GetSuccess() *bool
}

type RecoveryOriginImageResponseBody struct {
	AccessDeniedDetail *RecoveryOriginImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3BCAD49D-2AC1-13EB-AC19-8C7A46CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecoveryOriginImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoveryOriginImageResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponseBody) GetAccessDeniedDetail() *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RecoveryOriginImageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RecoveryOriginImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecoveryOriginImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoveryOriginImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecoveryOriginImageResponseBody) SetAccessDeniedDetail(v *RecoveryOriginImageResponseBodyAccessDeniedDetail) *RecoveryOriginImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetCode(v int64) *RecoveryOriginImageResponseBody {
	s.Code = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetMessage(v string) *RecoveryOriginImageResponseBody {
	s.Message = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetRequestId(v string) *RecoveryOriginImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) SetSuccess(v bool) *RecoveryOriginImageResponseBody {
	s.Success = &v
	return s
}

func (s *RecoveryOriginImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecoveryOriginImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s RecoveryOriginImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RecoveryOriginImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RecoveryOriginImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RecoveryOriginImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
