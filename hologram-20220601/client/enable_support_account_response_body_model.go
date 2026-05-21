// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSupportAccountResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v string) *EnableSupportAccountResponseBody
  GetData() *string 
  SetErrorCode(v string) *EnableSupportAccountResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableSupportAccountResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v string) *EnableSupportAccountResponseBody
  GetHttpStatusCode() *string 
  SetRequestId(v string) *EnableSupportAccountResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableSupportAccountResponseBody
  GetSuccess() *bool 
}

type EnableSupportAccountResponseBody struct {
  // example:
  // 
  // true
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // null
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // example:
  // 
  // null
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // AB71198A-2DB1-511B-AE4D-690BAA97F076
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSupportAccountResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSupportAccountResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSupportAccountResponseBody) GetData() *string  {
  return s.Data
}

func (s *EnableSupportAccountResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableSupportAccountResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableSupportAccountResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *EnableSupportAccountResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSupportAccountResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSupportAccountResponseBody) SetData(v string) *EnableSupportAccountResponseBody {
  s.Data = &v
  return s
}

func (s *EnableSupportAccountResponseBody) SetErrorCode(v string) *EnableSupportAccountResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableSupportAccountResponseBody) SetErrorMessage(v string) *EnableSupportAccountResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableSupportAccountResponseBody) SetHttpStatusCode(v string) *EnableSupportAccountResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableSupportAccountResponseBody) SetRequestId(v string) *EnableSupportAccountResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSupportAccountResponseBody) SetSuccess(v bool) *EnableSupportAccountResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSupportAccountResponseBody) Validate() error {
  return dara.Validate(s)
}

