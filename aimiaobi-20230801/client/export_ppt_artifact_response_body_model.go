// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPptArtifactResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportPptArtifactResponseBody
  GetCode() *string 
  SetData(v *ExportPptArtifactResponseBodyData) *ExportPptArtifactResponseBody
  GetData() *ExportPptArtifactResponseBodyData 
  SetHttpStatusCode(v int32) *ExportPptArtifactResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportPptArtifactResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportPptArtifactResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportPptArtifactResponseBody
  GetSuccess() *bool 
}

type ExportPptArtifactResponseBody struct {
  // example:
  // 
  // DataNotExists
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *ExportPptArtifactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 400
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // 错误消息
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // xxxxx
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportPptArtifactResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportPptArtifactResponseBody) GoString() string {
  return s.String()
}

func (s *ExportPptArtifactResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportPptArtifactResponseBody) GetData() *ExportPptArtifactResponseBodyData  {
  return s.Data
}

func (s *ExportPptArtifactResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportPptArtifactResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportPptArtifactResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportPptArtifactResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportPptArtifactResponseBody) SetCode(v string) *ExportPptArtifactResponseBody {
  s.Code = &v
  return s
}

func (s *ExportPptArtifactResponseBody) SetData(v *ExportPptArtifactResponseBodyData) *ExportPptArtifactResponseBody {
  s.Data = v
  return s
}

func (s *ExportPptArtifactResponseBody) SetHttpStatusCode(v int32) *ExportPptArtifactResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportPptArtifactResponseBody) SetMessage(v string) *ExportPptArtifactResponseBody {
  s.Message = &v
  return s
}

func (s *ExportPptArtifactResponseBody) SetRequestId(v string) *ExportPptArtifactResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportPptArtifactResponseBody) SetSuccess(v bool) *ExportPptArtifactResponseBody {
  s.Success = &v
  return s
}

func (s *ExportPptArtifactResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportPptArtifactResponseBodyData struct {
  // example:
  // 
  // adb8146b-146b-4206-bd40-19f591e85293
  ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
}

func (s ExportPptArtifactResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportPptArtifactResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportPptArtifactResponseBodyData) GetExportTaskId() *string  {
  return s.ExportTaskId
}

func (s *ExportPptArtifactResponseBodyData) SetExportTaskId(v string) *ExportPptArtifactResponseBodyData {
  s.ExportTaskId = &v
  return s
}

func (s *ExportPptArtifactResponseBodyData) Validate() error {
  return dara.Validate(s)
}

