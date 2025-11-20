// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoLogoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EraseVideoLogoResponseBodyData) *EraseVideoLogoResponseBody
  GetData() *EraseVideoLogoResponseBodyData 
  SetMessage(v string) *EraseVideoLogoResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EraseVideoLogoResponseBody
  GetRequestId() *string 
}

type EraseVideoLogoResponseBody struct {
  Data *EraseVideoLogoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 95532F36-98FC-4DCD-815C-282BB26D2DA1
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EraseVideoLogoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoResponseBody) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoResponseBody) GetData() *EraseVideoLogoResponseBodyData  {
  return s.Data
}

func (s *EraseVideoLogoResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EraseVideoLogoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EraseVideoLogoResponseBody) SetData(v *EraseVideoLogoResponseBodyData) *EraseVideoLogoResponseBody {
  s.Data = v
  return s
}

func (s *EraseVideoLogoResponseBody) SetMessage(v string) *EraseVideoLogoResponseBody {
  s.Message = &v
  return s
}

func (s *EraseVideoLogoResponseBody) SetRequestId(v string) *EraseVideoLogoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EraseVideoLogoResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EraseVideoLogoResponseBodyData struct {
  // example:
  // 
  // http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-delogo/2020-03-20-11/53%3A56-DGNUGG7AcRlAylhr.mp4?Expires=1584707036&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=7CvsX7X1rSU%2B%2FDxnw484lb3LCD****
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoLogoResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoResponseBodyData) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoResponseBodyData) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EraseVideoLogoResponseBodyData) SetVideoUrl(v string) *EraseVideoLogoResponseBodyData {
  s.VideoUrl = &v
  return s
}

func (s *EraseVideoLogoResponseBodyData) Validate() error {
  return dara.Validate(s)
}

