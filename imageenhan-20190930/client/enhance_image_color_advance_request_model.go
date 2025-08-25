// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceImageColorAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURLObject(v io.Reader) *EnhanceImageColorAdvanceRequest
  GetImageURLObject() io.Reader 
  SetMode(v string) *EnhanceImageColorAdvanceRequest
  GetMode() *string 
  SetOutputFormat(v string) *EnhanceImageColorAdvanceRequest
  GetOutputFormat() *string 
}

type EnhanceImageColorAdvanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/EnhanceImageColor/EnhanceImageColor1.jpg
  ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // LogC
  Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // png
  OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
}

func (s EnhanceImageColorAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceImageColorAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EnhanceImageColorAdvanceRequest) GetImageURLObject() io.Reader  {
  return s.ImageURLObject
}

func (s *EnhanceImageColorAdvanceRequest) GetMode() *string  {
  return s.Mode
}

func (s *EnhanceImageColorAdvanceRequest) GetOutputFormat() *string  {
  return s.OutputFormat
}

func (s *EnhanceImageColorAdvanceRequest) SetImageURLObject(v io.Reader) *EnhanceImageColorAdvanceRequest {
  s.ImageURLObject = v
  return s
}

func (s *EnhanceImageColorAdvanceRequest) SetMode(v string) *EnhanceImageColorAdvanceRequest {
  s.Mode = &v
  return s
}

func (s *EnhanceImageColorAdvanceRequest) SetOutputFormat(v string) *EnhanceImageColorAdvanceRequest {
  s.OutputFormat = &v
  return s
}

func (s *EnhanceImageColorAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

