// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PublishAppInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *PublishAppInstanceResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PublishAppInstanceResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PublishAppInstanceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PublishAppInstanceResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *PublishAppInstanceResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *PublishAppInstanceResponseBodyModule) *PublishAppInstanceResponseBody
	GetModule() *PublishAppInstanceResponseBodyModule
	SetRequestId(v string) *PublishAppInstanceResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *PublishAppInstanceResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *PublishAppInstanceResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *PublishAppInstanceResponseBody
	GetSynchro() *bool
}

type PublishAppInstanceResponseBody struct {
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
	// SYSTEM_ERROR
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                         `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *PublishAppInstanceResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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

func (s PublishAppInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAppInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PublishAppInstanceResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PublishAppInstanceResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PublishAppInstanceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PublishAppInstanceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PublishAppInstanceResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *PublishAppInstanceResponseBody) GetModule() *PublishAppInstanceResponseBodyModule {
	return s.Module
}

func (s *PublishAppInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishAppInstanceResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *PublishAppInstanceResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *PublishAppInstanceResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *PublishAppInstanceResponseBody) SetAccessDeniedDetail(v string) *PublishAppInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetAllowRetry(v bool) *PublishAppInstanceResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetAppName(v string) *PublishAppInstanceResponseBody {
	s.AppName = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetDynamicCode(v string) *PublishAppInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetDynamicMessage(v string) *PublishAppInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetErrorArgs(v []interface{}) *PublishAppInstanceResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *PublishAppInstanceResponseBody) SetModule(v *PublishAppInstanceResponseBodyModule) *PublishAppInstanceResponseBody {
	s.Module = v
	return s
}

func (s *PublishAppInstanceResponseBody) SetRequestId(v string) *PublishAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetRootErrorCode(v string) *PublishAppInstanceResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetRootErrorMsg(v string) *PublishAppInstanceResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *PublishAppInstanceResponseBody) SetSynchro(v bool) *PublishAppInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *PublishAppInstanceResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishAppInstanceResponseBodyModule struct {
	// example:
	//
	// 231
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
}

func (s PublishAppInstanceResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s PublishAppInstanceResponseBodyModule) GoString() string {
	return s.String()
}

func (s *PublishAppInstanceResponseBodyModule) GetPublishOrderId() *int64 {
	return s.PublishOrderId
}

func (s *PublishAppInstanceResponseBodyModule) SetPublishOrderId(v int64) *PublishAppInstanceResponseBodyModule {
	s.PublishOrderId = &v
	return s
}

func (s *PublishAppInstanceResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
