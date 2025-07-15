// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportContactFlowRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFlowId(v string) *ExportContactFlowRequest
  GetFlowId() *string 
  SetInstanceId(v string) *ExportContactFlowRequest
  GetInstanceId() *string 
  SetRequestId(v string) *ExportContactFlowRequest
  GetRequestId() *string 
}

type ExportContactFlowRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // lc-uf61xdtm0mf73k
  FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 9cfad875-6260-4a53-ab6e-b13e3fb3xxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // AF9834D8-6D09-4A1B-BADB-B019D9D444C8
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportContactFlowRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportContactFlowRequest) GoString() string {
  return s.String()
}

func (s *ExportContactFlowRequest) GetFlowId() *string  {
  return s.FlowId
}

func (s *ExportContactFlowRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportContactFlowRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportContactFlowRequest) SetFlowId(v string) *ExportContactFlowRequest {
  s.FlowId = &v
  return s
}

func (s *ExportContactFlowRequest) SetInstanceId(v string) *ExportContactFlowRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportContactFlowRequest) SetRequestId(v string) *ExportContactFlowRequest {
  s.RequestId = &v
  return s
}

func (s *ExportContactFlowRequest) Validate() error {
  return dara.Validate(s)
}

