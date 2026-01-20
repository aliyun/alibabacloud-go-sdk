// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRspDomainServerHoldStatusForGatewayOteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) *AddRspDomainServerHoldStatusForGatewayOteResponseBody
	GetAccessDeniedDetail() *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail
	SetData(v *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) *AddRspDomainServerHoldStatusForGatewayOteResponseBody
	GetData() *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData
	SetRecoverableError(v bool) *AddRspDomainServerHoldStatusForGatewayOteResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRspDomainServerHoldStatusForGatewayOteResponseBody
	GetSuccess() *bool
}

type AddRspDomainServerHoldStatusForGatewayOteResponseBody struct {
	AccessDeniedDetail *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	RecoverableError *bool `json:"RecoverableError,omitempty" xml:"RecoverableError,omitempty"`
	// example:
	//
	// 0629502C-6224-5DC9-A8ED-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBody) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) GetAccessDeniedDetail() *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) GetData() *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData {
	return s.Data
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) SetAccessDeniedDetail(v *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) SetData(v *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.Data = v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) SetRecoverableError(v bool) *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) SetRequestId(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) SetSuccess(v bool) *AddRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.Success = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBody) Validate() error {
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

type AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AddRspDomainServerHoldStatusForGatewayOteResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) SetDomainName(v string) *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *AddRspDomainServerHoldStatusForGatewayOteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
