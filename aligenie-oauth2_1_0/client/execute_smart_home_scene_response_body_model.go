// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSmartHomeSceneResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteSmartHomeSceneResponseBody
  GetRequestId() *string 
}

type ExecuteSmartHomeSceneResponseBody struct {
  // example:
  // 
  // 435CF567-12DC-5761-AFA8-650774502E2D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSmartHomeSceneResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSmartHomeSceneResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteSmartHomeSceneResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteSmartHomeSceneResponseBody) SetRequestId(v string) *ExecuteSmartHomeSceneResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteSmartHomeSceneResponseBody) Validate() error {
  return dara.Validate(s)
}

