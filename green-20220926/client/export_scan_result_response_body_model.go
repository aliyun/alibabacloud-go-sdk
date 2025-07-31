// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScanResultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportScanResultResponseBody
  GetCode() *int32 
  SetData(v string) *ExportScanResultResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportScanResultResponseBody
  GetHttpStatusCode() *int32 
  SetMsg(v string) *ExportScanResultResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportScanResultResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportScanResultResponseBody
  GetSuccess() *bool 
}

type ExportScanResultResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/scanResult/text/textScanResult_aliow2MAdWXCakCxlitVY8Lnn-1A9KEw.xlsx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // OK
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // example:
  // 
  // AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportScanResultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportScanResultResponseBody) GoString() string {
  return s.String()
}

func (s *ExportScanResultResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportScanResultResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportScanResultResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportScanResultResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportScanResultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportScanResultResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportScanResultResponseBody) SetCode(v int32) *ExportScanResultResponseBody {
  s.Code = &v
  return s
}

func (s *ExportScanResultResponseBody) SetData(v string) *ExportScanResultResponseBody {
  s.Data = &v
  return s
}

func (s *ExportScanResultResponseBody) SetHttpStatusCode(v int32) *ExportScanResultResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportScanResultResponseBody) SetMsg(v string) *ExportScanResultResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportScanResultResponseBody) SetRequestId(v string) *ExportScanResultResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportScanResultResponseBody) SetSuccess(v bool) *ExportScanResultResponseBody {
  s.Success = &v
  return s
}

func (s *ExportScanResultResponseBody) Validate() error {
  return dara.Validate(s)
}

