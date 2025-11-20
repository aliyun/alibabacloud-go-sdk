// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoSubtitlesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EraseVideoSubtitlesResponseBodyData) *EraseVideoSubtitlesResponseBody
  GetData() *EraseVideoSubtitlesResponseBodyData 
  SetMessage(v string) *EraseVideoSubtitlesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EraseVideoSubtitlesResponseBody
  GetRequestId() *string 
}

type EraseVideoSubtitlesResponseBody struct {
  Data *EraseVideoSubtitlesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // CCB082BF-A6B1-4C28-9E49-562EEE7DE639
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EraseVideoSubtitlesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoSubtitlesResponseBody) GoString() string {
  return s.String()
}

func (s *EraseVideoSubtitlesResponseBody) GetData() *EraseVideoSubtitlesResponseBodyData  {
  return s.Data
}

func (s *EraseVideoSubtitlesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EraseVideoSubtitlesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EraseVideoSubtitlesResponseBody) SetData(v *EraseVideoSubtitlesResponseBodyData) *EraseVideoSubtitlesResponseBody {
  s.Data = v
  return s
}

func (s *EraseVideoSubtitlesResponseBody) SetMessage(v string) *EraseVideoSubtitlesResponseBody {
  s.Message = &v
  return s
}

func (s *EraseVideoSubtitlesResponseBody) SetRequestId(v string) *EraseVideoSubtitlesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EraseVideoSubtitlesResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EraseVideoSubtitlesResponseBodyData struct {
  // example:
  // 
  // http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-desubtitle/2021-04-13-10/41%3A57-TcFd6Zug7gXwbeqs.mp4?Expires=1618312317&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=LZnGSQ8019%2Br5rcR4vKOaaT2UE****
  VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s EraseVideoSubtitlesResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoSubtitlesResponseBodyData) GoString() string {
  return s.String()
}

func (s *EraseVideoSubtitlesResponseBodyData) GetVideoUrl() *string  {
  return s.VideoUrl
}

func (s *EraseVideoSubtitlesResponseBodyData) SetVideoUrl(v string) *EraseVideoSubtitlesResponseBodyData {
  s.VideoUrl = &v
  return s
}

func (s *EraseVideoSubtitlesResponseBodyData) Validate() error {
  return dara.Validate(s)
}

