// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCipStatsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportCipStatsResponseBody
  GetCode() *int32 
  SetData(v string) *ExportCipStatsResponseBody
  GetData() *string 
  SetHttpStatusCode(v int32) *ExportCipStatsResponseBody
  GetHttpStatusCode() *int32 
  SetMsg(v string) *ExportCipStatsResponseBody
  GetMsg() *string 
  SetRequestId(v string) *ExportCipStatsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportCipStatsResponseBody
  GetSuccess() *bool 
}

type ExportCipStatsResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/cipStat/text/statistics1720597246783.xlsx
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

func (s ExportCipStatsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportCipStatsResponseBody) GoString() string {
  return s.String()
}

func (s *ExportCipStatsResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportCipStatsResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportCipStatsResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportCipStatsResponseBody) GetMsg() *string  {
  return s.Msg
}

func (s *ExportCipStatsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportCipStatsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportCipStatsResponseBody) SetCode(v int32) *ExportCipStatsResponseBody {
  s.Code = &v
  return s
}

func (s *ExportCipStatsResponseBody) SetData(v string) *ExportCipStatsResponseBody {
  s.Data = &v
  return s
}

func (s *ExportCipStatsResponseBody) SetHttpStatusCode(v int32) *ExportCipStatsResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportCipStatsResponseBody) SetMsg(v string) *ExportCipStatsResponseBody {
  s.Msg = &v
  return s
}

func (s *ExportCipStatsResponseBody) SetRequestId(v string) *ExportCipStatsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportCipStatsResponseBody) SetSuccess(v bool) *ExportCipStatsResponseBody {
  s.Success = &v
  return s
}

func (s *ExportCipStatsResponseBody) Validate() error {
  return dara.Validate(s)
}

