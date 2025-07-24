// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditWorkspaceQueueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EditWorkspaceQueueResponseBody
  GetRequestId() *string 
}

type EditWorkspaceQueueResponseBody struct {
  // 请求ID。
  // 
  // example:
  // 
  // DD6B1B2A-5837-5237-ABE4-FF0C8944****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditWorkspaceQueueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditWorkspaceQueueResponseBody) GoString() string {
  return s.String()
}

func (s *EditWorkspaceQueueResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditWorkspaceQueueResponseBody) SetRequestId(v string) *EditWorkspaceQueueResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditWorkspaceQueueResponseBody) Validate() error {
  return dara.Validate(s)
}

