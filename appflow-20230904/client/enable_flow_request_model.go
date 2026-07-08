// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFlowRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFlowId(v string) *EnableFlowRequest
  GetFlowId() *string 
  SetFlowVersion(v int32) *EnableFlowRequest
  GetFlowVersion() *int32 
}

type EnableFlowRequest struct {
  // The ID of the flow.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // flow-0bf98338eb1f4d10ad6a
  FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
  // The version of the flow.
  // 
  // example:
  // 
  // 2
  FlowVersion *int32 `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s EnableFlowRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableFlowRequest) GoString() string {
  return s.String()
}

func (s *EnableFlowRequest) GetFlowId() *string  {
  return s.FlowId
}

func (s *EnableFlowRequest) GetFlowVersion() *int32  {
  return s.FlowVersion
}

func (s *EnableFlowRequest) SetFlowId(v string) *EnableFlowRequest {
  s.FlowId = &v
  return s
}

func (s *EnableFlowRequest) SetFlowVersion(v int32) *EnableFlowRequest {
  s.FlowVersion = &v
  return s
}

func (s *EnableFlowRequest) Validate() error {
  return dara.Validate(s)
}

