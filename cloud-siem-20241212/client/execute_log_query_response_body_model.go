// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteLogQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetQueryResult(v []interface{}) *ExecuteLogQueryResponseBody
  GetQueryResult() []interface{} 
  SetRequestId(v string) *ExecuteLogQueryResponseBody
  GetRequestId() *string 
}

type ExecuteLogQueryResponseBody struct {
  // example:
  // 
  // []。
  QueryResult []interface{} `json:"QueryResult,omitempty" xml:"QueryResult,omitempty" type:"Repeated"`
  // example:
  // 
  // 6276D891-*****-55B2-87B9-74D413F7****。
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteLogQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteLogQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteLogQueryResponseBody) GetQueryResult() []interface{}  {
  return s.QueryResult
}

func (s *ExecuteLogQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteLogQueryResponseBody) SetQueryResult(v []interface{}) *ExecuteLogQueryResponseBody {
  s.QueryResult = v
  return s
}

func (s *ExecuteLogQueryResponseBody) SetRequestId(v string) *ExecuteLogQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteLogQueryResponseBody) Validate() error {
  return dara.Validate(s)
}

