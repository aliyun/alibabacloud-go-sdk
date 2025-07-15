// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterVideoResourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EffectCasterVideoResourceResponseBody
  GetRequestId() *string 
}

type EffectCasterVideoResourceResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // CF60DB6A-7FD6-426E-9288-122CC1A52FA7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EffectCasterVideoResourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterVideoResourceResponseBody) GoString() string {
  return s.String()
}

func (s *EffectCasterVideoResourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EffectCasterVideoResourceResponseBody) SetRequestId(v string) *EffectCasterVideoResourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EffectCasterVideoResourceResponseBody) Validate() error {
  return dara.Validate(s)
}

