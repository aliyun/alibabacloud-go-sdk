// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceFaceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURL(v string) *EnhanceFaceRequest
  GetImageURL() *string 
}

type EnhanceFaceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/EnhanceFace/EnhanceFace1.png
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceFaceRequest) GoString() string {
  return s.String()
}

func (s *EnhanceFaceRequest) GetImageURL() *string  {
  return s.ImageURL
}

func (s *EnhanceFaceRequest) SetImageURL(v string) *EnhanceFaceRequest {
  s.ImageURL = &v
  return s
}

func (s *EnhanceFaceRequest) Validate() error {
  return dara.Validate(s)
}

