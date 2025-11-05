// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportReversedDeductionHistoryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportReversedDeductionHistoryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportReversedDeductionHistoryResponse
  GetStatusCode() *int32 
  SetBody(v *ExportReversedDeductionHistoryResponseBody) *ExportReversedDeductionHistoryResponse
  GetBody() *ExportReversedDeductionHistoryResponseBody 
}

type ExportReversedDeductionHistoryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportReversedDeductionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportReversedDeductionHistoryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportReversedDeductionHistoryResponse) GoString() string {
  return s.String()
}

func (s *ExportReversedDeductionHistoryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportReversedDeductionHistoryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportReversedDeductionHistoryResponse) GetBody() *ExportReversedDeductionHistoryResponseBody  {
  return s.Body
}

func (s *ExportReversedDeductionHistoryResponse) SetHeaders(v map[string]*string) *ExportReversedDeductionHistoryResponse {
  s.Headers = v
  return s
}

func (s *ExportReversedDeductionHistoryResponse) SetStatusCode(v int32) *ExportReversedDeductionHistoryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportReversedDeductionHistoryResponse) SetBody(v *ExportReversedDeductionHistoryResponseBody) *ExportReversedDeductionHistoryResponse {
  s.Body = v
  return s
}

func (s *ExportReversedDeductionHistoryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

