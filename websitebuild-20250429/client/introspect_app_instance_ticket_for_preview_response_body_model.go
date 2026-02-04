// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntrospectAppInstanceTicketForPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *IntrospectAppInstanceTicketForPreviewResponseBodyModule) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetModule() *IntrospectAppInstanceTicketForPreviewResponseBodyModule
	SetRequestId(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *IntrospectAppInstanceTicketForPreviewResponseBody
	GetSynchro() *bool
}

type IntrospectAppInstanceTicketForPreviewResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                            `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *IntrospectAppInstanceTicketForPreviewResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s IntrospectAppInstanceTicketForPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntrospectAppInstanceTicketForPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetModule() *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	return s.Module
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetAccessDeniedDetail(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetAllowRetry(v bool) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetAppName(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.AppName = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetDynamicCode(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetDynamicMessage(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetErrorArgs(v []interface{}) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetModule(v *IntrospectAppInstanceTicketForPreviewResponseBodyModule) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.Module = v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetRequestId(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetRootErrorCode(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetRootErrorMsg(v string) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) SetSynchro(v bool) *IntrospectAppInstanceTicketForPreviewResponseBody {
	s.Synchro = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntrospectAppInstanceTicketForPreviewResponseBodyModule struct {
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	AccessTokenExpiresAt *string `json:"AccessTokenExpiresAt,omitempty" xml:"AccessTokenExpiresAt,omitempty"`
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	AccessTokenIssuedAt *string `json:"AccessTokenIssuedAt,omitempty" xml:"AccessTokenIssuedAt,omitempty"`
	// example:
	//
	// 2108341e17661121129745384e79f9
	AccessTokenValue *string `json:"AccessTokenValue,omitempty" xml:"AccessTokenValue,omitempty"`
	// aliyunPk
	//
	// example:
	//
	// 12343131221311
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// example:
	//
	// {}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// authorization_code
	AuthorizationGrantType *string `json:"AuthorizationGrantType,omitempty" xml:"AuthorizationGrantType,omitempty"`
	// bid
	//
	// example:
	//
	// 123131
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 12343131221311
	ParentPk *string `json:"ParentPk,omitempty" xml:"ParentPk,omitempty"`
	// example:
	//
	// 1768619049924
	RefreshTokenExpiresAt *string `json:"RefreshTokenExpiresAt,omitempty" xml:"RefreshTokenExpiresAt,omitempty"`
	// example:
	//
	// 1768619049924
	RefreshTokenIssuedAt *string `json:"RefreshTokenIssuedAt,omitempty" xml:"RefreshTokenIssuedAt,omitempty"`
	// example:
	//
	// be9750d595b6cd7c93a80b46
	RefreshTokenValue *string `json:"RefreshTokenValue,omitempty" xml:"RefreshTokenValue,omitempty"`
	// example:
	//
	// hdm_33be9750d595b6cd7c93a80b46734b22
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s IntrospectAppInstanceTicketForPreviewResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntrospectAppInstanceTicketForPreviewResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAccessTokenExpiresAt() *string {
	return s.AccessTokenExpiresAt
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAccessTokenIssuedAt() *string {
	return s.AccessTokenIssuedAt
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAccessTokenValue() *string {
	return s.AccessTokenValue
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAttributes() *string {
	return s.Attributes
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetAuthorizationGrantType() *string {
	return s.AuthorizationGrantType
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetBid() *string {
	return s.Bid
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetParentPk() *string {
	return s.ParentPk
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetRefreshTokenExpiresAt() *string {
	return s.RefreshTokenExpiresAt
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetRefreshTokenIssuedAt() *string {
	return s.RefreshTokenIssuedAt
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetRefreshTokenValue() *string {
	return s.RefreshTokenValue
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) GetUuid() *string {
	return s.Uuid
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAccessTokenExpiresAt(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.AccessTokenExpiresAt = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAccessTokenIssuedAt(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.AccessTokenIssuedAt = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAccessTokenValue(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.AccessTokenValue = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAliyunPk(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.AliyunPk = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAttributes(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.Attributes = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetAuthorizationGrantType(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.AuthorizationGrantType = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetBid(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.Bid = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetParentPk(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.ParentPk = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetRefreshTokenExpiresAt(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.RefreshTokenExpiresAt = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetRefreshTokenIssuedAt(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.RefreshTokenIssuedAt = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetRefreshTokenValue(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.RefreshTokenValue = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) SetUuid(v string) *IntrospectAppInstanceTicketForPreviewResponseBodyModule {
	s.Uuid = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
