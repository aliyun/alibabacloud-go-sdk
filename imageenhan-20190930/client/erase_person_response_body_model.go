// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErasePersonResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ErasePersonResponseBodyData) *ErasePersonResponseBody
  GetData() *ErasePersonResponseBodyData 
  SetRequestId(v string) *ErasePersonResponseBody
  GetRequestId() *string 
}

type ErasePersonResponseBody struct {
  Data *ErasePersonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 2FEDA495-9A5D-48B5-8922-98A4FE01D381
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ErasePersonResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ErasePersonResponseBody) GoString() string {
  return s.String()
}

func (s *ErasePersonResponseBody) GetData() *ErasePersonResponseBodyData  {
  return s.Data
}

func (s *ErasePersonResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ErasePersonResponseBody) SetData(v *ErasePersonResponseBodyData) *ErasePersonResponseBody {
  s.Data = v
  return s
}

func (s *ErasePersonResponseBody) SetRequestId(v string) *ErasePersonResponseBody {
  s.RequestId = &v
  return s
}

func (s *ErasePersonResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ErasePersonResponseBodyData struct {
  // example:
  // 
  // http://algo-app-isr-lab-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/remove-person/2020-10-29_10%3A59%3A21.421276_img19.png?Expires=1603970961&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=9lBakx0r6FOssTEYTcKs5pk8ta****
  ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ErasePersonResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ErasePersonResponseBodyData) GoString() string {
  return s.String()
}

func (s *ErasePersonResponseBodyData) GetImageUrl() *string  {
  return s.ImageUrl
}

func (s *ErasePersonResponseBodyData) SetImageUrl(v string) *ErasePersonResponseBodyData {
  s.ImageUrl = &v
  return s
}

func (s *ErasePersonResponseBodyData) Validate() error {
  return dara.Validate(s)
}

