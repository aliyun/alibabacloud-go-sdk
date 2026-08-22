// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchAccessProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) *ModifyOpenSearchAccessProtocolResponseBody
	GetAccessDeniedDetail() *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail
	SetMessage(v string) *ModifyOpenSearchAccessProtocolResponseBody
	GetMessage() *string
	SetProtocol(v string) *ModifyOpenSearchAccessProtocolResponseBody
	GetProtocol() *string
	SetRequestId(v string) *ModifyOpenSearchAccessProtocolResponseBody
	GetRequestId() *string
}

type ModifyOpenSearchAccessProtocolResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response message. "success" is returned for a successful request. An error code is returned for a failed request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The backend protocol. Valid values:
	//
	// - **HTTP*	- (default): supports association with HTTPS, HTTP, and QUIC listeners.
	//
	// - **HTTPS**: supports association with HTTPS listeners.
	//
	// - **gRPC**: supports association with HTTPS and QUIC listeners.
	//
	// > If **ServerGroupType*	- is set to **Fc**, you do not need to configure the backend protocol.
	//
	// example:
	//
	// icmp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOpenSearchAccessProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchAccessProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) GetAccessDeniedDetail() *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) SetAccessDeniedDetail(v *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) *ModifyOpenSearchAccessProtocolResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) SetMessage(v string) *ModifyOpenSearchAccessProtocolResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) SetProtocol(v string) *ModifyOpenSearchAccessProtocolResponseBody {
	s.Protocol = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) SetRequestId(v string) *ModifyOpenSearchAccessProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type used for authentication in the request. Valid values:
	//
	// - SubUser: RAM user
	//
	// - AssumedRoleUser: RAM role
	//
	// - Federated: SSO federated identity
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
	// The type of the permission denial.
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

func (s ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyOpenSearchAccessProtocolResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
