// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEnterpriseAcceleratePolicyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEapId(v string) *EnableEnterpriseAcceleratePolicyRequest
  GetEapId() *string 
}

type EnableEnterpriseAcceleratePolicyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // eap-530da9f7110441fb
  EapId *string `json:"EapId,omitempty" xml:"EapId,omitempty"`
}

func (s EnableEnterpriseAcceleratePolicyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableEnterpriseAcceleratePolicyRequest) GoString() string {
  return s.String()
}

func (s *EnableEnterpriseAcceleratePolicyRequest) GetEapId() *string  {
  return s.EapId
}

func (s *EnableEnterpriseAcceleratePolicyRequest) SetEapId(v string) *EnableEnterpriseAcceleratePolicyRequest {
  s.EapId = &v
  return s
}

func (s *EnableEnterpriseAcceleratePolicyRequest) Validate() error {
  return dara.Validate(s)
}

