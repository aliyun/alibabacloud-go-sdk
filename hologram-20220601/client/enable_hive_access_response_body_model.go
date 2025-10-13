// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHiveAccessResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableHiveAccessResponseBody
  GetData() *bool 
  SetErrorCode(v string) *EnableHiveAccessResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableHiveAccessResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v string) *EnableHiveAccessResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *EnableHiveAccessResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableHiveAccessResponseBody
  GetSuccess() *bool 
}

type EnableHiveAccessResponseBody struct {
  // The returned data.
  // 
  // Valid values:
  // 
  // 	- true
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // 	- false
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  //     <!-- -->
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The error code returned if the request failed.
  // 
  // example:
  // 
  // 404
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message returned.
  // 
  // example:
  // 
  // Internal server error.
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // EA8F0084-5831-5907-BB31-BD05D2617844
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHiveAccessResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHiveAccessResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHiveAccessResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableHiveAccessResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableHiveAccessResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableHiveAccessResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *EnableHiveAccessResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHiveAccessResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableHiveAccessResponseBody) SetData(v bool) *EnableHiveAccessResponseBody {
  s.Data = &v
  return s
}

func (s *EnableHiveAccessResponseBody) SetErrorCode(v string) *EnableHiveAccessResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableHiveAccessResponseBody) SetErrorMessage(v string) *EnableHiveAccessResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableHiveAccessResponseBody) SetHttpStatusCode(v string) *EnableHiveAccessResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableHiveAccessResponseBody) SetRequestId(v string) *EnableHiveAccessResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHiveAccessResponseBody) SetSuccess(v bool) *EnableHiveAccessResponseBody {
  s.Success = &v
  return s
}

func (s *EnableHiveAccessResponseBody) Validate() error {
  return dara.Validate(s)
}

