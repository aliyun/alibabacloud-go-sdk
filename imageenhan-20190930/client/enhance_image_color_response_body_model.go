// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceImageColorResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EnhanceImageColorResponseBodyData) *EnhanceImageColorResponseBody
  GetData() *EnhanceImageColorResponseBodyData 
  SetRequestId(v string) *EnhanceImageColorResponseBody
  GetRequestId() *string 
}

type EnhanceImageColorResponseBody struct {
  Data *EnhanceImageColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // 2F306ABD-5BC3-4FA0-89CF-0DED5B3654EB
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnhanceImageColorResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnhanceImageColorResponseBody) GoString() string {
  return s.String()
}

func (s *EnhanceImageColorResponseBody) GetData() *EnhanceImageColorResponseBodyData  {
  return s.Data
}

func (s *EnhanceImageColorResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnhanceImageColorResponseBody) SetData(v *EnhanceImageColorResponseBodyData) *EnhanceImageColorResponseBody {
  s.Data = v
  return s
}

func (s *EnhanceImageColorResponseBody) SetRequestId(v string) *EnhanceImageColorResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnhanceImageColorResponseBody) Validate() error {
  return dara.Validate(s)
}

type EnhanceImageColorResponseBodyData struct {
  // example:
  // 
  // http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/image-recolor/2020-06-23-10/24%3A14-3cf26e84-a5d2-49b0-8332-0e139e20c49e.png?Expires=1592909654&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=fHrYvitvm0qZJ9nrWYa%2Fjd4pQS****
  ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EnhanceImageColorResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnhanceImageColorResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnhanceImageColorResponseBodyData) GetImageURL() *string  {
  return s.ImageURL
}

func (s *EnhanceImageColorResponseBodyData) SetImageURL(v string) *EnhanceImageColorResponseBodyData {
  s.ImageURL = &v
  return s
}

func (s *EnhanceImageColorResponseBodyData) Validate() error {
  return dara.Validate(s)
}

