// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoLiveStreamRuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAutoLiveStreamRuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAutoLiveStreamRuleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAutoLiveStreamRuleResponseBody) *EnableAutoLiveStreamRuleResponse
  GetBody() *EnableAutoLiveStreamRuleResponseBody 
}

type EnableAutoLiveStreamRuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoLiveStreamRuleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoLiveStreamRuleResponse) GoString() string {
  return s.String()
}

func (s *EnableAutoLiveStreamRuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAutoLiveStreamRuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAutoLiveStreamRuleResponse) GetBody() *EnableAutoLiveStreamRuleResponseBody  {
  return s.Body
}

func (s *EnableAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *EnableAutoLiveStreamRuleResponse {
  s.Headers = v
  return s
}

func (s *EnableAutoLiveStreamRuleResponse) SetStatusCode(v int32) *EnableAutoLiveStreamRuleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAutoLiveStreamRuleResponse) SetBody(v *EnableAutoLiveStreamRuleResponseBody) *EnableAutoLiveStreamRuleResponse {
  s.Body = v
  return s
}

func (s *EnableAutoLiveStreamRuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

