// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportTextScanResultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportTextScanResultResponseBody
  GetCode() *int32 
  SetData(v string) *ExportTextScanResultResponseBody
  GetData() *string 
  SetMsg(v string) *ExportTextScanResultResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportTextScanResultResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportTextScanResultResponseBody
  GetSuccess() *bool 
}

type ExportTextScanResultResponseBody struct {
  // Error code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // Exported results.
  // 
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/scanResult/text/textScanResult_aliow2MAdWXCakCxlitVY8Lnn-1A9KEw.xlsx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // Further description of the error code.
  // 
  // example:
  // 
  // OK
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
  // 
  // example:
  // 
  // AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Success indicator.
  // 
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportTextScanResultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportTextScanResultResponseBody) GoString() string {
  return s.String()
}

func (s *ExportTextScanResultResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportTextScanResultResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportTextScanResultResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportTextScanResultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportTextScanResultResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportTextScanResultResponseBody) SetCode(v int32) *ExportTextScanResultResponseBody {
  s.Code = &v
  return s
}

func (s *ExportTextScanResultResponseBody) SetData(v string) *ExportTextScanResultResponseBody {
  s.Data = &v
  return s
}

func (s *ExportTextScanResultResponseBody) SetMsg(v string) *ExportTextScanResultResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportTextScanResultResponseBody) SetRequestId(v string) *ExportTextScanResultResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportTextScanResultResponseBody) SetSuccess(v bool) *ExportTextScanResultResponseBody {
  s.Success = &v
  return s
}

func (s *ExportTextScanResultResponseBody) Validate() error {
  return dara.Validate(s)
}

