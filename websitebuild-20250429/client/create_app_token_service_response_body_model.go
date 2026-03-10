// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppTokenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateAppTokenServiceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *CreateAppTokenServiceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *CreateAppTokenServiceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *CreateAppTokenServiceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateAppTokenServiceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *CreateAppTokenServiceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *CreateAppTokenServiceResponseBodyModule) *CreateAppTokenServiceResponseBody
	GetModule() *CreateAppTokenServiceResponseBodyModule
	SetRequestId(v string) *CreateAppTokenServiceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *CreateAppTokenServiceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *CreateAppTokenServiceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *CreateAppTokenServiceResponseBody
	GetSynchro() *bool
}

type CreateAppTokenServiceResponseBody struct {
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
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                            `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *CreateAppTokenServiceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s CreateAppTokenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppTokenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppTokenServiceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateAppTokenServiceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *CreateAppTokenServiceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppTokenServiceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateAppTokenServiceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateAppTokenServiceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *CreateAppTokenServiceResponseBody) GetModule() *CreateAppTokenServiceResponseBodyModule {
	return s.Module
}

func (s *CreateAppTokenServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppTokenServiceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *CreateAppTokenServiceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *CreateAppTokenServiceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *CreateAppTokenServiceResponseBody) SetAccessDeniedDetail(v string) *CreateAppTokenServiceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetAllowRetry(v bool) *CreateAppTokenServiceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetAppName(v string) *CreateAppTokenServiceResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetDynamicCode(v string) *CreateAppTokenServiceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetDynamicMessage(v string) *CreateAppTokenServiceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetErrorArgs(v []interface{}) *CreateAppTokenServiceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetModule(v *CreateAppTokenServiceResponseBodyModule) *CreateAppTokenServiceResponseBody {
	s.Module = v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetRequestId(v string) *CreateAppTokenServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetRootErrorCode(v string) *CreateAppTokenServiceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetRootErrorMsg(v string) *CreateAppTokenServiceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) SetSynchro(v bool) *CreateAppTokenServiceResponseBody {
	s.Synchro = &v
	return s
}

func (s *CreateAppTokenServiceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppTokenServiceResponseBodyModule struct {
	// example:
	//
	// {\\"serviceApi\\":\\"sendUserMsg\\",\\"bizId\\":\\"sp\\"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s CreateAppTokenServiceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateAppTokenServiceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateAppTokenServiceResponseBodyModule) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CreateAppTokenServiceResponseBodyModule) SetExtInfo(v string) *CreateAppTokenServiceResponseBodyModule {
	s.ExtInfo = &v
	return s
}

func (s *CreateAppTokenServiceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
