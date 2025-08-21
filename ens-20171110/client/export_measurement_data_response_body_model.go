// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMeasurementDataResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFilePath(v string) *ExportMeasurementDataResponseBody
  GetFilePath() *string 
  SetRequestId(v string) *ExportMeasurementDataResponseBody
  GetRequestId() *string 
}

type ExportMeasurementDataResponseBody struct {
  // The download path of the exported file.
  // 
  // example:
  // 
  // http://test-oss.com/image_01.jpeg
  FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // C0003E8B-B930-4F59-ADC0-0E209A9012A8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportMeasurementDataResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportMeasurementDataResponseBody) GoString() string {
  return s.String()
}

func (s *ExportMeasurementDataResponseBody) GetFilePath() *string  {
  return s.FilePath
}

func (s *ExportMeasurementDataResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportMeasurementDataResponseBody) SetFilePath(v string) *ExportMeasurementDataResponseBody {
  s.FilePath = &v
  return s
}

func (s *ExportMeasurementDataResponseBody) SetRequestId(v string) *ExportMeasurementDataResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportMeasurementDataResponseBody) Validate() error {
  return dara.Validate(s)
}

