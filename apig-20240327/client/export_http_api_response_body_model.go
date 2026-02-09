// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHttpApiResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExportHttpApiResponseBody
  GetCode() *string 
  SetData(v *ExportHttpApiResponseBodyData) *ExportHttpApiResponseBody
  GetData() *ExportHttpApiResponseBodyData 
  SetMessage(v string) *ExportHttpApiResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportHttpApiResponseBody
  GetRequestId() *string 
}

type ExportHttpApiResponseBody struct {
  // The status code.
  // 
  // example:
  // 
  // Ok
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The API definition.
  Data *ExportHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // The response message returned.
  // 
  // example:
  // 
  // success
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 4BACB05C-3FE2-588F-9148-700C5C026B74
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExportHttpApiResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportHttpApiResponseBody) GoString() string {
  return s.String()
}

func (s *ExportHttpApiResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExportHttpApiResponseBody) GetData() *ExportHttpApiResponseBodyData  {
  return s.Data
}

func (s *ExportHttpApiResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportHttpApiResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportHttpApiResponseBody) SetCode(v string) *ExportHttpApiResponseBody {
  s.Code = &v
  return s
}

func (s *ExportHttpApiResponseBody) SetData(v *ExportHttpApiResponseBodyData) *ExportHttpApiResponseBody {
  s.Data = v
  return s
}

func (s *ExportHttpApiResponseBody) SetMessage(v string) *ExportHttpApiResponseBody {
  s.Message = &v
  return s
}

func (s *ExportHttpApiResponseBody) SetRequestId(v string) *ExportHttpApiResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportHttpApiResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportHttpApiResponseBodyData struct {
  // The Base64-encoded API definition.
  // 
  // example:
  // 
  // b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
  SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
}

func (s ExportHttpApiResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportHttpApiResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportHttpApiResponseBodyData) GetSpecContentBase64() *string  {
  return s.SpecContentBase64
}

func (s *ExportHttpApiResponseBodyData) SetSpecContentBase64(v string) *ExportHttpApiResponseBodyData {
  s.SpecContentBase64 = &v
  return s
}

func (s *ExportHttpApiResponseBodyData) Validate() error {
  return dara.Validate(s)
}

