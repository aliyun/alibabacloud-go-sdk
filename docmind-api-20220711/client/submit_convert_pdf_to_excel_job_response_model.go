// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToExcelJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertPdfToExcelJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertPdfToExcelJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertPdfToExcelJobResponseBody) *SubmitConvertPdfToExcelJobResponse
	GetBody() *SubmitConvertPdfToExcelJobResponseBody
}

type SubmitConvertPdfToExcelJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertPdfToExcelJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertPdfToExcelJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertPdfToExcelJobResponse) GetBody() *SubmitConvertPdfToExcelJobResponseBody {
	return s.Body
}

func (s *SubmitConvertPdfToExcelJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToExcelJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToExcelJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponse) SetBody(v *SubmitConvertPdfToExcelJobResponseBody) *SubmitConvertPdfToExcelJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponse) Validate() error {
	return dara.Validate(s)
}
