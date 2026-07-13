// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
	GetAccessDeniedDetail() *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail
	SetData(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
	GetData() *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData
	SetRecoverableError(v bool) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
	GetRecoverableError() *bool
	SetRequestId(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody
	GetSuccess() *bool
}

type UpdateRspDomainServerProhibitStatusForGatewayResponseBody struct {
	// The details about the access denial. This field is returned only when Resource Access Management (RAM) authentication fails.
	AccessDeniedDetail *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data.
	Data *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the request can be retried if it fails. true: The request can be retried. false: The request cannot be retried.
	//
	// example:
	//
	// true
	RecoverableError *bool `json:"RecoverableError,omitempty" xml:"RecoverableError,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 0629502C-6224-5DC9-A8ED-2ED73A2E3931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GetAccessDeniedDetail() *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GetData() *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData {
	return s.Data
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GetRecoverableError() *bool {
	return s.RecoverableError
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) SetAccessDeniedDetail(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) SetData(v *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) SetRecoverableError(v bool) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	s.RecoverableError = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) SetRequestId(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) SetSuccess(v bool) *UpdateRspDomainServerProhibitStatusForGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBody) Validate() error {
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

type UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// UpdateRspDomainServerProhibitStatusForGateway
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
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted complete diagnostic information.
	//
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason why authentication failed. Valid values:
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

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthAction(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) SetPolicyType(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status information of the task.
	StatusList []*UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) GetStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList {
	return s.StatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) SetStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData {
	s.StatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyData) Validate() error {
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

type UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList struct {
	// The domain name.
	//
	// example:
	//
	// uptp.test.abchina.com.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The current status of the domain name.
	//
	// example:
	//
	// serverUpdateProhibited
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message for the domain name status.
	//
	// example:
	//
	// 实名认证未通过，增加serverUpdateProhibited状态
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayResponseBodyDataStatusList) Validate() error {
	return dara.Validate(s)
}
