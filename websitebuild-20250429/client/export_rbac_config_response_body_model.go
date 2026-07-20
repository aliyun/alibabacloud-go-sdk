// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRbacConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *ExportRbacConfigResponseBody
  GetAccessDeniedDetail() *string 
  SetAllowRetry(v bool) *ExportRbacConfigResponseBody
  GetAllowRetry() *bool 
  SetAppName(v string) *ExportRbacConfigResponseBody
  GetAppName() *string 
  SetDynamicCode(v string) *ExportRbacConfigResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *ExportRbacConfigResponseBody
  GetDynamicMessage() *string 
  SetErrorArgs(v []interface{}) *ExportRbacConfigResponseBody
  GetErrorArgs() []interface{} 
  SetModule(v string) *ExportRbacConfigResponseBody
  GetModule() *string 
  SetRequestId(v string) *ExportRbacConfigResponseBody
  GetRequestId() *string 
  SetRootErrorCode(v string) *ExportRbacConfigResponseBody
  GetRootErrorCode() *string 
  SetRootErrorMsg(v string) *ExportRbacConfigResponseBody
  GetRootErrorMsg() *string 
  SetSynchro(v bool) *ExportRbacConfigResponseBody
  GetSynchro() *bool 
}

type ExportRbacConfigResponseBody struct {
  // The detailed reason why access was denied.
  // 
  // example:
  // 
  // {}
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // Indicates whether retry is allowed.
  // 
  // example:
  // 
  // False
  AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
  // The application name.
  // 
  // example:
  // 
  // ish-intelligence-store-platform-admin-web
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // The dynamic error code.
  // 
  // example:
  // 
  // ERROR-oo1
  DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
  // The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- response parameter.
  // 
  // > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
  // 
  // example:
  // 
  // SYSTEM_ERROR
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  // The error parameters.
  ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
  // The response object.
  // 
  // example:
  // 
  // true
  Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 6C6B99AC-39EC-5350-874C-204128C905E6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The error code.
  // 
  // example:
  // 
  // SYSTEM.ERROR
  RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
  // The exception message.
  // 
  // example:
  // 
  // 系统异常
  RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
  // A reserved parameter.
  // 
  // example:
  // 
  // True
  Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ExportRbacConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportRbacConfigResponseBody) GoString() string {
  return s.String()
}

func (s *ExportRbacConfigResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *ExportRbacConfigResponseBody) GetAllowRetry() *bool  {
  return s.AllowRetry
}

func (s *ExportRbacConfigResponseBody) GetAppName() *string  {
  return s.AppName
}

func (s *ExportRbacConfigResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *ExportRbacConfigResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *ExportRbacConfigResponseBody) GetErrorArgs() []interface{}  {
  return s.ErrorArgs
}

func (s *ExportRbacConfigResponseBody) GetModule() *string  {
  return s.Module
}

func (s *ExportRbacConfigResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportRbacConfigResponseBody) GetRootErrorCode() *string  {
  return s.RootErrorCode
}

func (s *ExportRbacConfigResponseBody) GetRootErrorMsg() *string  {
  return s.RootErrorMsg
}

func (s *ExportRbacConfigResponseBody) GetSynchro() *bool  {
  return s.Synchro
}

func (s *ExportRbacConfigResponseBody) SetAccessDeniedDetail(v string) *ExportRbacConfigResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetAllowRetry(v bool) *ExportRbacConfigResponseBody {
  s.AllowRetry = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetAppName(v string) *ExportRbacConfigResponseBody {
  s.AppName = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetDynamicCode(v string) *ExportRbacConfigResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetDynamicMessage(v string) *ExportRbacConfigResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetErrorArgs(v []interface{}) *ExportRbacConfigResponseBody {
  s.ErrorArgs = v
  return s
}

func (s *ExportRbacConfigResponseBody) SetModule(v string) *ExportRbacConfigResponseBody {
  s.Module = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetRequestId(v string) *ExportRbacConfigResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetRootErrorCode(v string) *ExportRbacConfigResponseBody {
  s.RootErrorCode = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetRootErrorMsg(v string) *ExportRbacConfigResponseBody {
  s.RootErrorMsg = &v
  return s
}

func (s *ExportRbacConfigResponseBody) SetSynchro(v bool) *ExportRbacConfigResponseBody {
  s.Synchro = &v
  return s
}

func (s *ExportRbacConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

