// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhancePortraitVideoAdvanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetVideoUrlObject(v io.Reader) *EnhancePortraitVideoAdvanceRequest
  GetVideoUrlObject() io.Reader 
}

type EnhancePortraitVideoAdvanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/xxx/shang/video/SD%289516100%29.mp4
  VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoAdvanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnhancePortraitVideoAdvanceRequest) GoString() string {
  return s.String()
}

func (s *EnhancePortraitVideoAdvanceRequest) GetVideoUrlObject() io.Reader  {
  return s.VideoUrlObject
}

func (s *EnhancePortraitVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *EnhancePortraitVideoAdvanceRequest {
  s.VideoUrlObject = v
  return s
}

func (s *EnhancePortraitVideoAdvanceRequest) Validate() error {
  return dara.Validate(s)
}

