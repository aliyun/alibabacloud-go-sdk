// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportResultResponseBody
  GetCode() *int32 
  SetData(v string) *ExportResultResponseBody
  GetData() *string 
  SetMsg(v string) *ExportResultResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportResultResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportResultResponseBody
  GetSuccess() *bool 
}

type ExportResultResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/production/scanResult/osscheck/ossCheckResult_aliiGGXhSMvmIvsS7Lrl3LaUZ-1A9%24oZ.xlsx
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s ExportResultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportResultResponseBody) GoString() string {
  return s.String()
}

func (s *ExportResultResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportResultResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportResultResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportResultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportResultResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportResultResponseBody) SetCode(v int32) *ExportResultResponseBody {
  s.Code = &v
  return s
}

func (s *ExportResultResponseBody) SetData(v string) *ExportResultResponseBody {
  s.Data = &v
  return s
}

func (s *ExportResultResponseBody) SetMsg(v string) *ExportResultResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportResultResponseBody) SetRequestId(v string) *ExportResultResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportResultResponseBody) SetSuccess(v bool) *ExportResultResponseBody {
  s.Success = &v
  return s
}

func (s *ExportResultResponseBody) Validate() error {
  return dara.Validate(s)
}

