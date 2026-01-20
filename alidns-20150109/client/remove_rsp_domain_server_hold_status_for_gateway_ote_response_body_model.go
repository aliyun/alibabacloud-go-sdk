// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRspDomainServerHoldStatusForGatewayOteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
	GetAccessDeniedDetail() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail
	SetData(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
	GetData() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData
	SetRecoverableError(v bool) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody
	GetSuccess() *bool
}

type RemoveRspDomainServerHoldStatusForGatewayOteResponseBody struct {
	AccessDeniedDetail *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GetAccessDeniedDetail() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GetData() *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData {
	return s.Data
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) SetAccessDeniedDetail(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) SetData(v *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.Data = v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) SetRecoverableError(v bool) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) SetRequestId(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) SetSuccess(v bool) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBody) Validate() error {
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

type RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail struct {
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

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) SetDomainName(v string) *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *RemoveRspDomainServerHoldStatusForGatewayOteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
