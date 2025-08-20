// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeSetResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteChangeSetResponseBody
  GetRequestId() *string 
}

type ExecuteChangeSetResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // B288A0BE-D927-4888-B0F7-B35EF84B6E6F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteChangeSetResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeSetResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteChangeSetResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteChangeSetResponseBody) SetRequestId(v string) *ExecuteChangeSetResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteChangeSetResponseBody) Validate() error {
  return dara.Validate(s)
}

