// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCenVbrHealthCheckResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCenVbrHealthCheckResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCenVbrHealthCheckResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCenVbrHealthCheckResponseBody) *EnableCenVbrHealthCheckResponse
  GetBody() *EnableCenVbrHealthCheckResponseBody 
}

type EnableCenVbrHealthCheckResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCenVbrHealthCheckResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCenVbrHealthCheckResponse) GoString() string {
  return s.String()
}

func (s *EnableCenVbrHealthCheckResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCenVbrHealthCheckResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCenVbrHealthCheckResponse) GetBody() *EnableCenVbrHealthCheckResponseBody  {
  return s.Body
}

func (s *EnableCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *EnableCenVbrHealthCheckResponse {
  s.Headers = v
  return s
}

func (s *EnableCenVbrHealthCheckResponse) SetStatusCode(v int32) *EnableCenVbrHealthCheckResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCenVbrHealthCheckResponse) SetBody(v *EnableCenVbrHealthCheckResponseBody) *EnableCenVbrHealthCheckResponse {
  s.Body = v
  return s
}

func (s *EnableCenVbrHealthCheckResponse) Validate() error {
  return dara.Validate(s)
}

