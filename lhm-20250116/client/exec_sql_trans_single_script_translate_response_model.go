// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecSqlTransSingleScriptTranslateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecSqlTransSingleScriptTranslateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecSqlTransSingleScriptTranslateResponse
  GetStatusCode() *int32 
  SetBody(v *ExecSqlTransSingleScriptTranslateResponseBody) *ExecSqlTransSingleScriptTranslateResponse
  GetBody() *ExecSqlTransSingleScriptTranslateResponseBody 
}

type ExecSqlTransSingleScriptTranslateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecSqlTransSingleScriptTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecSqlTransSingleScriptTranslateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecSqlTransSingleScriptTranslateResponse) GoString() string {
  return s.String()
}

func (s *ExecSqlTransSingleScriptTranslateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecSqlTransSingleScriptTranslateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecSqlTransSingleScriptTranslateResponse) GetBody() *ExecSqlTransSingleScriptTranslateResponseBody  {
  return s.Body
}

func (s *ExecSqlTransSingleScriptTranslateResponse) SetHeaders(v map[string]*string) *ExecSqlTransSingleScriptTranslateResponse {
  s.Headers = v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponse) SetStatusCode(v int32) *ExecSqlTransSingleScriptTranslateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponse) SetBody(v *ExecSqlTransSingleScriptTranslateResponseBody) *ExecSqlTransSingleScriptTranslateResponse {
  s.Body = v
  return s
}

func (s *ExecSqlTransSingleScriptTranslateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

