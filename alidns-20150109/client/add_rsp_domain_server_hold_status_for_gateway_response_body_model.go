// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) *AddRspDomainServerHoldStatusForGatewayResponseBody
	GetAccessDeniedDetail() *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail
	SetData(v *AddRspDomainServerHoldStatusForGatewayResponseBodyData) *AddRspDomainServerHoldStatusForGatewayResponseBody
	GetData() *AddRspDomainServerHoldStatusForGatewayResponseBodyData
	SetRecoverableError(v bool) *AddRspDomainServerHoldStatusForGatewayResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *AddRspDomainServerHoldStatusForGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRspDomainServerHoldStatusForGatewayResponseBody
	GetSuccess() *bool
}

type AddRspDomainServerHoldStatusForGatewayResponseBody struct {
	// The details about the access denial. This field is returned only when Resource Access Management (RAM) verification fails.
	AccessDeniedDetail *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data list.
	Data *AddRspDomainServerHoldStatusForGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether you can retry the request if it fails. A value of `true` indicates that you can retry. A value of `false` indicates that you cannot.
	//
	// example:
	//
	// true
	RecoverableError *bool `json:"RecoverableError,omitempty" xml:"RecoverableError,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 0629502C-XXXX-5DC9-XXXX-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of `true` indicates success. A value of `false` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) GetAccessDeniedDetail() *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) GetData() *AddRspDomainServerHoldStatusForGatewayResponseBodyData {
	return s.Data
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) SetAccessDeniedDetail(v *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) *AddRspDomainServerHoldStatusForGatewayResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) SetData(v *AddRspDomainServerHoldStatusForGatewayResponseBodyData) *AddRspDomainServerHoldStatusForGatewayResponseBody {
	s.Data = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) SetRecoverableError(v bool) *AddRspDomainServerHoldStatusForGatewayResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) SetRequestId(v string) *AddRspDomainServerHoldStatusForGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) SetSuccess(v bool) *AddRspDomainServerHoldStatusForGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// AddRspDomainServerHoldStatusForGateway
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorized entity.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the owner of the authorized entity.
	//
	// example:
	//
	// 1046973331XXXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The complete diagnostic information after encryption.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason why the authorization failed. Valid values:
	//
	// - ExplicitDeny: The access is explicitly denied.
	//
	// - ImplicitDeny: The access is implicitly denied.
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

func (s AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AddRspDomainServerHoldStatusForGatewayResponseBodyData struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyData) SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
