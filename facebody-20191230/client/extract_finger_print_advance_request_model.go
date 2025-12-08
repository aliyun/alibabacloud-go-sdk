// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractFingerPrintAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageData(v string) *ExtractFingerPrintAdvanceRequest
  GetImageData() *string 
  SetImageURLObject(v io.Reader) *ExtractFingerPrintAdvanceRequest
  GetImageURLObject() io.Reader 
}

type ExtractFingerPrintAdvanceRequest struct {
  // example:
  // 
  // iVBORw0KGgoAAAANSUhEUgAAASUAA****
  ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
  // example:
  // 
  // https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xuhan/ExtractFingerPrint.png
  ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ExtractFingerPrintAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtractFingerPrintAdvanceRequest) GoString() string {
  return s.String()
}

func (s *ExtractFingerPrintAdvanceRequest) GetImageData() *string  {
  return s.ImageData
}

func (s *ExtractFingerPrintAdvanceRequest) GetImageURLObject() io.Reader  {
  return s.ImageURLObject
}

func (s *ExtractFingerPrintAdvanceRequest) SetImageData(v string) *ExtractFingerPrintAdvanceRequest {
  s.ImageData = &v
  return s
}

func (s *ExtractFingerPrintAdvanceRequest) SetImageURLObject(v io.Reader) *ExtractFingerPrintAdvanceRequest {
  s.ImageURLObject = v
  return s
}

func (s *ExtractFingerPrintAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

