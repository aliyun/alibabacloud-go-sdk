// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractFingerPrintRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageData(v string) *ExtractFingerPrintRequest
  GetImageData() *string 
  SetImageURL(v string) *ExtractFingerPrintRequest
  GetImageURL() *string 
}

type ExtractFingerPrintRequest struct {
  // example:
  // 
  // iVBORw0KGgoAAAANSUhEUgAAASUAA****
  ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
  // example:
  // 
  // https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xuhan/ExtractFingerPrint.png
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ExtractFingerPrintRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtractFingerPrintRequest) GoString() string {
  return s.String()
}

func (s *ExtractFingerPrintRequest) GetImageData() *string  {
  return s.ImageData
}

func (s *ExtractFingerPrintRequest) GetImageURL() *string  {
  return s.ImageURL
}

func (s *ExtractFingerPrintRequest) SetImageData(v string) *ExtractFingerPrintRequest {
  s.ImageData = &v
  return s
}

func (s *ExtractFingerPrintRequest) SetImageURL(v string) *ExtractFingerPrintRequest {
  s.ImageURL = &v
  return s
}

func (s *ExtractFingerPrintRequest) Validate() error {
  return dara.Validate(s)
}

