// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteSceneResponseBody
  GetRequestId() *string 
}

type ExecuteSceneResponseBody struct {
  // example:
  // 
  // 191C79AD-F9F9-531E-B8C1-73DF6433B920
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSceneResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSceneResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSceneResponseBody) SetRequestId(v string) *ExecuteSceneResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSceneResponseBody) Validate() error {
  return dara.Validate(s)
}

