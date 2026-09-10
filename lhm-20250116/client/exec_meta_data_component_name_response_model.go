// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaDataComponentNameResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecMetaDataComponentNameResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecMetaDataComponentNameResponse
  GetStatusCode() *int32 
  SetBody(v *ExecMetaDataComponentNameResponseBody) *ExecMetaDataComponentNameResponse
  GetBody() *ExecMetaDataComponentNameResponseBody 
}

type ExecMetaDataComponentNameResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecMetaDataComponentNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecMetaDataComponentNameResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaDataComponentNameResponse) GoString() string {
  return s.String()
}

func (s *ExecMetaDataComponentNameResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecMetaDataComponentNameResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecMetaDataComponentNameResponse) GetBody() *ExecMetaDataComponentNameResponseBody  {
  return s.Body
}

func (s *ExecMetaDataComponentNameResponse) SetHeaders(v map[string]*string) *ExecMetaDataComponentNameResponse {
  s.Headers = v
  return s
}

func (s *ExecMetaDataComponentNameResponse) SetStatusCode(v int32) *ExecMetaDataComponentNameResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecMetaDataComponentNameResponse) SetBody(v *ExecMetaDataComponentNameResponseBody) *ExecMetaDataComponentNameResponse {
  s.Body = v
  return s
}

func (s *ExecMetaDataComponentNameResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

