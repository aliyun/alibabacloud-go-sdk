// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportNacosConfigResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExportNacosConfigResponseBody
  GetCode() *int32 
  SetData(v *ExportNacosConfigResponseBodyData) *ExportNacosConfigResponseBody
  GetData() *ExportNacosConfigResponseBodyData 
  SetDynamicMessage(v string) *ExportNacosConfigResponseBody
  GetDynamicMessage() *string 
  SetErrorCode(v string) *ExportNacosConfigResponseBody
  GetErrorCode() *string 
  SetHttpStatusCode(v int32) *ExportNacosConfigResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExportNacosConfigResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportNacosConfigResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportNacosConfigResponseBody
  GetSuccess() *bool 
}

type ExportNacosConfigResponseBody struct {
  // The status code returned.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The details of the data.
  Data *ExportNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The dynamic part in the error message. This parameter is used to replace **%s*	- in the **ErrMessage*	- parameter.
  // 
  // > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
  // 
  // example:
  // 
  // The specified parameter is invalid.
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  // The error code returned if the request failed.
  // 
  // example:
  // 
  // mse-100-000
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The HTTP status code returned.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The message returned.
  // 
  // example:
  // 
  // The request was successfully processed.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 92245803-49B7-54CF-8D49-01A34A0E1CD6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- `true`: The request was successful.
  // 
  // 	- `false`: The request failed.
  // 
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportNacosConfigResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportNacosConfigResponseBody) GoString() string {
  return s.String()
}

func (s *ExportNacosConfigResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExportNacosConfigResponseBody) GetData() *ExportNacosConfigResponseBodyData  {
  return s.Data
}

func (s *ExportNacosConfigResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *ExportNacosConfigResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExportNacosConfigResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExportNacosConfigResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportNacosConfigResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportNacosConfigResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportNacosConfigResponseBody) SetCode(v int32) *ExportNacosConfigResponseBody {
  s.Code = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetData(v *ExportNacosConfigResponseBodyData) *ExportNacosConfigResponseBody {
  s.Data = v
  return s
}

func (s *ExportNacosConfigResponseBody) SetDynamicMessage(v string) *ExportNacosConfigResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetErrorCode(v string) *ExportNacosConfigResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetHttpStatusCode(v int32) *ExportNacosConfigResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetMessage(v string) *ExportNacosConfigResponseBody {
  s.Message = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetRequestId(v string) *ExportNacosConfigResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportNacosConfigResponseBody) SetSuccess(v bool) *ExportNacosConfigResponseBody {
  s.Success = &v
  return s
}

func (s *ExportNacosConfigResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExportNacosConfigResponseBodyData struct {
  // The URL that is used to download the exported configurations.
  // 
  // example:
  // 
  // http://xxxxxxxxx
  Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportNacosConfigResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportNacosConfigResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportNacosConfigResponseBodyData) GetUrl() *string  {
  return s.Url
}

func (s *ExportNacosConfigResponseBodyData) SetUrl(v string) *ExportNacosConfigResponseBodyData {
  s.Url = &v
  return s
}

func (s *ExportNacosConfigResponseBodyData) Validate() error {
  return dara.Validate(s)
}

