// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAutoDisposeRecordsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteAutoDisposeRecordsResponseBody
  GetRequestId() *string 
}

type ExecuteAutoDisposeRecordsResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 6276D891-*****-55B2-87B9-74D413F7****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteAutoDisposeRecordsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAutoDisposeRecordsResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAutoDisposeRecordsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAutoDisposeRecordsResponseBody) SetRequestId(v string) *ExecuteAutoDisposeRecordsResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsResponseBody) Validate() error {
  return dara.Validate(s)
}

