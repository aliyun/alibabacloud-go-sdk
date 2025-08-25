// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceVideoQualityAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBitrate(v int32) *EnhanceVideoQualityAdvanceRequest
  GetBitrate() *int32 
  SetFrameRate(v int32) *EnhanceVideoQualityAdvanceRequest
  GetFrameRate() *int32 
  SetHDRFormat(v string) *EnhanceVideoQualityAdvanceRequest
  GetHDRFormat() *string 
  SetMaxIlluminance(v int32) *EnhanceVideoQualityAdvanceRequest
  GetMaxIlluminance() *int32 
  SetOutPutHeight(v int32) *EnhanceVideoQualityAdvanceRequest
  GetOutPutHeight() *int32 
  SetOutPutWidth(v int32) *EnhanceVideoQualityAdvanceRequest
  GetOutPutWidth() *int32 
  SetVideoURLObject(v io.Reader) *EnhanceVideoQualityAdvanceRequest
  GetVideoURLObject() io.Reader 
}

type EnhanceVideoQualityAdvanceRequest struct {
  // example:
  // 
  // 20
  Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
  // example:
  // 
  // 50
  FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
  // example:
  // 
  // PQ
  HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
  // example:
  // 
  // 600
  MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
  // example:
  // 
  // 200
  OutPutHeight *int32 `json:"OutPutHeight,omitempty" xml:"OutPutHeight,omitempty"`
  // example:
  // 
  // 200
  OutPutWidth *int32 `json:"OutPutWidth,omitempty" xml:"OutPutWidth,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
  VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceVideoQualityAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EnhanceVideoQualityAdvanceRequest) GetBitrate() *int32  {
  return s.Bitrate
}

func (s *EnhanceVideoQualityAdvanceRequest) GetFrameRate() *int32  {
  return s.FrameRate
}

func (s *EnhanceVideoQualityAdvanceRequest) GetHDRFormat() *string  {
  return s.HDRFormat
}

func (s *EnhanceVideoQualityAdvanceRequest) GetMaxIlluminance() *int32  {
  return s.MaxIlluminance
}

func (s *EnhanceVideoQualityAdvanceRequest) GetOutPutHeight() *int32  {
  return s.OutPutHeight
}

func (s *EnhanceVideoQualityAdvanceRequest) GetOutPutWidth() *int32  {
  return s.OutPutWidth
}

func (s *EnhanceVideoQualityAdvanceRequest) GetVideoURLObject() io.Reader  {
  return s.VideoURLObject
}

func (s *EnhanceVideoQualityAdvanceRequest) SetBitrate(v int32) *EnhanceVideoQualityAdvanceRequest {
  s.Bitrate = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetFrameRate(v int32) *EnhanceVideoQualityAdvanceRequest {
  s.FrameRate = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetHDRFormat(v string) *EnhanceVideoQualityAdvanceRequest {
  s.HDRFormat = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetMaxIlluminance(v int32) *EnhanceVideoQualityAdvanceRequest {
  s.MaxIlluminance = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityAdvanceRequest {
  s.OutPutHeight = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityAdvanceRequest {
  s.OutPutWidth = &v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) SetVideoURLObject(v io.Reader) *EnhanceVideoQualityAdvanceRequest {
  s.VideoURLObject = v
  return s
}

func (s *EnhanceVideoQualityAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

