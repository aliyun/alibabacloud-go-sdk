// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTmpIdentityForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetUserTmpIdentityForPartnerResponseBodyData) *GetUserTmpIdentityForPartnerResponseBody
	GetData() *GetUserTmpIdentityForPartnerResponseBodyData
	SetErrorCode(v string) *GetUserTmpIdentityForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetUserTmpIdentityForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetUserTmpIdentityForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserTmpIdentityForPartnerResponseBody
	GetSuccess() *bool
}

type GetUserTmpIdentityForPartnerResponseBody struct {
	Data *GetUserTmpIdentityForPartnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserTmpIdentityForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserTmpIdentityForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTmpIdentityForPartnerResponseBody) GetData() *GetUserTmpIdentityForPartnerResponseBodyData {
	return s.Data
}

func (s *GetUserTmpIdentityForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserTmpIdentityForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetUserTmpIdentityForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserTmpIdentityForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserTmpIdentityForPartnerResponseBody) SetData(v *GetUserTmpIdentityForPartnerResponseBodyData) *GetUserTmpIdentityForPartnerResponseBody {
	s.Data = v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBody) SetErrorCode(v string) *GetUserTmpIdentityForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBody) SetErrorMsg(v string) *GetUserTmpIdentityForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBody) SetRequestId(v string) *GetUserTmpIdentityForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBody) SetSuccess(v bool) *GetUserTmpIdentityForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserTmpIdentityForPartnerResponseBodyData struct {
	Credentials       *GetUserTmpIdentityForPartnerResponseBodyDataCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	HasCustomRoleAuth *bool                                                    `json:"HasCustomRoleAuth,omitempty" xml:"HasCustomRoleAuth,omitempty"`
}

func (s GetUserTmpIdentityForPartnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserTmpIdentityForPartnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserTmpIdentityForPartnerResponseBodyData) GetCredentials() *GetUserTmpIdentityForPartnerResponseBodyDataCredentials {
	return s.Credentials
}

func (s *GetUserTmpIdentityForPartnerResponseBodyData) GetHasCustomRoleAuth() *bool {
	return s.HasCustomRoleAuth
}

func (s *GetUserTmpIdentityForPartnerResponseBodyData) SetCredentials(v *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) *GetUserTmpIdentityForPartnerResponseBodyData {
	s.Credentials = v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyData) SetHasCustomRoleAuth(v bool) *GetUserTmpIdentityForPartnerResponseBodyData {
	s.HasCustomRoleAuth = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetUserTmpIdentityForPartnerResponseBodyDataCredentials struct {
	// example:
	//
	// 7t2Rzt37ND7-xxx
	EncryptedAccessKeyId *string `json:"EncryptedAccessKeyId,omitempty" xml:"EncryptedAccessKeyId,omitempty"`
	// example:
	//
	// 7t2Rzt37ND7-xxx
	EncryptedAccessKeySecret *string `json:"EncryptedAccessKeySecret,omitempty" xml:"EncryptedAccessKeySecret,omitempty"`
	// example:
	//
	// 7t2Rzt37ND7-xxx
	EncryptedSecurityToken *string `json:"EncryptedSecurityToken,omitempty" xml:"EncryptedSecurityToken,omitempty"`
	// example:
	//
	// 2025-04-29T14:00:00Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
}

func (s GetUserTmpIdentityForPartnerResponseBodyDataCredentials) String() string {
	return dara.Prettify(s)
}

func (s GetUserTmpIdentityForPartnerResponseBodyDataCredentials) GoString() string {
	return s.String()
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) GetEncryptedAccessKeyId() *string {
	return s.EncryptedAccessKeyId
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) GetEncryptedAccessKeySecret() *string {
	return s.EncryptedAccessKeySecret
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) GetEncryptedSecurityToken() *string {
	return s.EncryptedSecurityToken
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) SetEncryptedAccessKeyId(v string) *GetUserTmpIdentityForPartnerResponseBodyDataCredentials {
	s.EncryptedAccessKeyId = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) SetEncryptedAccessKeySecret(v string) *GetUserTmpIdentityForPartnerResponseBodyDataCredentials {
	s.EncryptedAccessKeySecret = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) SetEncryptedSecurityToken(v string) *GetUserTmpIdentityForPartnerResponseBodyDataCredentials {
	s.EncryptedSecurityToken = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) SetExpiration(v string) *GetUserTmpIdentityForPartnerResponseBodyDataCredentials {
	s.Expiration = &v
	return s
}

func (s *GetUserTmpIdentityForPartnerResponseBodyDataCredentials) Validate() error {
	return dara.Validate(s)
}
