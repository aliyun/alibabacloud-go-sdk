// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataIngestionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDataIngestionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDataIngestionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDataIngestionResponseBody) *EnableDataIngestionResponse
  GetBody() *EnableDataIngestionResponseBody 
}

type EnableDataIngestionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDataIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDataIngestionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDataIngestionResponse) GoString() string {
  return s.String()
}

func (s *EnableDataIngestionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDataIngestionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDataIngestionResponse) GetBody() *EnableDataIngestionResponseBody  {
  return s.Body
}

func (s *EnableDataIngestionResponse) SetHeaders(v map[string]*string) *EnableDataIngestionResponse {
  s.Headers = v
  return s
}

func (s *EnableDataIngestionResponse) SetStatusCode(v int32) *EnableDataIngestionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDataIngestionResponse) SetBody(v *EnableDataIngestionResponseBody) *EnableDataIngestionResponse {
  s.Body = v
  return s
}

func (s *EnableDataIngestionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

