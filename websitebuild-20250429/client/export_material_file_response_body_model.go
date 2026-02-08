// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMaterialFileResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *ExportMaterialFileResponseBody
  GetAccessDeniedDetail() *string 
  SetAllowRetry(v bool) *ExportMaterialFileResponseBody
  GetAllowRetry() *bool 
  SetAppName(v string) *ExportMaterialFileResponseBody
  GetAppName() *string 
  SetDynamicCode(v string) *ExportMaterialFileResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *ExportMaterialFileResponseBody
  GetDynamicMessage() *string 
  SetErrorArgs(v []interface{}) *ExportMaterialFileResponseBody
  GetErrorArgs() []interface{} 
  SetErrorCode(v string) *ExportMaterialFileResponseBody
  GetErrorCode() *string 
  SetErrorMsg(v string) *ExportMaterialFileResponseBody
  GetErrorMsg() *string 
  SetModule(v *ExportMaterialFileResponseBodyModule) *ExportMaterialFileResponseBody
  GetModule() *ExportMaterialFileResponseBodyModule 
  SetRequestId(v string) *ExportMaterialFileResponseBody
  GetRequestId() *string 
  SetRootErrorCode(v string) *ExportMaterialFileResponseBody
  GetRootErrorCode() *string 
  SetRootErrorMsg(v string) *ExportMaterialFileResponseBody
  GetRootErrorMsg() *string 
  SetSuccess(v bool) *ExportMaterialFileResponseBody
  GetSuccess() *bool 
  SetSynchro(v bool) *ExportMaterialFileResponseBody
  GetSynchro() *bool 
}

type ExportMaterialFileResponseBody struct {
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
  // 0
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // example:
  // 
  // aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
  ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
  Module *ExportMaterialFileResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
  RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // example:
  // 
  // True
  Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ExportMaterialFileResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportMaterialFileResponseBody) GoString() string {
  return s.String()
}

func (s *ExportMaterialFileResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *ExportMaterialFileResponseBody) GetAllowRetry() *bool  {
  return s.AllowRetry
}

func (s *ExportMaterialFileResponseBody) GetAppName() *string  {
  return s.AppName
}

func (s *ExportMaterialFileResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *ExportMaterialFileResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *ExportMaterialFileResponseBody) GetErrorArgs() []interface{}  {
  return s.ErrorArgs
}

func (s *ExportMaterialFileResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExportMaterialFileResponseBody) GetErrorMsg() *string  {
  return s.ErrorMsg
}

func (s *ExportMaterialFileResponseBody) GetModule() *ExportMaterialFileResponseBodyModule  {
  return s.Module
}

func (s *ExportMaterialFileResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportMaterialFileResponseBody) GetRootErrorCode() *string  {
  return s.RootErrorCode
}

func (s *ExportMaterialFileResponseBody) GetRootErrorMsg() *string  {
  return s.RootErrorMsg
}

func (s *ExportMaterialFileResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportMaterialFileResponseBody) GetSynchro() *bool  {
  return s.Synchro
}

func (s *ExportMaterialFileResponseBody) SetAccessDeniedDetail(v string) *ExportMaterialFileResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetAllowRetry(v bool) *ExportMaterialFileResponseBody {
  s.AllowRetry = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetAppName(v string) *ExportMaterialFileResponseBody {
  s.AppName = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetDynamicCode(v string) *ExportMaterialFileResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetDynamicMessage(v string) *ExportMaterialFileResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetErrorArgs(v []interface{}) *ExportMaterialFileResponseBody {
  s.ErrorArgs = v
  return s
}

func (s *ExportMaterialFileResponseBody) SetErrorCode(v string) *ExportMaterialFileResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetErrorMsg(v string) *ExportMaterialFileResponseBody {
  s.ErrorMsg = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetModule(v *ExportMaterialFileResponseBodyModule) *ExportMaterialFileResponseBody {
  s.Module = v
  return s
}

func (s *ExportMaterialFileResponseBody) SetRequestId(v string) *ExportMaterialFileResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetRootErrorCode(v string) *ExportMaterialFileResponseBody {
  s.RootErrorCode = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetRootErrorMsg(v string) *ExportMaterialFileResponseBody {
  s.RootErrorMsg = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetSuccess(v bool) *ExportMaterialFileResponseBody {
  s.Success = &v
  return s
}

func (s *ExportMaterialFileResponseBody) SetSynchro(v bool) *ExportMaterialFileResponseBody {
  s.Synchro = &v
  return s
}

func (s *ExportMaterialFileResponseBody) Validate() error {
  if s.Module != nil {
    if err := s.Module.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportMaterialFileResponseBodyModule struct {
  // example:
  // 
  // https://xxx.xxx.cn/original-data/pdf/mndj_report/cd649b2cc2102c0df57bfa1ae62931d6.zip
  FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ExportMaterialFileResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s ExportMaterialFileResponseBodyModule) GoString() string {
  return s.String()
}

func (s *ExportMaterialFileResponseBodyModule) GetFileUrl() *string  {
  return s.FileUrl
}

func (s *ExportMaterialFileResponseBodyModule) SetFileUrl(v string) *ExportMaterialFileResponseBodyModule {
  s.FileUrl = &v
  return s
}

func (s *ExportMaterialFileResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

