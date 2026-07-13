// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiRegistrantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateAtiRegistrantResponseBodyAccessDeniedDetail) *CreateAtiRegistrantResponseBody
	GetAccessDeniedDetail() *CreateAtiRegistrantResponseBodyAccessDeniedDetail
	SetCreateTimestamp(v int64) *CreateAtiRegistrantResponseBody
	GetCreateTimestamp() *int64
	SetName(v string) *CreateAtiRegistrantResponseBody
	GetName() *string
	SetRegistrantId(v string) *CreateAtiRegistrantResponseBody
	GetRegistrantId() *string
	SetRequestId(v string) *CreateAtiRegistrantResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateAtiRegistrantResponseBody
	GetStatus() *string
}

type CreateAtiRegistrantResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *CreateAtiRegistrantResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The name of the real-name verified registrant.
	//
	// example:
	//
	// 张xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the real-name verified registrant.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The real-name verification status. Valid values:
	//
	// - Approved.
	//
	// - Under review.
	//
	// - Rejected.
	//
	// example:
	//
	// 审核通过
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAtiRegistrantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiRegistrantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAtiRegistrantResponseBody) GetAccessDeniedDetail() *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateAtiRegistrantResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *CreateAtiRegistrantResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateAtiRegistrantResponseBody) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *CreateAtiRegistrantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAtiRegistrantResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateAtiRegistrantResponseBody) SetAccessDeniedDetail(v *CreateAtiRegistrantResponseBodyAccessDeniedDetail) *CreateAtiRegistrantResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateAtiRegistrantResponseBody) SetCreateTimestamp(v int64) *CreateAtiRegistrantResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *CreateAtiRegistrantResponseBody) SetName(v string) *CreateAtiRegistrantResponseBody {
	s.Name = &v
	return s
}

func (s *CreateAtiRegistrantResponseBody) SetRegistrantId(v string) *CreateAtiRegistrantResponseBody {
	s.RegistrantId = &v
	return s
}

func (s *CreateAtiRegistrantResponseBody) SetRequestId(v string) *CreateAtiRegistrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAtiRegistrantResponseBody) SetStatus(v string) *CreateAtiRegistrantResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAtiRegistrantResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAtiRegistrantResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted complete diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: Explicit deny.
	//
	// - ImplicitDeny: Implicit deny.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateAtiRegistrantResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiRegistrantResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateAtiRegistrantResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateAtiRegistrantResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
