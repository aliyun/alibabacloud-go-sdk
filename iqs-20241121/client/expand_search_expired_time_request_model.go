// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandSearchExpiredTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v *ExpendExpiredTimeRequest) *ExpandSearchExpiredTimeRequest
  GetBody() *ExpendExpiredTimeRequest 
}

type ExpandSearchExpiredTimeRequest struct {
  Body *ExpendExpiredTimeRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandSearchExpiredTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpandSearchExpiredTimeRequest) GoString() string {
  return s.String()
}

func (s *ExpandSearchExpiredTimeRequest) GetBody() *ExpendExpiredTimeRequest  {
  return s.Body
}

func (s *ExpandSearchExpiredTimeRequest) SetBody(v *ExpendExpiredTimeRequest) *ExpandSearchExpiredTimeRequest {
  s.Body = v
  return s
}

func (s *ExpandSearchExpiredTimeRequest) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

