// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGatewayApisResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *ExportGatewayApisResponseBody
  GetAccessDeniedDetail() *string 
  SetApiCount(v int32) *ExportGatewayApisResponseBody
  GetApiCount() *int32 
  SetDownloadUrl(v string) *ExportGatewayApisResponseBody
  GetDownloadUrl() *string 
  SetErrorField(v string) *ExportGatewayApisResponseBody
  GetErrorField() *string 
  SetGroupCount(v int32) *ExportGatewayApisResponseBody
  GetGroupCount() *int32 
  SetHttpStatusCode(v int32) *ExportGatewayApisResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *ExportGatewayApisResponseBody
  GetRequestId() *string 
  SetResultCode(v string) *ExportGatewayApisResponseBody
  GetResultCode() *string 
  SetResultMessage(v string) *ExportGatewayApisResponseBody
  GetResultMessage() *string 
  SetSuccess(v bool) *ExportGatewayApisResponseBody
  GetSuccess() *bool 
}

type ExportGatewayApisResponseBody struct {
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  ApiCount *int32 `json:"ApiCount,omitempty" xml:"ApiCount,omitempty"`
  DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
  ErrorField *string `json:"ErrorField,omitempty" xml:"ErrorField,omitempty"`
  GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
  ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportGatewayApisResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportGatewayApisResponseBody) GoString() string {
  return s.String()
}

func (s *ExportGatewayApisResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *ExportGatewayApisResponseBody) GetApiCount() *int32  {
  return s.ApiCount
}

func (s *ExportGatewayApisResponseBody) GetDownloadUrl() *string  {
  return s.DownloadUrl
}

func (s *ExportGatewayApisResponseBody) GetErrorField() *string  {
  return s.ErrorField
}

func (s *ExportGatewayApisResponseBody) GetGroupCount() *int32  {
  return s.GroupCount
}

func (s *ExportGatewayApisResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportGatewayApisResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportGatewayApisResponseBody) GetResultCode() *string  {
  return s.ResultCode
}

func (s *ExportGatewayApisResponseBody) GetResultMessage() *string  {
  return s.ResultMessage
}

func (s *ExportGatewayApisResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportGatewayApisResponseBody) SetAccessDeniedDetail(v string) *ExportGatewayApisResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetApiCount(v int32) *ExportGatewayApisResponseBody {
  s.ApiCount = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetDownloadUrl(v string) *ExportGatewayApisResponseBody {
  s.DownloadUrl = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetErrorField(v string) *ExportGatewayApisResponseBody {
  s.ErrorField = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetGroupCount(v int32) *ExportGatewayApisResponseBody {
  s.GroupCount = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetHttpStatusCode(v int32) *ExportGatewayApisResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetRequestId(v string) *ExportGatewayApisResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetResultCode(v string) *ExportGatewayApisResponseBody {
  s.ResultCode = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetResultMessage(v string) *ExportGatewayApisResponseBody {
  s.ResultMessage = &v
  return s
}

func (s *ExportGatewayApisResponseBody) SetSuccess(v bool) *ExportGatewayApisResponseBody {
  s.Success = &v
  return s
}

func (s *ExportGatewayApisResponseBody) Validate() error {
  return dara.Validate(s)
}

