// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceImageColorRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURL(v string) *EnhanceImageColorRequest
  GetImageURL() *string 
  SetMode(v string) *EnhanceImageColorRequest
  GetMode() *string 
  SetOutputFormat(v string) *EnhanceImageColorRequest
  GetOutputFormat() *string 
}

type EnhanceImageColorRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/EnhanceImageColor/EnhanceImageColor1.jpg
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s EnhanceImageColorRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceImageColorRequest) GoString() string {
  return s.String()
}

func (s *EnhanceImageColorRequest) GetImageURL() *string  {
  return s.ImageURL
}

func (s *EnhanceImageColorRequest) GetMode() *string  {
  return s.Mode
}

func (s *EnhanceImageColorRequest) GetOutputFormat() *string  {
  return s.OutputFormat
}

func (s *EnhanceImageColorRequest) SetImageURL(v string) *EnhanceImageColorRequest {
  s.ImageURL = &v
  return s
}

func (s *EnhanceImageColorRequest) SetMode(v string) *EnhanceImageColorRequest {
  s.Mode = &v
  return s
}

func (s *EnhanceImageColorRequest) SetOutputFormat(v string) *EnhanceImageColorRequest {
  s.OutputFormat = &v
  return s
}

func (s *EnhanceImageColorRequest) Validate() error {
  return dara.Validate(s)
}

