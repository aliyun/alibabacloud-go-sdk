// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAssociatedTransferResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAssociatedTransferResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAssociatedTransferResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAssociatedTransferResponseBody) *EnableAssociatedTransferResponse
  GetBody() *EnableAssociatedTransferResponseBody 
}

type EnableAssociatedTransferResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAssociatedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAssociatedTransferResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAssociatedTransferResponse) GoString() string {
  return s.String()
}

func (s *EnableAssociatedTransferResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAssociatedTransferResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAssociatedTransferResponse) GetBody() *EnableAssociatedTransferResponseBody  {
  return s.Body
}

func (s *EnableAssociatedTransferResponse) SetHeaders(v map[string]*string) *EnableAssociatedTransferResponse {
  s.Headers = v
  return s
}

func (s *EnableAssociatedTransferResponse) SetStatusCode(v int32) *EnableAssociatedTransferResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAssociatedTransferResponse) SetBody(v *EnableAssociatedTransferResponseBody) *EnableAssociatedTransferResponse {
  s.Body = v
  return s
}

func (s *EnableAssociatedTransferResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

