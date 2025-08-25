// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateVideoQualityAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMode(v string) *EvaluateVideoQualityAdvanceRequest
  GetMode() *string 
  SetVideoUrlObject(v io.Reader) *EvaluateVideoQualityAdvanceRequest
  GetVideoUrlObject() io.Reader 
}

type EvaluateVideoQualityAdvanceRequest struct {
  // example:
  // 
  // vqa_plus
  Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // http://public-vigen-video.oss-cn-shanghai.aliyuncs.com/Common/xxx/dont_delete/decaption/123.mp4
  VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EvaluateVideoQualityAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateVideoQualityAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EvaluateVideoQualityAdvanceRequest) GetMode() *string  {
  return s.Mode
}

func (s *EvaluateVideoQualityAdvanceRequest) GetVideoUrlObject() io.Reader  {
  return s.VideoUrlObject
}

func (s *EvaluateVideoQualityAdvanceRequest) SetMode(v string) *EvaluateVideoQualityAdvanceRequest {
  s.Mode = &v
  return s
}

func (s *EvaluateVideoQualityAdvanceRequest) SetVideoUrlObject(v io.Reader) *EvaluateVideoQualityAdvanceRequest {
  s.VideoUrlObject = v
  return s
}

func (s *EvaluateVideoQualityAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

