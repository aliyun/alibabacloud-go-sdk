// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendImageStyleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExtendImageStyleResponseBodyData) *ExtendImageStyleResponseBody
  GetData() *ExtendImageStyleResponseBodyData 
  SetRequestId(v string) *ExtendImageStyleResponseBody
  GetRequestId() *string 
}

type ExtendImageStyleResponseBody struct {
  Data *ExtendImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // F1ABC965-2612-4680-9DE3-B5C77434B9B7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExtendImageStyleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExtendImageStyleResponseBody) GoString() string {
  return s.String()
}

func (s *ExtendImageStyleResponseBody) GetData() *ExtendImageStyleResponseBodyData  {
  return s.Data
}

func (s *ExtendImageStyleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExtendImageStyleResponseBody) SetData(v *ExtendImageStyleResponseBodyData) *ExtendImageStyleResponseBody {
  s.Data = v
  return s
}

func (s *ExtendImageStyleResponseBody) SetRequestId(v string) *ExtendImageStyleResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExtendImageStyleResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExtendImageStyleResponseBodyData struct {
  // example:
  // 
  // -
  MajorUrl *string `json:"MajorUrl,omitempty" xml:"MajorUrl,omitempty"`
  // example:
  // 
  // http://luban-vgd-invi.oss-cn-hangzhou.aliyuncs.com/upload/result_/2019-9-26/invi__015694889247201016812_spCq4n.jpg?Expires=1569492524&OSSAccessKeyId=LTAI4Fc5SVvzUQ19K1Cz****&Signature=uOEP8gKvdgPtDcrXxRz1v37dsT****
  Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExtendImageStyleResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExtendImageStyleResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExtendImageStyleResponseBodyData) GetMajorUrl() *string  {
  return s.MajorUrl
}

func (s *ExtendImageStyleResponseBodyData) GetUrl() *string  {
  return s.Url
}

func (s *ExtendImageStyleResponseBodyData) SetMajorUrl(v string) *ExtendImageStyleResponseBodyData {
  s.MajorUrl = &v
  return s
}

func (s *ExtendImageStyleResponseBodyData) SetUrl(v string) *ExtendImageStyleResponseBodyData {
  s.Url = &v
  return s
}

func (s *ExtendImageStyleResponseBodyData) Validate() error {
  return dara.Validate(s)
}

