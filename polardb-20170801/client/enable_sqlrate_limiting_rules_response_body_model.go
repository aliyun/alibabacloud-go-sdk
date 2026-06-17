// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSQLRateLimitingRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMessage(v string) *EnableSQLRateLimitingRulesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableSQLRateLimitingRulesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableSQLRateLimitingRulesResponseBody
  GetSuccess() *bool 
}

type EnableSQLRateLimitingRulesResponseBody struct {
  // The response message.
  // 
  // > If the request is successful, `Successful` is returned. If the request fails, an error message is returned.
  // 
  // example:
  // 
  // Successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 4CE6DF97-AEA4-484F-906F-C407EE******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // - **true**: The request was successful.
  // 
  // - **false**: The request failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSQLRateLimitingRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSQLRateLimitingRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSQLRateLimitingRulesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableSQLRateLimitingRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSQLRateLimitingRulesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSQLRateLimitingRulesResponseBody) SetMessage(v string) *EnableSQLRateLimitingRulesResponseBody {
  s.Message = &v
  return s
}

func (s *EnableSQLRateLimitingRulesResponseBody) SetRequestId(v string) *EnableSQLRateLimitingRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSQLRateLimitingRulesResponseBody) SetSuccess(v bool) *EnableSQLRateLimitingRulesResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSQLRateLimitingRulesResponseBody) Validate() error {
  return dara.Validate(s)
}

