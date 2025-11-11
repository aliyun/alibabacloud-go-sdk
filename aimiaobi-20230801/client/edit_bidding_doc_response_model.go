// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditBiddingDocResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditBiddingDocResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditBiddingDocResponse
  GetStatusCode() *int32 
  SetBody(v *EditBiddingDocResponseBody) *EditBiddingDocResponse
  GetBody() *EditBiddingDocResponseBody 
}

type EditBiddingDocResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditBiddingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditBiddingDocResponse) String() string {
  return dara.Prettify(s)
}

func (s EditBiddingDocResponse) GoString() string {
  return s.String()
}

func (s *EditBiddingDocResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditBiddingDocResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditBiddingDocResponse) GetBody() *EditBiddingDocResponseBody  {
  return s.Body
}

func (s *EditBiddingDocResponse) SetHeaders(v map[string]*string) *EditBiddingDocResponse {
  s.Headers = v
  return s
}

func (s *EditBiddingDocResponse) SetStatusCode(v int32) *EditBiddingDocResponse {
  s.StatusCode = &v
  return s
}

func (s *EditBiddingDocResponse) SetBody(v *EditBiddingDocResponseBody) *EditBiddingDocResponse {
  s.Body = v
  return s
}

func (s *EditBiddingDocResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

