// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) *RemoveRspDomainServerHoldStatusForGatewayResponseBody
	GetAccessDeniedDetail() *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail
	SetData(v *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) *RemoveRspDomainServerHoldStatusForGatewayResponseBody
	GetData() *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData
	SetRecoverableError(v bool) *RemoveRspDomainServerHoldStatusForGatewayResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveRspDomainServerHoldStatusForGatewayResponseBody
	GetSuccess() *bool
}

type RemoveRspDomainServerHoldStatusForGatewayResponseBody struct {
	AccessDeniedDetail *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) GetAccessDeniedDetail() *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) GetData() *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData {
	return s.Data
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) SetAccessDeniedDetail(v *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) SetData(v *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	s.Data = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) SetRecoverableError(v bool) *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) SetRequestId(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) SetSuccess(v bool) *RemoveRspDomainServerHoldStatusForGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBody) Validate() error {
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

type RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail struct {
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
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
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

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RemoveRspDomainServerHoldStatusForGatewayResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ServerHoldStatus *string `json:"ServerHoldStatus,omitempty" xml:"ServerHoldStatus,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) GetServerHoldStatus() *string {
	return s.ServerHoldStatus
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) SetServerHoldStatus(v string) *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData {
	s.ServerHoldStatus = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
