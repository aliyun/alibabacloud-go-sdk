// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEdgeClusterAddEdgeMachineResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetEdgeMachineId(v string) *EdgeClusterAddEdgeMachineResponseBody
  GetEdgeMachineId() *string 
  SetRequestId(v string) *EdgeClusterAddEdgeMachineResponseBody
  GetRequestId() *string 
}

type EdgeClusterAddEdgeMachineResponseBody struct {
  // The ID of the cloud-native box.
  // 
  // example:
  // 
  // 0f4bf70a-caff-4b26-a679-fb0188a1****
  EdgeMachineId *string `json:"edge_machine_id,omitempty" xml:"edge_machine_id,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 0adf3a23-6841-41e8-9f55-7b290216c980
  RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s EdgeClusterAddEdgeMachineResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EdgeClusterAddEdgeMachineResponseBody) GoString() string {
  return s.String()
}

func (s *EdgeClusterAddEdgeMachineResponseBody) GetEdgeMachineId() *string  {
  return s.EdgeMachineId
}

func (s *EdgeClusterAddEdgeMachineResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EdgeClusterAddEdgeMachineResponseBody) SetEdgeMachineId(v string) *EdgeClusterAddEdgeMachineResponseBody {
  s.EdgeMachineId = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineResponseBody) SetRequestId(v string) *EdgeClusterAddEdgeMachineResponseBody {
  s.RequestId = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineResponseBody) Validate() error {
  return dara.Validate(s)
}

