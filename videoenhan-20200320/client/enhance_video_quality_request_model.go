// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceVideoQualityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBitrate(v int32) *EnhanceVideoQualityRequest
  GetBitrate() *int32 
  SetFrameRate(v int32) *EnhanceVideoQualityRequest
  GetFrameRate() *int32 
  SetHDRFormat(v string) *EnhanceVideoQualityRequest
  GetHDRFormat() *string 
  SetMaxIlluminance(v int32) *EnhanceVideoQualityRequest
  GetMaxIlluminance() *int32 
  SetOutPutHeight(v int32) *EnhanceVideoQualityRequest
  GetOutPutHeight() *int32 
  SetOutPutWidth(v int32) *EnhanceVideoQualityRequest
  GetOutPutWidth() *int32 
  SetVideoURL(v string) *EnhanceVideoQualityRequest
  GetVideoURL() *string 
}

type EnhanceVideoQualityRequest struct {
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
  VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhanceVideoQualityRequest) GoString() string {
  return s.String()
}

func (s *EnhanceVideoQualityRequest) GetBitrate() *int32  {
  return s.Bitrate
}

func (s *EnhanceVideoQualityRequest) GetFrameRate() *int32  {
  return s.FrameRate
}

func (s *EnhanceVideoQualityRequest) GetHDRFormat() *string  {
  return s.HDRFormat
}

func (s *EnhanceVideoQualityRequest) GetMaxIlluminance() *int32  {
  return s.MaxIlluminance
}

func (s *EnhanceVideoQualityRequest) GetOutPutHeight() *int32  {
  return s.OutPutHeight
}

func (s *EnhanceVideoQualityRequest) GetOutPutWidth() *int32  {
  return s.OutPutWidth
}

func (s *EnhanceVideoQualityRequest) GetVideoURL() *string  {
  return s.VideoURL
}

func (s *EnhanceVideoQualityRequest) SetBitrate(v int32) *EnhanceVideoQualityRequest {
  s.Bitrate = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetFrameRate(v int32) *EnhanceVideoQualityRequest {
  s.FrameRate = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetHDRFormat(v string) *EnhanceVideoQualityRequest {
  s.HDRFormat = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetMaxIlluminance(v int32) *EnhanceVideoQualityRequest {
  s.MaxIlluminance = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutHeight(v int32) *EnhanceVideoQualityRequest {
  s.OutPutHeight = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetOutPutWidth(v int32) *EnhanceVideoQualityRequest {
  s.OutPutWidth = &v
  return s
}

func (s *EnhanceVideoQualityRequest) SetVideoURL(v string) *EnhanceVideoQualityRequest {
  s.VideoURL = &v
  return s
}

func (s *EnhanceVideoQualityRequest) Validate() error {
  return dara.Validate(s)
}

