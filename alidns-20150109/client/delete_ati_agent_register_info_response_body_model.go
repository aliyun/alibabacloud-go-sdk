// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAtiAgentRegisterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *DeleteAtiAgentRegisterInfoResponseBody
	GetAccessDeniedDetail() *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail
	SetRequestId(v string) *DeleteAtiAgentRegisterInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAtiAgentRegisterInfoResponseBody
	GetSuccess() *bool
}

type DeleteAtiAgentRegisterInfoResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAtiAgentRegisterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiAgentRegisterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) GetAccessDeniedDetail() *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) SetAccessDeniedDetail(v *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) *DeleteAtiAgentRegisterInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) SetRequestId(v string) *DeleteAtiAgentRegisterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) SetSuccess(v bool) *DeleteAtiAgentRegisterInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGateway
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
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
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

func (s DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
