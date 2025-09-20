// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractDocumentTextResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExtractDocumentTextResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExtractDocumentTextResponse
  GetStatusCode() *int32 
  SetBody(v *ExtractDocumentTextResponseBody) *ExtractDocumentTextResponse
  GetBody() *ExtractDocumentTextResponseBody 
}

type ExtractDocumentTextResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExtractDocumentTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtractDocumentTextResponse) String() string {
  return dara.Prettify(s)
}

func (s ExtractDocumentTextResponse) GoString() string {
  return s.String()
}

func (s *ExtractDocumentTextResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExtractDocumentTextResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExtractDocumentTextResponse) GetBody() *ExtractDocumentTextResponseBody  {
  return s.Body
}

func (s *ExtractDocumentTextResponse) SetHeaders(v map[string]*string) *ExtractDocumentTextResponse {
  s.Headers = v
  return s
}

func (s *ExtractDocumentTextResponse) SetStatusCode(v int32) *ExtractDocumentTextResponse {
  s.StatusCode = &v
  return s
}

func (s *ExtractDocumentTextResponse) SetBody(v *ExtractDocumentTextResponseBody) *ExtractDocumentTextResponse {
  s.Body = v
  return s
}

func (s *ExtractDocumentTextResponse) Validate() error {
  return dara.Validate(s)
}

