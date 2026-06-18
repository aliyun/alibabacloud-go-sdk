// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApiKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableApiKeyResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *EnableApiKeyResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableApiKeyResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableApiKeyResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableApiKeyResponseBody
  GetSuccess() *bool 
}

type EnableApiKeyResponseBody struct {
  // The status code.
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // The response message.
  // 
  // example:
  // 
  // ok
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // BB521414-5D38-5E66-AA66-963B2B4200E2
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the API call was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EnableApiKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApiKeyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApiKeyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableApiKeyResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableApiKeyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableApiKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApiKeyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableApiKeyResponseBody) SetCode(v string) *EnableApiKeyResponseBody {
  s.Code = &v
  return s
}

func (s *EnableApiKeyResponseBody) SetHttpStatusCode(v int32) *EnableApiKeyResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableApiKeyResponseBody) SetMessage(v string) *EnableApiKeyResponseBody {
  s.Message = &v
  return s
}

func (s *EnableApiKeyResponseBody) SetRequestId(v string) *EnableApiKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApiKeyResponseBody) SetSuccess(v bool) *EnableApiKeyResponseBody {
  s.Success = &v
  return s
}

func (s *EnableApiKeyResponseBody) Validate() error {
  return dara.Validate(s)
}

