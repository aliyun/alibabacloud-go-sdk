// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppInstanceTicketResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppInstanceTicketResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppInstanceTicketResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppInstanceTicketResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppInstanceTicketResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppInstanceTicketResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppInstanceTicketResponseBodyModule) *CreateAppInstanceTicketResponseBody
	GetModule() *CreateAppInstanceTicketResponseBodyModule
	SetRequestId(v string) *CreateAppInstanceTicketResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppInstanceTicketResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppInstanceTicketResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppInstanceTicketResponseBody
	GetSynchro() *bool
}

type CreateAppInstanceTicketResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Whether to allow retry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// App name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it means that the input request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERRROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Response data
	Module *CreateAppInstanceTicketResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// Exception message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s CreateAppInstanceTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppInstanceTicketResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppInstanceTicketResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInstanceTicketResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppInstanceTicketResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppInstanceTicketResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppInstanceTicketResponseBody) GetModule() *CreateAppInstanceTicketResponseBodyModule {
	return s.Module
}

func (s *CreateAppInstanceTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppInstanceTicketResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppInstanceTicketResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppInstanceTicketResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppInstanceTicketResponseBody) SetAccessDeniedDetail(v string) *CreateAppInstanceTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetAllowRetry(v bool) *CreateAppInstanceTicketResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetAppName(v string) *CreateAppInstanceTicketResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetDynamicCode(v string) *CreateAppInstanceTicketResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetDynamicMessage(v string) *CreateAppInstanceTicketResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetErrorArgs(v []interface{}) *CreateAppInstanceTicketResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetModule(v *CreateAppInstanceTicketResponseBodyModule) *CreateAppInstanceTicketResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetRequestId(v string) *CreateAppInstanceTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetRootErrorCode(v string) *CreateAppInstanceTicketResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetRootErrorMsg(v string) *CreateAppInstanceTicketResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) SetSynchro(v bool) *CreateAppInstanceTicketResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppInstanceTicketResponseBodyModule struct {
	// Access token expiration time
	//
	// example:
	//
	// 1768619049924
	AccessTokenExpiresAt *string `json:"AccessTokenExpiresAt,omitempty" xml:"AccessTokenExpiresAt,omitempty"`
	// Access token issuance time
	//
	// example:
	//
	// 1768619049924
	AccessTokenIssuedAt *string `json:"AccessTokenIssuedAt,omitempty" xml:"AccessTokenIssuedAt,omitempty"`
	// Access token value
	//
	// example:
	//
	// be9750d595b6cd7c93a80b46
	AccessTokenValue *string `json:"AccessTokenValue,omitempty" xml:"AccessTokenValue,omitempty"`
	// User ID
	//
	// example:
	//
	// 12343131221311
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// Extended attributes
	//
	// example:
	//
	// {\\"resourceGroupId\\":\\"\\",\\"page\\":1,\\"size\\":10,\\"selected\\":\\"[\\\\\\"kvstore\\\\\\"]\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Authorization grant type
	//
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
	// User ID
	//
	// example:
	//
	// 12343131221311
	ParentPk *string `json:"ParentPk,omitempty" xml:"ParentPk,omitempty"`
	// Refresh token expiration time
	//
	// example:
	//
	// 1768619049924
	RefreshTokenExpiresAt *string `json:"RefreshTokenExpiresAt,omitempty" xml:"RefreshTokenExpiresAt,omitempty"`
	// Refresh token issuance time
	//
	// example:
	//
	// 1768619049924
	RefreshTokenIssuedAt *string `json:"RefreshTokenIssuedAt,omitempty" xml:"RefreshTokenIssuedAt,omitempty"`
	// Refresh token value
	//
	// example:
	//
	// be9750d595b6cd7c93a80b46
	RefreshTokenValue *string `json:"RefreshTokenValue,omitempty" xml:"RefreshTokenValue,omitempty"`
	// Unique identifier externally
	//
	// example:
	//
	// hdm_33be9750d595b6cd7c93a80b46734b22
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateAppInstanceTicketResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceTicketResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAccessTokenExpiresAt() *string {
	return s.AccessTokenExpiresAt
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAccessTokenIssuedAt() *string {
	return s.AccessTokenIssuedAt
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAccessTokenValue() *string {
	return s.AccessTokenValue
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetAuthorizationGrantType() *string {
	return s.AuthorizationGrantType
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetBid() *string {
	return s.Bid
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetParentPk() *string {
	return s.ParentPk
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetRefreshTokenExpiresAt() *string {
	return s.RefreshTokenExpiresAt
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetRefreshTokenIssuedAt() *string {
	return s.RefreshTokenIssuedAt
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetRefreshTokenValue() *string {
	return s.RefreshTokenValue
}

func (s *CreateAppInstanceTicketResponseBodyModule) GetUuid() *string {
	return s.Uuid
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAccessTokenExpiresAt(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.AccessTokenExpiresAt = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAccessTokenIssuedAt(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.AccessTokenIssuedAt = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAccessTokenValue(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.AccessTokenValue = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAliyunPk(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.AliyunPk = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAttributes(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.Attributes = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetAuthorizationGrantType(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.AuthorizationGrantType = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetBid(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.Bid = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetParentPk(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.ParentPk = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetRefreshTokenExpiresAt(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.RefreshTokenExpiresAt = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetRefreshTokenIssuedAt(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.RefreshTokenIssuedAt = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetRefreshTokenValue(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.RefreshTokenValue = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) SetUuid(v string) *CreateAppInstanceTicketResponseBodyModule {
	s.Uuid = &v
	return s
}

func (s *CreateAppInstanceTicketResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
