// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVocabularyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportVocabularyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportVocabularyResponse
  GetStatusCode() *int32 
  SetBody(v *ExportVocabularyResponseBody) *ExportVocabularyResponse
  GetBody() *ExportVocabularyResponseBody 
}

type ExportVocabularyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportVocabularyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportVocabularyResponse) GoString() string {
  return s.String()
}

func (s *ExportVocabularyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportVocabularyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportVocabularyResponse) GetBody() *ExportVocabularyResponseBody  {
  return s.Body
}

func (s *ExportVocabularyResponse) SetHeaders(v map[string]*string) *ExportVocabularyResponse {
  s.Headers = v
  return s
}

func (s *ExportVocabularyResponse) SetStatusCode(v int32) *ExportVocabularyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportVocabularyResponse) SetBody(v *ExportVocabularyResponseBody) *ExportVocabularyResponse {
  s.Body = v
  return s
}

func (s *ExportVocabularyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

