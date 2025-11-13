// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainStatusOteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) *UpdateRspDomainStatusOteResponseBody
	GetAccessDeniedDetail() *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail
	SetData(v *UpdateRspDomainStatusOteResponseBodyData) *UpdateRspDomainStatusOteResponseBody
	GetData() *UpdateRspDomainStatusOteResponseBodyData
	SetRecoverableError(v bool) *UpdateRspDomainStatusOteResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *UpdateRspDomainStatusOteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRspDomainStatusOteResponseBody
	GetSuccess() *bool
}

type UpdateRspDomainStatusOteResponseBody struct {
	AccessDeniedDetail *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *UpdateRspDomainStatusOteResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	RecoverableError *bool `json:"RecoverableError,omitempty" xml:"RecoverableError,omitempty"`
	// example:
	//
	// 0629502C-XXXX-5DC9-A8ED-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRspDomainStatusOteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteResponseBody) GetAccessDeniedDetail() *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateRspDomainStatusOteResponseBody) GetData() *UpdateRspDomainStatusOteResponseBodyData {
	return s.Data
}

func (s *UpdateRspDomainStatusOteResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *UpdateRspDomainStatusOteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRspDomainStatusOteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRspDomainStatusOteResponseBody) SetAccessDeniedDetail(v *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) *UpdateRspDomainStatusOteResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBody) SetData(v *UpdateRspDomainStatusOteResponseBodyData) *UpdateRspDomainStatusOteResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBody) SetRecoverableError(v bool) *UpdateRspDomainStatusOteResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBody) SetRequestId(v string) *UpdateRspDomainStatusOteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBody) SetSuccess(v bool) *UpdateRspDomainStatusOteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBody) Validate() error {
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

type UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail struct {
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

func (s UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainStatusOteResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StatusList []*UpdateRspDomainStatusOteResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s UpdateRspDomainStatusOteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainStatusOteResponseBodyData) GetStatusList() []*UpdateRspDomainStatusOteResponseBodyDataStatusList {
	return s.StatusList
}

func (s *UpdateRspDomainStatusOteResponseBodyData) SetDomainName(v string) *UpdateRspDomainStatusOteResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyData) SetStatusList(v []*UpdateRspDomainStatusOteResponseBodyDataStatusList) *UpdateRspDomainStatusOteResponseBodyData {
	s.StatusList = v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyData) Validate() error {
	if s.StatusList != nil {
		for _, item := range s.StatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRspDomainStatusOteResponseBodyDataStatusList struct {
	// example:
	//
	// uptp.test.abchina.com.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// serverUpdateProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainStatusOteResponseBodyDataStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) SetDomainName(v string) *UpdateRspDomainStatusOteResponseBodyDataStatusList {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) SetStatus(v string) *UpdateRspDomainStatusOteResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) SetStatusMsg(v string) *UpdateRspDomainStatusOteResponseBodyDataStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainStatusOteResponseBodyDataStatusList) Validate() error {
	return dara.Validate(s)
}
