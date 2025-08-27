// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExceedApplySyncResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExceedApplySyncResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExceedApplySyncResponse
  GetStatusCode() *int32 
  SetBody(v *ExceedApplySyncResponseBody) *ExceedApplySyncResponse
  GetBody() *ExceedApplySyncResponseBody 
}

type ExceedApplySyncResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExceedApplySyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExceedApplySyncResponse) String() string {
  return dara.Prettify(s)
}

func (s ExceedApplySyncResponse) GoString() string {
  return s.String()
}

func (s *ExceedApplySyncResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExceedApplySyncResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExceedApplySyncResponse) GetBody() *ExceedApplySyncResponseBody  {
  return s.Body
}

func (s *ExceedApplySyncResponse) SetHeaders(v map[string]*string) *ExceedApplySyncResponse {
  s.Headers = v
  return s
}

func (s *ExceedApplySyncResponse) SetStatusCode(v int32) *ExceedApplySyncResponse {
  s.StatusCode = &v
  return s
}

func (s *ExceedApplySyncResponse) SetBody(v *ExceedApplySyncResponseBody) *ExceedApplySyncResponse {
  s.Body = v
  return s
}

func (s *ExceedApplySyncResponse) Validate() error {
  return dara.Validate(s)
}

