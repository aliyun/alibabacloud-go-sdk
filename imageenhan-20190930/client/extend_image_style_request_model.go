// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendImageStyleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMajorUrl(v string) *ExtendImageStyleRequest
  GetMajorUrl() *string 
  SetStyleUrl(v string) *ExtendImageStyleRequest
  GetStyleUrl() *string 
}

type ExtendImageStyleRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ExtendImageStyle/ExtendImageStyle1.jpg
  MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ExtendImageStyle/ExtendImageStyle6.jpg
  StyleUrl *string `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ExtendImageStyleRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtendImageStyleRequest) GoString() string {
  return s.String()
}

func (s *ExtendImageStyleRequest) GetMajorUrl() *string  {
  return s.MajorUrl
}

func (s *ExtendImageStyleRequest) GetStyleUrl() *string  {
  return s.StyleUrl
}

func (s *ExtendImageStyleRequest) SetMajorUrl(v string) *ExtendImageStyleRequest {
  s.MajorUrl = &v
  return s
}

func (s *ExtendImageStyleRequest) SetStyleUrl(v string) *ExtendImageStyleRequest {
  s.StyleUrl = &v
  return s
}

func (s *ExtendImageStyleRequest) Validate() error {
  return dara.Validate(s)
}

