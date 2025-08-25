// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iErasePersonAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURLObject(v io.Reader) *ErasePersonAdvanceRequest
  GetImageURLObject() io.Reader 
  SetUserMaskObject(v io.Reader) *ErasePersonAdvanceRequest
  GetUserMaskObject() io.Reader 
}

type ErasePersonAdvanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ErasePerson/ErasePerson1.jpg
  ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ErasePerson/ErasePerson6.jpg
  UserMaskObject io.Reader `json:"UserMask,omitempty" xml:"UserMask,omitempty"`
}

func (s ErasePersonAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s ErasePersonAdvanceRequest) GoString() string {
  return s.String()
}

func (s *ErasePersonAdvanceRequest) GetImageURLObject() io.Reader  {
  return s.ImageURLObject
}

func (s *ErasePersonAdvanceRequest) GetUserMaskObject() io.Reader  {
  return s.UserMaskObject
}

func (s *ErasePersonAdvanceRequest) SetImageURLObject(v io.Reader) *ErasePersonAdvanceRequest {
  s.ImageURLObject = v
  return s
}

func (s *ErasePersonAdvanceRequest) SetUserMaskObject(v io.Reader) *ErasePersonAdvanceRequest {
  s.UserMaskObject = v
  return s
}

func (s *ErasePersonAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

