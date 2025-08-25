// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoSubtitlesAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBH(v float32) *EraseVideoSubtitlesAdvanceRequest
  GetBH() *float32 
  SetBW(v float32) *EraseVideoSubtitlesAdvanceRequest
  GetBW() *float32 
  SetBX(v float32) *EraseVideoSubtitlesAdvanceRequest
  GetBX() *float32 
  SetBY(v float32) *EraseVideoSubtitlesAdvanceRequest
  GetBY() *float32 
  SetVideoUrlObject(v io.Reader) *EraseVideoSubtitlesAdvanceRequest
  GetVideoUrlObject() io.Reader 
}

type EraseVideoSubtitlesAdvanceRequest struct {
  // example:
  // 
  // 0.25
  BH *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
  // example:
  // 
  // 1
  BW *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
  // example:
  // 
  // 0
  BX *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
  // example:
  // 
  // 0.75
  BY *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/EraseVideoSubtitles/EraseVideoSubtitles1.mp4
  VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoSubtitlesAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EraseVideoSubtitlesAdvanceRequest) GetBH() *float32  {
  return s.BH
}

func (s *EraseVideoSubtitlesAdvanceRequest) GetBW() *float32  {
  return s.BW
}

func (s *EraseVideoSubtitlesAdvanceRequest) GetBX() *float32  {
  return s.BX
}

func (s *EraseVideoSubtitlesAdvanceRequest) GetBY() *float32  {
  return s.BY
}

func (s *EraseVideoSubtitlesAdvanceRequest) GetVideoUrlObject() io.Reader  {
  return s.VideoUrlObject
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBH(v float32) *EraseVideoSubtitlesAdvanceRequest {
  s.BH = &v
  return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBW(v float32) *EraseVideoSubtitlesAdvanceRequest {
  s.BW = &v
  return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBX(v float32) *EraseVideoSubtitlesAdvanceRequest {
  s.BX = &v
  return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetBY(v float32) *EraseVideoSubtitlesAdvanceRequest {
  s.BY = &v
  return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) SetVideoUrlObject(v io.Reader) *EraseVideoSubtitlesAdvanceRequest {
  s.VideoUrlObject = v
  return s
}

func (s *EraseVideoSubtitlesAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

