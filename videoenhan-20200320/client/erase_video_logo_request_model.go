// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoLogoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBoxes(v []*EraseVideoLogoRequestBoxes) *EraseVideoLogoRequest
  GetBoxes() []*EraseVideoLogoRequestBoxes 
  SetVideoUrl(v string) *EraseVideoLogoRequest
  GetVideoUrl() *string 
}

type EraseVideoLogoRequest struct {
  Boxes []*EraseVideoLogoRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoLogo/EraseVideoLogo1.mp4
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoRequest) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoRequest) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoRequest) GetBoxes() []*EraseVideoLogoRequestBoxes  {
  return s.Boxes
}

func (s *EraseVideoLogoRequest) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EraseVideoLogoRequest) SetBoxes(v []*EraseVideoLogoRequestBoxes) *EraseVideoLogoRequest {
  s.Boxes = v
  return s
}

func (s *EraseVideoLogoRequest) SetVideoUrl(v string) *EraseVideoLogoRequest {
  s.VideoUrl = &v
  return s
}

func (s *EraseVideoLogoRequest) Validate() error {
  return dara.Validate(s)
}

type EraseVideoLogoRequestBoxes struct {
  // example:
  // 
  // 1.0
  H *float32 `json:"H,omitempty" xml:"H,omitempty"`
  // example:
  // 
  // 1.0
  W *float32 `json:"W,omitempty" xml:"W,omitempty"`
  // example:
  // 
  // 0.0
  X *float32 `json:"X,omitempty" xml:"X,omitempty"`
  // example:
  // 
  // 0.0
  Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s EraseVideoLogoRequestBoxes) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoRequestBoxes) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoRequestBoxes) GetH() *float32  {
  return s.H
}

func (s *EraseVideoLogoRequestBoxes) GetW() *float32  {
  return s.W
}

func (s *EraseVideoLogoRequestBoxes) GetX() *float32  {
  return s.X
}

func (s *EraseVideoLogoRequestBoxes) GetY() *float32  {
  return s.Y
}

func (s *EraseVideoLogoRequestBoxes) SetH(v float32) *EraseVideoLogoRequestBoxes {
  s.H = &v
  return s
}

func (s *EraseVideoLogoRequestBoxes) SetW(v float32) *EraseVideoLogoRequestBoxes {
  s.W = &v
  return s
}

func (s *EraseVideoLogoRequestBoxes) SetX(v float32) *EraseVideoLogoRequestBoxes {
  s.X = &v
  return s
}

func (s *EraseVideoLogoRequestBoxes) SetY(v float32) *EraseVideoLogoRequestBoxes {
  s.Y = &v
  return s
}

func (s *EraseVideoLogoRequestBoxes) Validate() error {
  return dara.Validate(s)
}

