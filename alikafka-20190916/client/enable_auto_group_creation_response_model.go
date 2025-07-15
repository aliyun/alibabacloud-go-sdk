// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoGroupCreationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAutoGroupCreationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAutoGroupCreationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAutoGroupCreationResponseBody) *EnableAutoGroupCreationResponse
  GetBody() *EnableAutoGroupCreationResponseBody 
}

type EnableAutoGroupCreationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAutoGroupCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoGroupCreationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoGroupCreationResponse) GoString() string {
  return s.String()
}

func (s *EnableAutoGroupCreationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAutoGroupCreationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAutoGroupCreationResponse) GetBody() *EnableAutoGroupCreationResponseBody  {
  return s.Body
}

func (s *EnableAutoGroupCreationResponse) SetHeaders(v map[string]*string) *EnableAutoGroupCreationResponse {
  s.Headers = v
  return s
}

func (s *EnableAutoGroupCreationResponse) SetStatusCode(v int32) *EnableAutoGroupCreationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAutoGroupCreationResponse) SetBody(v *EnableAutoGroupCreationResponseBody) *EnableAutoGroupCreationResponse {
  s.Body = v
  return s
}

func (s *EnableAutoGroupCreationResponse) Validate() error {
  return dara.Validate(s)
}

