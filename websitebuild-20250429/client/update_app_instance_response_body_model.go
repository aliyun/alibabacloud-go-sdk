// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *UpdateAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *UpdateAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *UpdateAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *UpdateAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *UpdateAppInstanceResponseBodyModule) *UpdateAppInstanceResponseBody
	GetModule() *UpdateAppInstanceResponseBodyModule
	SetRequestId(v string) *UpdateAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *UpdateAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *UpdateAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *UpdateAppInstanceResponseBody
	GetSynchro() *bool
}

type UpdateAppInstanceResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                        `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *UpdateAppInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s UpdateAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *UpdateAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *UpdateAppInstanceResponseBody) GetModule() *UpdateAppInstanceResponseBodyModule {
	return s.Module
}

func (s *UpdateAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *UpdateAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *UpdateAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *UpdateAppInstanceResponseBody) SetAccessDeniedDetail(v string) *UpdateAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetAllowRetry(v bool) *UpdateAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetAppName(v string) *UpdateAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetDynamicCode(v string) *UpdateAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetDynamicMessage(v string) *UpdateAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetErrorArgs(v []interface{}) *UpdateAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetModule(v *UpdateAppInstanceResponseBodyModule) *UpdateAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetRequestId(v string) *UpdateAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetRootErrorCode(v string) *UpdateAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetRootErrorMsg(v string) *UpdateAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) SetSynchro(v bool) *UpdateAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *UpdateAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppInstanceResponseBodyModule struct {
	// example:
	//
	// WD20250703155602000001
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// example:
	//
	// 250822465990301
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// placeHolder
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s UpdateAppInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceResponseBodyModule) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppInstanceResponseBodyModule) GetIconUrl() *string {
	return s.IconUrl
}

func (s *UpdateAppInstanceResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateAppInstanceResponseBodyModule) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *UpdateAppInstanceResponseBodyModule) SetBizId(v string) *UpdateAppInstanceResponseBodyModule {
	s.BizId = &v
	return s
}

func (s *UpdateAppInstanceResponseBodyModule) SetIconUrl(v string) *UpdateAppInstanceResponseBodyModule {
	s.IconUrl = &v
	return s
}

func (s *UpdateAppInstanceResponseBodyModule) SetOrderId(v string) *UpdateAppInstanceResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *UpdateAppInstanceResponseBodyModule) SetThumbnailUrl(v string) *UpdateAppInstanceResponseBodyModule {
	s.ThumbnailUrl = &v
	return s
}

func (s *UpdateAppInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
