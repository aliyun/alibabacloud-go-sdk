// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKeywordResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportKeywordResponseBody
  GetCode() *int32 
  SetData(v string) *ExportKeywordResponseBody
  GetData() *string 
  SetMsg(v string) *ExportKeywordResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportKeywordResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportKeywordResponseBody
  GetSuccess() *bool 
}

type ExportKeywordResponseBody struct {
  // Error code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // Export result.
  // 
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/export/production/keyword/export_keywordO4ee1Bok1R8IIDVpcT9viU-1xxWr
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // Further description of the error code.
  // 
  // example:
  // 
  // OK
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
  // Request ID.
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

func (s ExportKeywordResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportKeywordResponseBody) GoString() string {
  return s.String()
}

func (s *ExportKeywordResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportKeywordResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportKeywordResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportKeywordResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportKeywordResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportKeywordResponseBody) SetCode(v int32) *ExportKeywordResponseBody {
  s.Code = &v
  return s
}

func (s *ExportKeywordResponseBody) SetData(v string) *ExportKeywordResponseBody {
  s.Data = &v
  return s
}

func (s *ExportKeywordResponseBody) SetMsg(v string) *ExportKeywordResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportKeywordResponseBody) SetRequestId(v string) *ExportKeywordResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportKeywordResponseBody) SetSuccess(v bool) *ExportKeywordResponseBody {
  s.Success = &v
  return s
}

func (s *ExportKeywordResponseBody) Validate() error {
  return dara.Validate(s)
}

