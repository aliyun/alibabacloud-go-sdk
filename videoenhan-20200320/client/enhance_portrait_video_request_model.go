// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhancePortraitVideoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetVideoUrl(v string) *EnhancePortraitVideoRequest
  GetVideoUrl() *string 
}

type EnhancePortraitVideoRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhancePortraitVideoRequest) GoString() string {
  return s.String()
}

func (s *EnhancePortraitVideoRequest) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EnhancePortraitVideoRequest) SetVideoUrl(v string) *EnhancePortraitVideoRequest {
  s.VideoUrl = &v
  return s
}

func (s *EnhancePortraitVideoRequest) Validate() error {
  return dara.Validate(s)
}

