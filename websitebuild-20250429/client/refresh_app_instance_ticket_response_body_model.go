// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAppInstanceTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RefreshAppInstanceTicketResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RefreshAppInstanceTicketResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RefreshAppInstanceTicketResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RefreshAppInstanceTicketResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RefreshAppInstanceTicketResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RefreshAppInstanceTicketResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RefreshAppInstanceTicketResponseBodyModule) *RefreshAppInstanceTicketResponseBody
	GetModule() *RefreshAppInstanceTicketResponseBodyModule
	SetRequestId(v string) *RefreshAppInstanceTicketResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RefreshAppInstanceTicketResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RefreshAppInstanceTicketResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RefreshAppInstanceTicketResponseBody
	GetSynchro() *bool
}

type RefreshAppInstanceTicketResponseBody struct {
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
	// mar
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the `%s` in the **ErrMessage*	- error message.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it indicates that the input request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Returned error parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Application module
	Module *RefreshAppInstanceTicketResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s RefreshAppInstanceTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAppInstanceTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAppInstanceTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RefreshAppInstanceTicketResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RefreshAppInstanceTicketResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RefreshAppInstanceTicketResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RefreshAppInstanceTicketResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RefreshAppInstanceTicketResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RefreshAppInstanceTicketResponseBody) GetModule() *RefreshAppInstanceTicketResponseBodyModule {
	return s.Module
}

func (s *RefreshAppInstanceTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAppInstanceTicketResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RefreshAppInstanceTicketResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RefreshAppInstanceTicketResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RefreshAppInstanceTicketResponseBody) SetAccessDeniedDetail(v string) *RefreshAppInstanceTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetAllowRetry(v bool) *RefreshAppInstanceTicketResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetAppName(v string) *RefreshAppInstanceTicketResponseBody {
	s.AppName = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetDynamicCode(v string) *RefreshAppInstanceTicketResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetDynamicMessage(v string) *RefreshAppInstanceTicketResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetErrorArgs(v []interface{}) *RefreshAppInstanceTicketResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetModule(v *RefreshAppInstanceTicketResponseBodyModule) *RefreshAppInstanceTicketResponseBody {
	s.Module = v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetRequestId(v string) *RefreshAppInstanceTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetRootErrorCode(v string) *RefreshAppInstanceTicketResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetRootErrorMsg(v string) *RefreshAppInstanceTicketResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) SetSynchro(v bool) *RefreshAppInstanceTicketResponseBody {
	s.Synchro = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshAppInstanceTicketResponseBodyModule struct {
	// Access token expiration time
	//
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	AccessTokenExpiresAt *string `json:"AccessTokenExpiresAt,omitempty" xml:"AccessTokenExpiresAt,omitempty"`
	// Access token issuance time
	//
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	AccessTokenIssuedAt *string `json:"AccessTokenIssuedAt,omitempty" xml:"AccessTokenIssuedAt,omitempty"`
	// Access token value
	//
	// example:
	//
	// 2108341e17661121129745384e79f9
	AccessTokenValue *string `json:"AccessTokenValue,omitempty" xml:"AccessTokenValue,omitempty"`
	// User ID
	//
	// example:
	//
	// 1907880872137014
	AliyunPk *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// Extended attributes
	//
	// example:
	//
	// {}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Authorization grant type
	//
	// example:
	//
	// authorization_code
	AuthorizationGrantType *string `json:"AuthorizationGrantType,omitempty" xml:"AuthorizationGrantType,omitempty"`
	// business id for user
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// User ID
	//
	// example:
	//
	// 1907880872137014
	ParentPk *string `json:"ParentPk,omitempty" xml:"ParentPk,omitempty"`
	// 刷新token过期时间
	//
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	RefreshTokenExpiresAt *string `json:"RefreshTokenExpiresAt,omitempty" xml:"RefreshTokenExpiresAt,omitempty"`
	// 刷新token颁发时间
	//
	// example:
	//
	// 2025-12-18T22:30:00+08:00
	RefreshTokenIssuedAt *string `json:"RefreshTokenIssuedAt,omitempty" xml:"RefreshTokenIssuedAt,omitempty"`
	// 刷新token值
	//
	// example:
	//
	// 2108341e17661121129745384e79f9
	RefreshTokenValue *string `json:"RefreshTokenValue,omitempty" xml:"RefreshTokenValue,omitempty"`
	// Unique identifier externally
	//
	// example:
	//
	// 357504C7F21FCAE502756332ECE8B533
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RefreshAppInstanceTicketResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RefreshAppInstanceTicketResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAccessTokenExpiresAt() *string {
	return s.AccessTokenExpiresAt
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAccessTokenIssuedAt() *string {
	return s.AccessTokenIssuedAt
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAccessTokenValue() *string {
	return s.AccessTokenValue
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAttributes() *string {
	return s.Attributes
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetAuthorizationGrantType() *string {
	return s.AuthorizationGrantType
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetBid() *string {
	return s.Bid
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetParentPk() *string {
	return s.ParentPk
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetRefreshTokenExpiresAt() *string {
	return s.RefreshTokenExpiresAt
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetRefreshTokenIssuedAt() *string {
	return s.RefreshTokenIssuedAt
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetRefreshTokenValue() *string {
	return s.RefreshTokenValue
}

func (s *RefreshAppInstanceTicketResponseBodyModule) GetUuid() *string {
	return s.Uuid
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAccessTokenExpiresAt(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.AccessTokenExpiresAt = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAccessTokenIssuedAt(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.AccessTokenIssuedAt = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAccessTokenValue(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.AccessTokenValue = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAliyunPk(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.AliyunPk = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAttributes(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.Attributes = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetAuthorizationGrantType(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.AuthorizationGrantType = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetBid(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.Bid = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetParentPk(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.ParentPk = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetRefreshTokenExpiresAt(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.RefreshTokenExpiresAt = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetRefreshTokenIssuedAt(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.RefreshTokenIssuedAt = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetRefreshTokenValue(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.RefreshTokenValue = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) SetUuid(v string) *RefreshAppInstanceTicketResponseBodyModule {
	s.Uuid = &v
	return s
}

func (s *RefreshAppInstanceTicketResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
