// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEdgeClusterAddEdgeMachineRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExpired(v int64) *EdgeClusterAddEdgeMachineRequest
  GetExpired() *int64 
  SetNodepoolId(v string) *EdgeClusterAddEdgeMachineRequest
  GetNodepoolId() *string 
  SetOptions(v string) *EdgeClusterAddEdgeMachineRequest
  GetOptions() *string 
}

type EdgeClusterAddEdgeMachineRequest struct {
  // The timeout period of sessions. Unit: seconds.
  // 
  // example:
  // 
  // 1024
  Expired *int64 `json:"expired,omitempty" xml:"expired,omitempty"`
  // The node pool ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // c26607f52179f4472a0d9723e7595****
  NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
  // The options that you want to configure.
  // 
  // example:
  // 
  // "{\\"enableIptables\\":true,\\"quiet\\":true,\\"manageRuntime\\":true,\\"allowedClusterAddons\\":[\\"kube-proxy\\",\\"flannel\\",\\"coredns\\"]}"
  Options *string `json:"options,omitempty" xml:"options,omitempty"`
}

func (s EdgeClusterAddEdgeMachineRequest) String() string {
  return dara.Prettify(s)
}

func (s EdgeClusterAddEdgeMachineRequest) GoString() string {
  return s.String()
}

func (s *EdgeClusterAddEdgeMachineRequest) GetExpired() *int64  {
  return s.Expired
}

func (s *EdgeClusterAddEdgeMachineRequest) GetNodepoolId() *string  {
  return s.NodepoolId
}

func (s *EdgeClusterAddEdgeMachineRequest) GetOptions() *string  {
  return s.Options
}

func (s *EdgeClusterAddEdgeMachineRequest) SetExpired(v int64) *EdgeClusterAddEdgeMachineRequest {
  s.Expired = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineRequest) SetNodepoolId(v string) *EdgeClusterAddEdgeMachineRequest {
  s.NodepoolId = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineRequest) SetOptions(v string) *EdgeClusterAddEdgeMachineRequest {
  s.Options = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineRequest) Validate() error {
  return dara.Validate(s)
}

