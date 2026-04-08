// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstablishRelationTableToBusinessResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EstablishRelationTableToBusinessResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EstablishRelationTableToBusinessResponse
  GetStatusCode() *int32 
  SetBody(v *EstablishRelationTableToBusinessResponseBody) *EstablishRelationTableToBusinessResponse
  GetBody() *EstablishRelationTableToBusinessResponseBody 
}

type EstablishRelationTableToBusinessResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EstablishRelationTableToBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EstablishRelationTableToBusinessResponse) String() string {
  return dara.Prettify(s)
}

func (s EstablishRelationTableToBusinessResponse) GoString() string {
  return s.String()
}

func (s *EstablishRelationTableToBusinessResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EstablishRelationTableToBusinessResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EstablishRelationTableToBusinessResponse) GetBody() *EstablishRelationTableToBusinessResponseBody  {
  return s.Body
}

func (s *EstablishRelationTableToBusinessResponse) SetHeaders(v map[string]*string) *EstablishRelationTableToBusinessResponse {
  s.Headers = v
  return s
}

func (s *EstablishRelationTableToBusinessResponse) SetStatusCode(v int32) *EstablishRelationTableToBusinessResponse {
  s.StatusCode = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponse) SetBody(v *EstablishRelationTableToBusinessResponseBody) *EstablishRelationTableToBusinessResponse {
  s.Body = v
  return s
}

func (s *EstablishRelationTableToBusinessResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

