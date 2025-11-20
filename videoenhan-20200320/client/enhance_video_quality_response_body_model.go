// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceVideoQualityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EnhanceVideoQualityResponseBodyData) *EnhanceVideoQualityResponseBody
  GetData() *EnhanceVideoQualityResponseBodyData 
  SetMessage(v string) *EnhanceVideoQualityResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnhanceVideoQualityResponseBody
  GetRequestId() *string 
}

type EnhanceVideoQualityResponseBody struct {
  Data *EnhanceVideoQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 881F39DC-C107-4817-A6D5-000BE833CC2A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceVideoQualityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnhanceVideoQualityResponseBody) GoString() string {
  return s.String()
}

func (s *EnhanceVideoQualityResponseBody) GetData() *EnhanceVideoQualityResponseBodyData  {
  return s.Data
}

func (s *EnhanceVideoQualityResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnhanceVideoQualityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnhanceVideoQualityResponseBody) SetData(v *EnhanceVideoQualityResponseBodyData) *EnhanceVideoQualityResponseBody {
  s.Data = v
  return s
}

func (s *EnhanceVideoQualityResponseBody) SetMessage(v string) *EnhanceVideoQualityResponseBody {
  s.Message = &v
  return s
}

func (s *EnhanceVideoQualityResponseBody) SetRequestId(v string) *EnhanceVideoQualityResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnhanceVideoQualityResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnhanceVideoQualityResponseBodyData struct {
  // example:
  // 
  // http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/uhd-enhance/20-11-20/Wwzf9z75GO5XdisS_20-11-20-07-13-48.mp4?Expires=1605858272&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=vvY0D%2Bl5eEzp%2BD7mPOWz0zMU7v****
  VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EnhanceVideoQualityResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnhanceVideoQualityResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnhanceVideoQualityResponseBodyData) GetVideoURL() *string  {
  return s.VideoURL
}

func (s *EnhanceVideoQualityResponseBodyData) SetVideoURL(v string) *EnhanceVideoQualityResponseBodyData {
  s.VideoURL = &v
  return s
}

func (s *EnhanceVideoQualityResponseBodyData) Validate() error {
  return dara.Validate(s)
}

