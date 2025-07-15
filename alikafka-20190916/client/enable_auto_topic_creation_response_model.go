// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoTopicCreationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAutoTopicCreationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAutoTopicCreationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAutoTopicCreationResponseBody) *EnableAutoTopicCreationResponse
  GetBody() *EnableAutoTopicCreationResponseBody 
}

type EnableAutoTopicCreationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAutoTopicCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoTopicCreationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoTopicCreationResponse) GoString() string {
  return s.String()
}

func (s *EnableAutoTopicCreationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAutoTopicCreationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAutoTopicCreationResponse) GetBody() *EnableAutoTopicCreationResponseBody  {
  return s.Body
}

func (s *EnableAutoTopicCreationResponse) SetHeaders(v map[string]*string) *EnableAutoTopicCreationResponse {
  s.Headers = v
  return s
}

func (s *EnableAutoTopicCreationResponse) SetStatusCode(v int32) *EnableAutoTopicCreationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAutoTopicCreationResponse) SetBody(v *EnableAutoTopicCreationResponseBody) *EnableAutoTopicCreationResponse {
  s.Body = v
  return s
}

func (s *EnableAutoTopicCreationResponse) Validate() error {
  return dara.Validate(s)
}

