// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPluginConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EditPluginConfigResponseBody
  GetAccessDeniedDetail() *string 
  SetAllowRetry(v bool) *EditPluginConfigResponseBody
  GetAllowRetry() *bool 
  SetAppName(v string) *EditPluginConfigResponseBody
  GetAppName() *string 
  SetDynamicCode(v string) *EditPluginConfigResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *EditPluginConfigResponseBody
  GetDynamicMessage() *string 
  SetErrorArgs(v []interface{}) *EditPluginConfigResponseBody
  GetErrorArgs() []interface{} 
  SetModule(v bool) *EditPluginConfigResponseBody
  GetModule() *bool 
  SetRequestId(v string) *EditPluginConfigResponseBody
  GetRequestId() *string 
  SetRootErrorCode(v string) *EditPluginConfigResponseBody
  GetRootErrorCode() *string 
  SetRootErrorMsg(v string) *EditPluginConfigResponseBody
  GetRootErrorMsg() *string 
  SetSynchro(v bool) *EditPluginConfigResponseBody
  GetSynchro() *bool 
}

type EditPluginConfigResponseBody struct {
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
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
  // example:
  // 
  // true
  Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
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
  // example:
  // 
  // 系统异常
  RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
  // example:
  // 
  // True
  Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s EditPluginConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditPluginConfigResponseBody) GoString() string {
  return s.String()
}

func (s *EditPluginConfigResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EditPluginConfigResponseBody) GetAllowRetry() *bool  {
  return s.AllowRetry
}

func (s *EditPluginConfigResponseBody) GetAppName() *string  {
  return s.AppName
}

func (s *EditPluginConfigResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *EditPluginConfigResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *EditPluginConfigResponseBody) GetErrorArgs() []interface{}  {
  return s.ErrorArgs
}

func (s *EditPluginConfigResponseBody) GetModule() *bool  {
  return s.Module
}

func (s *EditPluginConfigResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditPluginConfigResponseBody) GetRootErrorCode() *string  {
  return s.RootErrorCode
}

func (s *EditPluginConfigResponseBody) GetRootErrorMsg() *string  {
  return s.RootErrorMsg
}

func (s *EditPluginConfigResponseBody) GetSynchro() *bool  {
  return s.Synchro
}

func (s *EditPluginConfigResponseBody) SetAccessDeniedDetail(v string) *EditPluginConfigResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetAllowRetry(v bool) *EditPluginConfigResponseBody {
  s.AllowRetry = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetAppName(v string) *EditPluginConfigResponseBody {
  s.AppName = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetDynamicCode(v string) *EditPluginConfigResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetDynamicMessage(v string) *EditPluginConfigResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetErrorArgs(v []interface{}) *EditPluginConfigResponseBody {
  s.ErrorArgs = v
  return s
}

func (s *EditPluginConfigResponseBody) SetModule(v bool) *EditPluginConfigResponseBody {
  s.Module = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetRequestId(v string) *EditPluginConfigResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetRootErrorCode(v string) *EditPluginConfigResponseBody {
  s.RootErrorCode = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetRootErrorMsg(v string) *EditPluginConfigResponseBody {
  s.RootErrorMsg = &v
  return s
}

func (s *EditPluginConfigResponseBody) SetSynchro(v bool) *EditPluginConfigResponseBody {
  s.Synchro = &v
  return s
}

func (s *EditPluginConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

