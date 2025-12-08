// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceFaceAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURLObject(v io.Reader) *EnhanceFaceAdvanceRequest
  GetImageURLObject() io.Reader 
}

type EnhanceFaceAdvanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/EnhanceFace/EnhanceFace1.png
  ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceFaceAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EnhanceFaceAdvanceRequest) GetImageURLObject() io.Reader  {
  return s.ImageURLObject
}

func (s *EnhanceFaceAdvanceRequest) SetImageURLObject(v io.Reader) *EnhanceFaceAdvanceRequest {
  s.ImageURLObject = v
  return s
}

func (s *EnhanceFaceAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

