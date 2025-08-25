// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoLogoAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBoxes(v []*EraseVideoLogoAdvanceRequestBoxes) *EraseVideoLogoAdvanceRequest
  GetBoxes() []*EraseVideoLogoAdvanceRequestBoxes 
  SetVideoUrlObject(v io.Reader) *EraseVideoLogoAdvanceRequest
  GetVideoUrlObject() io.Reader 
}

type EraseVideoLogoAdvanceRequest struct {
  Boxes []*EraseVideoLogoAdvanceRequestBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoLogo/EraseVideoLogo1.mp4
  VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoAdvanceRequest) GetBoxes() []*EraseVideoLogoAdvanceRequestBoxes  {
  return s.Boxes
}

func (s *EraseVideoLogoAdvanceRequest) GetVideoUrlObject() io.Reader  {
  return s.VideoUrlObject
}

func (s *EraseVideoLogoAdvanceRequest) SetBoxes(v []*EraseVideoLogoAdvanceRequestBoxes) *EraseVideoLogoAdvanceRequest {
  s.Boxes = v
  return s
}

func (s *EraseVideoLogoAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoLogoAdvanceRequest {
  s.VideoUrlObject = v
  return s
}

func (s *EraseVideoLogoAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

type EraseVideoLogoAdvanceRequestBoxes struct {
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

func (s EraseVideoLogoAdvanceRequestBoxes) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoAdvanceRequestBoxes) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoAdvanceRequestBoxes) GetH() *float32  {
  return s.H
}

func (s *EraseVideoLogoAdvanceRequestBoxes) GetW() *float32  {
  return s.W
}

func (s *EraseVideoLogoAdvanceRequestBoxes) GetX() *float32  {
  return s.X
}

func (s *EraseVideoLogoAdvanceRequestBoxes) GetY() *float32  {
  return s.Y
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetH(v float32) *EraseVideoLogoAdvanceRequestBoxes {
  s.H = &v
  return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetW(v float32) *EraseVideoLogoAdvanceRequestBoxes {
  s.W = &v
  return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetX(v float32) *EraseVideoLogoAdvanceRequestBoxes {
  s.X = &v
  return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) SetY(v float32) *EraseVideoLogoAdvanceRequestBoxes {
  s.Y = &v
  return s
}

func (s *EraseVideoLogoAdvanceRequestBoxes) Validate() error {
  return dara.Validate(s)
}

