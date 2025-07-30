// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInitDomainAutoRedirectRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableInitDomainAutoRedirectRequest
  GetInstanceId() *string 
}

type EnableInitDomainAutoRedirectRequest struct {
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInitDomainAutoRedirectRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInitDomainAutoRedirectRequest) GoString() string {
  return s.String()
}

func (s *EnableInitDomainAutoRedirectRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInitDomainAutoRedirectRequest) SetInstanceId(v string) *EnableInitDomainAutoRedirectRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInitDomainAutoRedirectRequest) Validate() error {
  return dara.Validate(s)
}

