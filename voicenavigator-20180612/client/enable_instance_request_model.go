// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableInstanceRequest
  GetInstanceId() *string 
}

type EnableInstanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // da37319b-6c83-4268-9f19-814aed62e401
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceRequest) GoString() string {
  return s.String()
}

func (s *EnableInstanceRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInstanceRequest) SetInstanceId(v string) *EnableInstanceRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInstanceRequest) Validate() error {
  return dara.Validate(s)
}

