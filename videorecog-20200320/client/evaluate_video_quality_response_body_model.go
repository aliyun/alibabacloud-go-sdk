// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateVideoQualityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EvaluateVideoQualityResponseBodyData) *EvaluateVideoQualityResponseBody
  GetData() *EvaluateVideoQualityResponseBodyData 
  SetMessage(v string) *EvaluateVideoQualityResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EvaluateVideoQualityResponseBody
  GetRequestId() *string 
}

type EvaluateVideoQualityResponseBody struct {
  Data *EvaluateVideoQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 1d33e538-c949-4fcd-83f6-4d57e4b31527
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateVideoQualityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityResponseBody) GetData() *EvaluateVideoQualityResponseBodyData  {
  return s.Data
}

func (s *EvaluateVideoQualityResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EvaluateVideoQualityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateVideoQualityResponseBody) SetData(v *EvaluateVideoQualityResponseBodyData) *EvaluateVideoQualityResponseBody {
  s.Data = v
  return s
}

func (s *EvaluateVideoQualityResponseBody) SetMessage(v string) *EvaluateVideoQualityResponseBody {
  s.Message = &v
  return s
}

func (s *EvaluateVideoQualityResponseBody) SetRequestId(v string) *EvaluateVideoQualityResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateVideoQualityResponseBody) Validate() error {
  return dara.Validate(s)
}

type EvaluateVideoQualityResponseBodyData struct {
  // example:
  // 
  // http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/eas-video-quality-assessment/2023-01-13-10/31%3A08-cVeN9ZQlzIPfGqsa.json?Expires=1673578869&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=AiSsOsZ7rYfhf9w3Mxn%2Fq4GKKy****
  JsonUrl *string `json:"JsonUrl,omitempty" xml:"JsonUrl,omitempty"`
  // example:
  // 
  // http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/eas-video-quality-assessment/2023-01-13-10/31%3A08-cVeN9ZQlzIPfGqsa.pdf?Expires=1673578869&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=xULlZzVuhoYWAXRbp9A4EzzZcS****
  PdfUrl *string `json:"PdfUrl,omitempty" xml:"PdfUrl,omitempty"`
  VideoQualityInfo *EvaluateVideoQualityResponseBodyDataVideoQualityInfo `json:"VideoQualityInfo,omitempty" xml:"VideoQualityInfo,omitempty" type:"Struct"`
}

func (s EvaluateVideoQualityResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityResponseBodyData) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityResponseBodyData) GetJsonUrl() *string  {
  return s.JsonUrl
}

func (s *EvaluateVideoQualityResponseBodyData) GetPdfUrl() *string  {
  return s.PdfUrl
}

func (s *EvaluateVideoQualityResponseBodyData) GetVideoQualityInfo() *EvaluateVideoQualityResponseBodyDataVideoQualityInfo  {
  return s.VideoQualityInfo
}

func (s *EvaluateVideoQualityResponseBodyData) SetJsonUrl(v string) *EvaluateVideoQualityResponseBodyData {
  s.JsonUrl = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyData) SetPdfUrl(v string) *EvaluateVideoQualityResponseBodyData {
  s.PdfUrl = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyData) SetVideoQualityInfo(v *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) *EvaluateVideoQualityResponseBodyData {
  s.VideoQualityInfo = v
  return s
}

func (s *EvaluateVideoQualityResponseBodyData) Validate() error {
  return dara.Validate(s)
}

type EvaluateVideoQualityResponseBodyDataVideoQualityInfo struct {
  // example:
  // 
  // 0.15
  Blurriness *float32 `json:"Blurriness,omitempty" xml:"Blurriness,omitempty"`
  // example:
  // 
  // 0.55
  ColorContrast *float32 `json:"ColorContrast,omitempty" xml:"ColorContrast,omitempty"`
  // example:
  // 
  // 0.17
  ColorSaturation *float32 `json:"ColorSaturation,omitempty" xml:"ColorSaturation,omitempty"`
  // example:
  // 
  // 0.48
  Colorfulness *float32 `json:"Colorfulness,omitempty" xml:"Colorfulness,omitempty"`
  // example:
  // 
  // 0.25
  CompressiveStrength *float32 `json:"CompressiveStrength,omitempty" xml:"CompressiveStrength,omitempty"`
  // example:
  // 
  // 0.51
  Luminance *float32 `json:"Luminance,omitempty" xml:"Luminance,omitempty"`
  // example:
  // 
  // 0.7048
  MosScore *float32 `json:"MosScore,omitempty" xml:"MosScore,omitempty"`
  // example:
  // 
  // 0.01
  NoiseIntensity *float32 `json:"NoiseIntensity,omitempty" xml:"NoiseIntensity,omitempty"`
}

func (s EvaluateVideoQualityResponseBodyDataVideoQualityInfo) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetBlurriness() *float32  {
  return s.Blurriness
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetColorContrast() *float32  {
  return s.ColorContrast
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetColorSaturation() *float32  {
  return s.ColorSaturation
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetColorfulness() *float32  {
  return s.Colorfulness
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetCompressiveStrength() *float32  {
  return s.CompressiveStrength
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetLuminance() *float32  {
  return s.Luminance
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetMosScore() *float32  {
  return s.MosScore
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) GetNoiseIntensity() *float32  {
  return s.NoiseIntensity
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetBlurriness(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.Blurriness = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorContrast(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.ColorContrast = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorSaturation(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.ColorSaturation = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetColorfulness(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.Colorfulness = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetCompressiveStrength(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.CompressiveStrength = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetLuminance(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.Luminance = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetMosScore(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.MosScore = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) SetNoiseIntensity(v float32) *EvaluateVideoQualityResponseBodyDataVideoQualityInfo {
  s.NoiseIntensity = &v
  return s
}

func (s *EvaluateVideoQualityResponseBodyDataVideoQualityInfo) Validate() error {
  return dara.Validate(s)
}

