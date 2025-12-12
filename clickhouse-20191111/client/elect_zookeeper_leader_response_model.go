// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectZookeeperLeaderResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ElectZookeeperLeaderResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ElectZookeeperLeaderResponse
  GetStatusCode() *int32 
  SetBody(v *ElectZookeeperLeaderResponseBody) *ElectZookeeperLeaderResponse
  GetBody() *ElectZookeeperLeaderResponseBody 
}

type ElectZookeeperLeaderResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ElectZookeeperLeaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ElectZookeeperLeaderResponse) String() string {
  return dara.Prettify(s)
}

func (s ElectZookeeperLeaderResponse) GoString() string {
  return s.String()
}

func (s *ElectZookeeperLeaderResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ElectZookeeperLeaderResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ElectZookeeperLeaderResponse) GetBody() *ElectZookeeperLeaderResponseBody  {
  return s.Body
}

func (s *ElectZookeeperLeaderResponse) SetHeaders(v map[string]*string) *ElectZookeeperLeaderResponse {
  s.Headers = v
  return s
}

func (s *ElectZookeeperLeaderResponse) SetStatusCode(v int32) *ElectZookeeperLeaderResponse {
  s.StatusCode = &v
  return s
}

func (s *ElectZookeeperLeaderResponse) SetBody(v *ElectZookeeperLeaderResponseBody) *ElectZookeeperLeaderResponse {
  s.Body = v
  return s
}

func (s *ElectZookeeperLeaderResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

