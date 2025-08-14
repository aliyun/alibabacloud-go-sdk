// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelIProductionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CancelIProductionJobResponseBodyAccessDeniedDetail) *CancelIProductionJobResponseBody
	GetAccessDeniedDetail() *CancelIProductionJobResponseBodyAccessDeniedDetail
	SetMessage(v string) *CancelIProductionJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelIProductionJobResponseBody
	GetRequestId() *string
}

type CancelIProductionJobResponseBody struct {
	AccessDeniedDetail *CancelIProductionJobResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelIProductionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobResponseBody) GetAccessDeniedDetail() *CancelIProductionJobResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CancelIProductionJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelIProductionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelIProductionJobResponseBody) SetAccessDeniedDetail(v *CancelIProductionJobResponseBodyAccessDeniedDetail) *CancelIProductionJobResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CancelIProductionJobResponseBody) SetMessage(v string) *CancelIProductionJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelIProductionJobResponseBody) SetRequestId(v string) *CancelIProductionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelIProductionJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type CancelIProductionJobResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// ice:CancelIProductionJob
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// ****4522705967****
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// ****82303720****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// ******AAZ/h8jzNEODc5QUUyLUZCOTAtNUQyQy1BMEFBLUUzODQxODUx******==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// AssumeRolePolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CancelIProductionJobResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CancelIProductionJobResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CancelIProductionJobResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CancelIProductionJobResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
