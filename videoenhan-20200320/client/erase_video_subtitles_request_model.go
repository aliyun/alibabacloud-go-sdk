// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoSubtitlesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBH(v float32) *EraseVideoSubtitlesRequest
  GetBH() *float32 
  SetBW(v float32) *EraseVideoSubtitlesRequest
  GetBW() *float32 
  SetBX(v float32) *EraseVideoSubtitlesRequest
  GetBX() *float32 
  SetBY(v float32) *EraseVideoSubtitlesRequest
  GetBY() *float32 
  SetVideoUrl(v string) *EraseVideoSubtitlesRequest
  GetVideoUrl() *string 
}

type EraseVideoSubtitlesRequest struct {
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
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesRequest) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoSubtitlesRequest) GoString() string {
  return s.String()
}

func (s *EraseVideoSubtitlesRequest) GetBH() *float32  {
  return s.BH
}

func (s *EraseVideoSubtitlesRequest) GetBW() *float32  {
  return s.BW
}

func (s *EraseVideoSubtitlesRequest) GetBX() *float32  {
  return s.BX
}

func (s *EraseVideoSubtitlesRequest) GetBY() *float32  {
  return s.BY
}

func (s *EraseVideoSubtitlesRequest) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EraseVideoSubtitlesRequest) SetBH(v float32) *EraseVideoSubtitlesRequest {
  s.BH = &v
  return s
}

func (s *EraseVideoSubtitlesRequest) SetBW(v float32) *EraseVideoSubtitlesRequest {
  s.BW = &v
  return s
}

func (s *EraseVideoSubtitlesRequest) SetBX(v float32) *EraseVideoSubtitlesRequest {
  s.BX = &v
  return s
}

func (s *EraseVideoSubtitlesRequest) SetBY(v float32) *EraseVideoSubtitlesRequest {
  s.BY = &v
  return s
}

func (s *EraseVideoSubtitlesRequest) SetVideoUrl(v string) *EraseVideoSubtitlesRequest {
  s.VideoUrl = &v
  return s
}

func (s *EraseVideoSubtitlesRequest) Validate() error {
  return dara.Validate(s)
}

