// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndpointStatus interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndpointStatus
  GetCode() *string 
  SetDetail(v *EndpointStatusDetail) *EndpointStatus
  GetDetail() *EndpointStatusDetail 
  SetMessage(v string) *EndpointStatus
  GetMessage() *string 
  SetPhase(v string) *EndpointStatus
  GetPhase() *string 
}

type EndpointStatus struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Detail *EndpointStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty"`
  // example:
  // 
  // Init Succeed
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // Ready
  Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
}

func (s EndpointStatus) String() string {
  return dara.Prettify(s)
}

func (s EndpointStatus) GoString() string {
  return s.String()
}

func (s *EndpointStatus) GetCode() *string  {
  return s.Code
}

func (s *EndpointStatus) GetDetail() *EndpointStatusDetail  {
  return s.Detail
}

func (s *EndpointStatus) GetMessage() *string  {
  return s.Message
}

func (s *EndpointStatus) GetPhase() *string  {
  return s.Phase
}

func (s *EndpointStatus) SetCode(v string) *EndpointStatus {
  s.Code = &v
  return s
}

func (s *EndpointStatus) SetDetail(v *EndpointStatusDetail) *EndpointStatus {
  s.Detail = v
  return s
}

func (s *EndpointStatus) SetMessage(v string) *EndpointStatus {
  s.Message = &v
  return s
}

func (s *EndpointStatus) SetPhase(v string) *EndpointStatus {
  s.Phase = &v
  return s
}

func (s *EndpointStatus) Validate() error {
  if s.Detail != nil {
    if err := s.Detail.Validate(); err != nil {
      return err
    }
  }
  return nil
}

