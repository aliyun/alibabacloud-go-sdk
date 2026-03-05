// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppTokenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppTokenServiceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppTokenServiceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppTokenServiceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppTokenServiceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppTokenServiceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppTokenServiceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppTokenServiceResponseBodyModule) *GetAppTokenServiceResponseBody
	GetModule() *GetAppTokenServiceResponseBodyModule
	SetRequestId(v string) *GetAppTokenServiceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppTokenServiceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppTokenServiceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppTokenServiceResponseBody
	GetSynchro() *bool
}

type GetAppTokenServiceResponseBody struct {
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
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                         `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppTokenServiceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s GetAppTokenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppTokenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppTokenServiceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppTokenServiceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppTokenServiceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppTokenServiceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppTokenServiceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppTokenServiceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppTokenServiceResponseBody) GetModule() *GetAppTokenServiceResponseBodyModule {
	return s.Module
}

func (s *GetAppTokenServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppTokenServiceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppTokenServiceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppTokenServiceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppTokenServiceResponseBody) SetAccessDeniedDetail(v string) *GetAppTokenServiceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetAllowRetry(v bool) *GetAppTokenServiceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetAppName(v string) *GetAppTokenServiceResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetDynamicCode(v string) *GetAppTokenServiceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetDynamicMessage(v string) *GetAppTokenServiceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetErrorArgs(v []interface{}) *GetAppTokenServiceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetModule(v *GetAppTokenServiceResponseBodyModule) *GetAppTokenServiceResponseBody {
	s.Module = v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetRequestId(v string) *GetAppTokenServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetRootErrorCode(v string) *GetAppTokenServiceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetRootErrorMsg(v string) *GetAppTokenServiceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) SetSynchro(v bool) *GetAppTokenServiceResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppTokenServiceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppTokenServiceResponseBodyModule struct {
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppTokenServiceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppTokenServiceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppTokenServiceResponseBodyModule) GetExtend() *string {
	return s.Extend
}

func (s *GetAppTokenServiceResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *GetAppTokenServiceResponseBodyModule) SetExtend(v string) *GetAppTokenServiceResponseBodyModule {
	s.Extend = &v
	return s
}

func (s *GetAppTokenServiceResponseBodyModule) SetStatus(v string) *GetAppTokenServiceResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetAppTokenServiceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
