// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterUrgentResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EffectCasterUrgentResponseBody
  GetRequestId() *string 
}

type EffectCasterUrgentResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // CF60DB6A-7FD6-426E-9288-122CC1A52FA7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EffectCasterUrgentResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterUrgentResponseBody) GoString() string {
  return s.String()
}

func (s *EffectCasterUrgentResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EffectCasterUrgentResponseBody) SetRequestId(v string) *EffectCasterUrgentResponseBody {
  s.RequestId = &v
  return s
}

func (s *EffectCasterUrgentResponseBody) Validate() error {
  return dara.Validate(s)
}

