// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalSlbRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableInternalSlbRequest
  GetInstanceId() *string 
}

type EnableInternalSlbRequest struct {
  // example:
  // 
  // c-b25e21e24388****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInternalSlbRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalSlbRequest) GoString() string {
  return s.String()
}

func (s *EnableInternalSlbRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInternalSlbRequest) SetInstanceId(v string) *EnableInternalSlbRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInternalSlbRequest) Validate() error {
  return dara.Validate(s)
}

