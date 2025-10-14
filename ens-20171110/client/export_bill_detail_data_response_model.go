// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportBillDetailDataResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportBillDetailDataResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportBillDetailDataResponse
  GetStatusCode() *int32 
  SetBody(v *ExportBillDetailDataResponseBody) *ExportBillDetailDataResponse
  GetBody() *ExportBillDetailDataResponseBody 
}

type ExportBillDetailDataResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportBillDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportBillDetailDataResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportBillDetailDataResponse) GoString() string {
  return s.String()
}

func (s *ExportBillDetailDataResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportBillDetailDataResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportBillDetailDataResponse) GetBody() *ExportBillDetailDataResponseBody  {
  return s.Body
}

func (s *ExportBillDetailDataResponse) SetHeaders(v map[string]*string) *ExportBillDetailDataResponse {
  s.Headers = v
  return s
}

func (s *ExportBillDetailDataResponse) SetStatusCode(v int32) *ExportBillDetailDataResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportBillDetailDataResponse) SetBody(v *ExportBillDetailDataResponseBody) *ExportBillDetailDataResponse {
  s.Body = v
  return s
}

func (s *ExportBillDetailDataResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

