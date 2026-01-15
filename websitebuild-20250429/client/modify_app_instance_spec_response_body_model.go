// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyAppInstanceSpecResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ModifyAppInstanceSpecResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ModifyAppInstanceSpecResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ModifyAppInstanceSpecResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyAppInstanceSpecResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ModifyAppInstanceSpecResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *ModifyAppInstanceSpecResponseBodyModule) *ModifyAppInstanceSpecResponseBody
	GetModule() *ModifyAppInstanceSpecResponseBodyModule
	SetRequestId(v string) *ModifyAppInstanceSpecResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ModifyAppInstanceSpecResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ModifyAppInstanceSpecResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ModifyAppInstanceSpecResponseBody
	GetSynchro() *bool
}

type ModifyAppInstanceSpecResponseBody struct {
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
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// https://check-result-file-sh.oss-cn-shanghai.aliyuncs.com/u6qw3gxzu3b7sbj/u6qw3gxzu3b7sbj.diff.zip?Expires=1740975709&OSSAccessKeyId=LTAI5tKUErVCETM4ev9SELNb&Signature=FP7dDnkrLlOZHmRRORVqbLOtv9c%3D
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                            `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *ModifyAppInstanceSpecResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s ModifyAppInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceSpecResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyAppInstanceSpecResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ModifyAppInstanceSpecResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ModifyAppInstanceSpecResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyAppInstanceSpecResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyAppInstanceSpecResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ModifyAppInstanceSpecResponseBody) GetModule() *ModifyAppInstanceSpecResponseBodyModule {
	return s.Module
}

func (s *ModifyAppInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppInstanceSpecResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ModifyAppInstanceSpecResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ModifyAppInstanceSpecResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ModifyAppInstanceSpecResponseBody) SetAccessDeniedDetail(v string) *ModifyAppInstanceSpecResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetAllowRetry(v bool) *ModifyAppInstanceSpecResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetAppName(v string) *ModifyAppInstanceSpecResponseBody {
	s.AppName = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetDynamicCode(v string) *ModifyAppInstanceSpecResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetDynamicMessage(v string) *ModifyAppInstanceSpecResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetErrorArgs(v []interface{}) *ModifyAppInstanceSpecResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetModule(v *ModifyAppInstanceSpecResponseBodyModule) *ModifyAppInstanceSpecResponseBody {
	s.Module = v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetRequestId(v string) *ModifyAppInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetRootErrorCode(v string) *ModifyAppInstanceSpecResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetRootErrorMsg(v string) *ModifyAppInstanceSpecResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) SetSynchro(v bool) *ModifyAppInstanceSpecResponseBody {
	s.Synchro = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppInstanceSpecResponseBodyModule struct {
	// example:
	//
	// 247748990880615
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyAppInstanceSpecResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceSpecResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceSpecResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyAppInstanceSpecResponseBodyModule) SetOrderId(v string) *ModifyAppInstanceSpecResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *ModifyAppInstanceSpecResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
