// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RenewAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *RenewAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RenewAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RenewAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RenewAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RenewAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *RenewAppInstanceResponseBodyModule) *RenewAppInstanceResponseBody
	GetModule() *RenewAppInstanceResponseBodyModule
	SetRequestId(v string) *RenewAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *RenewAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *RenewAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *RenewAppInstanceResponseBody
	GetSynchro() *bool
}

type RenewAppInstanceResponseBody struct {
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
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/gl3d6l3817id8p1/gl3d6l3817id8p1.diff.zip?Expires=1750392068&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=Bcj3eohy8nmlSQ7AAGdq7JZoLjM%3D
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                       `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *RenewAppInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s RenewAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RenewAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RenewAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RenewAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RenewAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RenewAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RenewAppInstanceResponseBody) GetModule() *RenewAppInstanceResponseBodyModule {
	return s.Module
}

func (s *RenewAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *RenewAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *RenewAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *RenewAppInstanceResponseBody) SetAccessDeniedDetail(v string) *RenewAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetAllowRetry(v bool) *RenewAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetAppName(v string) *RenewAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetDynamicCode(v string) *RenewAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetDynamicMessage(v string) *RenewAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetErrorArgs(v []interface{}) *RenewAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RenewAppInstanceResponseBody) SetModule(v *RenewAppInstanceResponseBodyModule) *RenewAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *RenewAppInstanceResponseBody) SetRequestId(v string) *RenewAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetRootErrorCode(v string) *RenewAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetRootErrorMsg(v string) *RenewAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *RenewAppInstanceResponseBody) SetSynchro(v bool) *RenewAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *RenewAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewAppInstanceResponseBodyModule struct {
	// example:
	//
	// 252126898290411
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewAppInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewAppInstanceResponseBodyModule) SetOrderId(v string) *RenewAppInstanceResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *RenewAppInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
