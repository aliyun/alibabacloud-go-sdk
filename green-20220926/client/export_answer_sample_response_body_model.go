// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnswerSampleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *ExportAnswerSampleResponseBody
  GetData() *string 
  SetRequestId(v string) *ExportAnswerSampleResponseBody
  GetRequestId() *string 
}

type ExportAnswerSampleResponseBody struct {
  // example:
  // 
  // True
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportAnswerSampleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportAnswerSampleResponseBody) GoString() string {
  return s.String()
}

func (s *ExportAnswerSampleResponseBody) GetData() *string  {
  return s.Data
}

func (s *ExportAnswerSampleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportAnswerSampleResponseBody) SetData(v string) *ExportAnswerSampleResponseBody {
  s.Data = &v
  return s
}

func (s *ExportAnswerSampleResponseBody) SetRequestId(v string) *ExportAnswerSampleResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportAnswerSampleResponseBody) Validate() error {
  return dara.Validate(s)
}

