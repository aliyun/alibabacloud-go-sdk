// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndCoordinationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCoordinatorUserId(v string) *EndCoordinationRequest
  GetCoordinatorUserId() *string 
  SetInstanceId(v string) *EndCoordinationRequest
  GetInstanceId() *string 
  SetOwnerUserId(v string) *EndCoordinationRequest
  GetOwnerUserId() *string 
}

type EndCoordinationRequest struct {
  // example:
  // 
  // lina
  CoordinatorUserId *string `json:"CoordinatorUserId,omitempty" xml:"CoordinatorUserId,omitempty"`
  // example:
  // 
  // acp-2zecay9ponatdc4m****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // xiaoming
  OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s EndCoordinationRequest) String() string {
  return dara.Prettify(s)
}

func (s EndCoordinationRequest) GoString() string {
  return s.String()
}

func (s *EndCoordinationRequest) GetCoordinatorUserId() *string  {
  return s.CoordinatorUserId
}

func (s *EndCoordinationRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndCoordinationRequest) GetOwnerUserId() *string  {
  return s.OwnerUserId
}

func (s *EndCoordinationRequest) SetCoordinatorUserId(v string) *EndCoordinationRequest {
  s.CoordinatorUserId = &v
  return s
}

func (s *EndCoordinationRequest) SetInstanceId(v string) *EndCoordinationRequest {
  s.InstanceId = &v
  return s
}

func (s *EndCoordinationRequest) SetOwnerUserId(v string) *EndCoordinationRequest {
  s.OwnerUserId = &v
  return s
}

func (s *EndCoordinationRequest) Validate() error {
  return dara.Validate(s)
}

