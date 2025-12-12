// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectZookeeperLeaderResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ElectZookeeperLeaderResponseBody
  GetRequestId() *string 
}

type ElectZookeeperLeaderResponseBody struct {
  // example:
  // 
  // 7D3ECB0E-98CA-5E08-A9CA-F70C5A1E9BDF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ElectZookeeperLeaderResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ElectZookeeperLeaderResponseBody) GoString() string {
  return s.String()
}

func (s *ElectZookeeperLeaderResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ElectZookeeperLeaderResponseBody) SetRequestId(v string) *ElectZookeeperLeaderResponseBody {
  s.RequestId = &v
  return s
}

func (s *ElectZookeeperLeaderResponseBody) Validate() error {
  return dara.Validate(s)
}

