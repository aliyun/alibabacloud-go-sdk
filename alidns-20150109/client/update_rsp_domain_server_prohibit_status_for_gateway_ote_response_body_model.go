// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayOteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
	GetAccessDeniedDetail() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail
	SetData(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
	GetData() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData
	SetRecoverableError(v bool) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody
	GetSuccess() *bool
}

type UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody struct {
	AccessDeniedDetail *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GetAccessDeniedDetail() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GetData() *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData {
	return s.Data
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) SetAccessDeniedDetail(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) SetData(v *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) SetRecoverableError(v bool) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) SetRequestId(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) SetSuccess(v bool) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBody) Validate() error {
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

type UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail struct {
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

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData struct {
	// example:
	//
	// example.com
	DomainName *string                                                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StatusList []*UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) GetStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList {
	return s.StatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) SetStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData {
	s.StatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyData) Validate() error {
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

type UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList struct {
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

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteResponseBodyDataStatusList) Validate() error {
	return dara.Validate(s)
}
