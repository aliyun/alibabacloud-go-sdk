// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendImageStyleAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMajorUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest
  GetMajorUrlObject() io.Reader 
  SetStyleUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest
  GetStyleUrlObject() io.Reader 
}

type ExtendImageStyleAdvanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ExtendImageStyle/ExtendImageStyle1.jpg
  MajorUrlObject io.Reader `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ExtendImageStyle/ExtendImageStyle6.jpg
  StyleUrlObject io.Reader `json:"StyleUrl,omitempty" xml:"StyleUrl,omitempty"`
}

func (s ExtendImageStyleAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtendImageStyleAdvanceRequest) GoString() string {
  return s.String()
}

func (s *ExtendImageStyleAdvanceRequest) GetMajorUrlObject() io.Reader  {
  return s.MajorUrlObject
}

func (s *ExtendImageStyleAdvanceRequest) GetStyleUrlObject() io.Reader  {
  return s.StyleUrlObject
}

func (s *ExtendImageStyleAdvanceRequest) SetMajorUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest {
  s.MajorUrlObject = v
  return s
}

func (s *ExtendImageStyleAdvanceRequest) SetStyleUrlObject(v io.Reader) *ExtendImageStyleAdvanceRequest {
  s.StyleUrlObject = v
  return s
}

func (s *ExtendImageStyleAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

