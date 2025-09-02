// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRecognizeRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v interface{}) *EditRecognizeRuleResponseBody
  GetData() interface{} 
  SetErrorCode(v string) *EditRecognizeRuleResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EditRecognizeRuleResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v int32) *EditRecognizeRuleResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *EditRecognizeRuleResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditRecognizeRuleResponseBody
  GetSuccess() *bool 
}

type EditRecognizeRuleResponseBody struct {
  // The returned result in the JSON format.
  // 
  // example:
  // 
  // { "HttpStatusCode": 200, "Success": true }
  Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
  // The error code.
  // 
  // example:
  // 
  // 9990030003
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // Missing parameter
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 10000001
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditRecognizeRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditRecognizeRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EditRecognizeRuleResponseBody) GetData() interface{}  {
  return s.Data
}

func (s *EditRecognizeRuleResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EditRecognizeRuleResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EditRecognizeRuleResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EditRecognizeRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditRecognizeRuleResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditRecognizeRuleResponseBody) SetData(v interface{}) *EditRecognizeRuleResponseBody {
  s.Data = v
  return s
}

func (s *EditRecognizeRuleResponseBody) SetErrorCode(v string) *EditRecognizeRuleResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EditRecognizeRuleResponseBody) SetErrorMessage(v string) *EditRecognizeRuleResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EditRecognizeRuleResponseBody) SetHttpStatusCode(v int32) *EditRecognizeRuleResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EditRecognizeRuleResponseBody) SetRequestId(v string) *EditRecognizeRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditRecognizeRuleResponseBody) SetSuccess(v bool) *EditRecognizeRuleResponseBody {
  s.Success = &v
  return s
}

func (s *EditRecognizeRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

