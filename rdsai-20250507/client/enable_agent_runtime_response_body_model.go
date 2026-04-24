// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAgentRuntimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceName(v string) *EnableAgentRuntimeResponseBody
  GetInstanceName() *string 
  SetRequestId(v string) *EnableAgentRuntimeResponseBody
  GetRequestId() *string 
}

type EnableAgentRuntimeResponseBody struct {
  // example:
  // 
  // ra-supabase-8moov5lxba****
  InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // FE9C65D7-930F-57A5-A207-8C396329****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAgentRuntimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAgentRuntimeResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAgentRuntimeResponseBody) GetInstanceName() *string  {
  return s.InstanceName
}

func (s *EnableAgentRuntimeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAgentRuntimeResponseBody) SetInstanceName(v string) *EnableAgentRuntimeResponseBody {
  s.InstanceName = &v
  return s
}

func (s *EnableAgentRuntimeResponseBody) SetRequestId(v string) *EnableAgentRuntimeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAgentRuntimeResponseBody) Validate() error {
  return dara.Validate(s)
}

