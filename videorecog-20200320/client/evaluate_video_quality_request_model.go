// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateVideoQualityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMode(v string) *EvaluateVideoQualityRequest
  GetMode() *string 
  SetVideoUrl(v string) *EvaluateVideoQualityRequest
  GetVideoUrl() *string 
}

type EvaluateVideoQualityRequest struct {
  // example:
  // 
  // vqa_plus
  Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://public-vigen-video.oss-cn-shanghai.aliyuncs.com/Common/xxx/dont_delete/decaption/123.mp4
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EvaluateVideoQualityRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityRequest) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityRequest) GetMode() *string  {
  return s.Mode
}

func (s *EvaluateVideoQualityRequest) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EvaluateVideoQualityRequest) SetMode(v string) *EvaluateVideoQualityRequest {
  s.Mode = &v
  return s
}

func (s *EvaluateVideoQualityRequest) SetVideoUrl(v string) *EvaluateVideoQualityRequest {
  s.VideoUrl = &v
  return s
}

func (s *EvaluateVideoQualityRequest) Validate() error {
  return dara.Validate(s)
}

