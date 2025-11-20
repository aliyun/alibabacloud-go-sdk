// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhancePortraitVideoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EnhancePortraitVideoResponseBodyData) *EnhancePortraitVideoResponseBody
  GetData() *EnhancePortraitVideoResponseBodyData 
  SetMessage(v string) *EnhancePortraitVideoResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnhancePortraitVideoResponseBody
  GetRequestId() *string 
}

type EnhancePortraitVideoResponseBody struct {
  Data *EnhancePortraitVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // d21a2afa-4d52-4bca-803b-e65028146603
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhancePortraitVideoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnhancePortraitVideoResponseBody) GoString() string {
  return s.String()
}

func (s *EnhancePortraitVideoResponseBody) GetData() *EnhancePortraitVideoResponseBodyData  {
  return s.Data
}

func (s *EnhancePortraitVideoResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnhancePortraitVideoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnhancePortraitVideoResponseBody) SetData(v *EnhancePortraitVideoResponseBodyData) *EnhancePortraitVideoResponseBody {
  s.Data = v
  return s
}

func (s *EnhancePortraitVideoResponseBody) SetMessage(v string) *EnhancePortraitVideoResponseBody {
  s.Message = &v
  return s
}

func (s *EnhancePortraitVideoResponseBody) SetRequestId(v string) *EnhancePortraitVideoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnhancePortraitVideoResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnhancePortraitVideoResponseBodyData struct {
  // example:
  // 
  // http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EnhancePortraitVideoResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnhancePortraitVideoResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnhancePortraitVideoResponseBodyData) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EnhancePortraitVideoResponseBodyData) SetVideoUrl(v string) *EnhancePortraitVideoResponseBodyData {
  s.VideoUrl = &v
  return s
}

func (s *EnhancePortraitVideoResponseBodyData) Validate() error {
  return dara.Validate(s)
}

