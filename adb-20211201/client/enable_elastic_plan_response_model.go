// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableElasticPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableElasticPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableElasticPlanResponse
  GetStatusCode() *int32 
  SetBody(v *EnableElasticPlanResponseBody) *EnableElasticPlanResponse
  GetBody() *EnableElasticPlanResponseBody 
}

type EnableElasticPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableElasticPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableElasticPlanResponse) GoString() string {
  return s.String()
}

func (s *EnableElasticPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableElasticPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableElasticPlanResponse) GetBody() *EnableElasticPlanResponseBody  {
  return s.Body
}

func (s *EnableElasticPlanResponse) SetHeaders(v map[string]*string) *EnableElasticPlanResponse {
  s.Headers = v
  return s
}

func (s *EnableElasticPlanResponse) SetStatusCode(v int32) *EnableElasticPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableElasticPlanResponse) SetBody(v *EnableElasticPlanResponseBody) *EnableElasticPlanResponse {
  s.Body = v
  return s
}

func (s *EnableElasticPlanResponse) Validate() error {
  return dara.Validate(s)
}

