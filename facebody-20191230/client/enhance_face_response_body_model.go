// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceFaceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EnhanceFaceResponseBodyData) *EnhanceFaceResponseBody
  GetData() *EnhanceFaceResponseBodyData 
  SetRequestId(v string) *EnhanceFaceResponseBody
  GetRequestId() *string 
}

type EnhanceFaceResponseBody struct {
  Data *EnhanceFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 8B031473-6773-4A4C-AF33-A233758E6E98
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceFaceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnhanceFaceResponseBody) GoString() string {
  return s.String()
}

func (s *EnhanceFaceResponseBody) GetData() *EnhanceFaceResponseBodyData  {
  return s.Data
}

func (s *EnhanceFaceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnhanceFaceResponseBody) SetData(v *EnhanceFaceResponseBodyData) *EnhanceFaceResponseBody {
  s.Data = v
  return s
}

func (s *EnhanceFaceResponseBody) SetRequestId(v string) *EnhanceFaceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnhanceFaceResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnhanceFaceResponseBodyData struct {
  // example:
  // 
  // http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/face-enhancement/2021-11-30/8e34b61c-abcf-4708-9d9d-6501968ee806/20211130_080644126523_47ct9w3pgh.jpg?Expires=1638261404&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2wPdcuCmr%2F86WpBL3HQJw5uCFl****
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceFaceResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnhanceFaceResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnhanceFaceResponseBodyData) GetImageURL() *string  {
  return s.ImageURL
}

func (s *EnhanceFaceResponseBodyData) SetImageURL(v string) *EnhanceFaceResponseBodyData {
  s.ImageURL = &v
  return s
}

func (s *EnhanceFaceResponseBodyData) Validate() error {
  return dara.Validate(s)
}

