// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErasePersonRequest interface {
  dara.Model
  String() string
  GoString() string
  SetImageURL(v string) *ErasePersonRequest
  GetImageURL() *string 
  SetUserMask(v string) *ErasePersonRequest
  GetUserMask() *string 
}

type ErasePersonRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ErasePerson/ErasePerson1.jpg
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ErasePerson/ErasePerson6.jpg
  UserMask *string `json:"UserMask,omitempty" xml:"UserMask,omitempty"`
}

func (s ErasePersonRequest) String() string {
  return dara.Prettify(s)
}

func (s ErasePersonRequest) GoString() string {
  return s.String()
}

func (s *ErasePersonRequest) GetImageURL() *string  {
  return s.ImageURL
}

func (s *ErasePersonRequest) GetUserMask() *string  {
  return s.UserMask
}

func (s *ErasePersonRequest) SetImageURL(v string) *ErasePersonRequest {
  s.ImageURL = &v
  return s
}

func (s *ErasePersonRequest) SetUserMask(v string) *ErasePersonRequest {
  s.UserMask = &v
  return s
}

func (s *ErasePersonRequest) Validate() error {
  return dara.Validate(s)
}

