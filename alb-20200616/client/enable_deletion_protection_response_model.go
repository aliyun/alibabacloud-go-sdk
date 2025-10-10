// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeletionProtectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDeletionProtectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDeletionProtectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDeletionProtectionResponseBody) *EnableDeletionProtectionResponse
  GetBody() *EnableDeletionProtectionResponseBody 
}

type EnableDeletionProtectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDeletionProtectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDeletionProtectionResponse) GoString() string {
  return s.String()
}

func (s *EnableDeletionProtectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDeletionProtectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDeletionProtectionResponse) GetBody() *EnableDeletionProtectionResponseBody  {
  return s.Body
}

func (s *EnableDeletionProtectionResponse) SetHeaders(v map[string]*string) *EnableDeletionProtectionResponse {
  s.Headers = v
  return s
}

func (s *EnableDeletionProtectionResponse) SetStatusCode(v int32) *EnableDeletionProtectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDeletionProtectionResponse) SetBody(v *EnableDeletionProtectionResponseBody) *EnableDeletionProtectionResponse {
  s.Body = v
  return s
}

func (s *EnableDeletionProtectionResponse) Validate() error {
  return dara.Validate(s)
}

