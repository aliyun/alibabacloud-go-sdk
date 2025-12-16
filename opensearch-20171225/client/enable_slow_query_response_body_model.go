// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSlowQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSlowQueryResponseBody
  GetRequestId() *string 
  SetResult(v map[string]interface{}) *EnableSlowQueryResponseBody
  GetResult() map[string]interface{} 
}

type EnableSlowQueryResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 728E89C6-8673-D39B-39A1-5BA2B56D448F
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // The returned data.
  // 
  // example:
  // 
  // {}
  Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EnableSlowQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSlowQueryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSlowQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSlowQueryResponseBody) GetResult() map[string]interface{}  {
  return s.Result
}

func (s *EnableSlowQueryResponseBody) SetRequestId(v string) *EnableSlowQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSlowQueryResponseBody) SetResult(v map[string]interface{}) *EnableSlowQueryResponseBody {
  s.Result = v
  return s
}

func (s *EnableSlowQueryResponseBody) Validate() error {
  return dara.Validate(s)
}

