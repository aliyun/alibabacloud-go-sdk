// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerHoldStatusOteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) *UpdateRspDomainServerHoldStatusOteResponseBody
	GetAccessDeniedDetail() *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail
	SetData(v *UpdateRspDomainServerHoldStatusOteResponseBodyData) *UpdateRspDomainServerHoldStatusOteResponseBody
	GetData() *UpdateRspDomainServerHoldStatusOteResponseBodyData
	SetRecoverableError(v bool) *UpdateRspDomainServerHoldStatusOteResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *UpdateRspDomainServerHoldStatusOteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRspDomainServerHoldStatusOteResponseBody
	GetSuccess() *bool
}

type UpdateRspDomainServerHoldStatusOteResponseBody struct {
	AccessDeniedDetail *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *UpdateRspDomainServerHoldStatusOteResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	RecoverableError *bool `json:"RecoverableError,omitempty" xml:"RecoverableError,omitempty"`
	// example:
	//
	// 0629502C-XXXX-XXXX-XXXX-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRspDomainServerHoldStatusOteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerHoldStatusOteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) GetAccessDeniedDetail() *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) GetData() *UpdateRspDomainServerHoldStatusOteResponseBodyData {
	return s.Data
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) SetAccessDeniedDetail(v *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) *UpdateRspDomainServerHoldStatusOteResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) SetData(v *UpdateRspDomainServerHoldStatusOteResponseBodyData) *UpdateRspDomainServerHoldStatusOteResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) SetRecoverableError(v bool) *UpdateRspDomainServerHoldStatusOteResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) SetRequestId(v string) *UpdateRspDomainServerHoldStatusOteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) SetSuccess(v bool) *UpdateRspDomainServerHoldStatusOteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBody) Validate() error {
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

type UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail struct {
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
	// 1046973331XXXX
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

func (s UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainServerHoldStatusOteResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// enable
	ServerHoldStatus *string `json:"ServerHoldStatus,omitempty" xml:"ServerHoldStatus,omitempty"`
}

func (s UpdateRspDomainServerHoldStatusOteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerHoldStatusOteResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyData) GetServerHoldStatus() *string {
	return s.ServerHoldStatus
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyData) SetDomainName(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyData) SetServerHoldStatus(v string) *UpdateRspDomainServerHoldStatusOteResponseBodyData {
	s.ServerHoldStatus = &v
	return s
}

func (s *UpdateRspDomainServerHoldStatusOteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
