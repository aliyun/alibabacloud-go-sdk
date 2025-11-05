// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomerQuotaRecordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportCustomerQuotaRecordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportCustomerQuotaRecordResponse
  GetStatusCode() *int32 
  SetBody(v *ExportCustomerQuotaRecordResponseBody) *ExportCustomerQuotaRecordResponse
  GetBody() *ExportCustomerQuotaRecordResponseBody 
}

type ExportCustomerQuotaRecordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportCustomerQuotaRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCustomerQuotaRecordResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomerQuotaRecordResponse) GoString() string {
  return s.String()
}

func (s *ExportCustomerQuotaRecordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportCustomerQuotaRecordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportCustomerQuotaRecordResponse) GetBody() *ExportCustomerQuotaRecordResponseBody  {
  return s.Body
}

func (s *ExportCustomerQuotaRecordResponse) SetHeaders(v map[string]*string) *ExportCustomerQuotaRecordResponse {
  s.Headers = v
  return s
}

func (s *ExportCustomerQuotaRecordResponse) SetStatusCode(v int32) *ExportCustomerQuotaRecordResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportCustomerQuotaRecordResponse) SetBody(v *ExportCustomerQuotaRecordResponseBody) *ExportCustomerQuotaRecordResponse {
  s.Body = v
  return s
}

func (s *ExportCustomerQuotaRecordResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

