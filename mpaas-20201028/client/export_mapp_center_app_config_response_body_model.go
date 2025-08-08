// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMappCenterAppConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetExportMappCenterAppConfigResult(v *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) *ExportMappCenterAppConfigResponseBody
  GetExportMappCenterAppConfigResult() *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult 
  SetRequestId(v string) *ExportMappCenterAppConfigResponseBody
  GetRequestId() *string 
  SetResultCode(v string) *ExportMappCenterAppConfigResponseBody
  GetResultCode() *string 
  SetResultMessage(v string) *ExportMappCenterAppConfigResponseBody
  GetResultMessage() *string 
}

type ExportMappCenterAppConfigResponseBody struct {
  ExportMappCenterAppConfigResult *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult `json:"ExportMappCenterAppConfigResult,omitempty" xml:"ExportMappCenterAppConfigResult,omitempty" type:"Struct"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
  ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ExportMappCenterAppConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportMappCenterAppConfigResponseBody) GoString() string {
  return s.String()
}

func (s *ExportMappCenterAppConfigResponseBody) GetExportMappCenterAppConfigResult() *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult  {
  return s.ExportMappCenterAppConfigResult
}

func (s *ExportMappCenterAppConfigResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportMappCenterAppConfigResponseBody) GetResultCode() *string  {
  return s.ResultCode
}

func (s *ExportMappCenterAppConfigResponseBody) GetResultMessage() *string  {
  return s.ResultMessage
}

func (s *ExportMappCenterAppConfigResponseBody) SetExportMappCenterAppConfigResult(v *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) *ExportMappCenterAppConfigResponseBody {
  s.ExportMappCenterAppConfigResult = v
  return s
}

func (s *ExportMappCenterAppConfigResponseBody) SetRequestId(v string) *ExportMappCenterAppConfigResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBody) SetResultCode(v string) *ExportMappCenterAppConfigResponseBody {
  s.ResultCode = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBody) SetResultMessage(v string) *ExportMappCenterAppConfigResponseBody {
  s.ResultMessage = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult struct {
  ConfigDownloadUrl *string `json:"ConfigDownloadUrl,omitempty" xml:"ConfigDownloadUrl,omitempty"`
  ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) String() string {
  return dara.Prettify(s)
}

func (s ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) GoString() string {
  return s.String()
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) GetConfigDownloadUrl() *string  {
  return s.ConfigDownloadUrl
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) GetResultMsg() *string  {
  return s.ResultMsg
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) SetConfigDownloadUrl(v string) *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult {
  s.ConfigDownloadUrl = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) SetResultMsg(v string) *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult {
  s.ResultMsg = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) SetSuccess(v bool) *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult {
  s.Success = &v
  return s
}

func (s *ExportMappCenterAppConfigResponseBodyExportMappCenterAppConfigResult) Validate() error {
  return dara.Validate(s)
}

