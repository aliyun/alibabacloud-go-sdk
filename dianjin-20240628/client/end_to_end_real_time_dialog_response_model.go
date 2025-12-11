// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndToEndRealTimeDialogResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndToEndRealTimeDialogResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndToEndRealTimeDialogResponse
  GetStatusCode() *int32 
  SetBody(v *EndToEndRealTimeDialogResponseBody) *EndToEndRealTimeDialogResponse
  GetBody() *EndToEndRealTimeDialogResponseBody 
}

type EndToEndRealTimeDialogResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndToEndRealTimeDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndToEndRealTimeDialogResponse) String() string {
  return dara.Prettify(s)
}

func (s EndToEndRealTimeDialogResponse) GoString() string {
  return s.String()
}

func (s *EndToEndRealTimeDialogResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndToEndRealTimeDialogResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndToEndRealTimeDialogResponse) GetBody() *EndToEndRealTimeDialogResponseBody  {
  return s.Body
}

func (s *EndToEndRealTimeDialogResponse) SetHeaders(v map[string]*string) *EndToEndRealTimeDialogResponse {
  s.Headers = v
  return s
}

func (s *EndToEndRealTimeDialogResponse) SetStatusCode(v int32) *EndToEndRealTimeDialogResponse {
  s.StatusCode = &v
  return s
}

func (s *EndToEndRealTimeDialogResponse) SetBody(v *EndToEndRealTimeDialogResponseBody) *EndToEndRealTimeDialogResponse {
  s.Body = v
  return s
}

func (s *EndToEndRealTimeDialogResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

