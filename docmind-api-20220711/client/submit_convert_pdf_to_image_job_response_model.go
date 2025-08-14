// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToImageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertPdfToImageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertPdfToImageJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertPdfToImageJobResponseBody) *SubmitConvertPdfToImageJobResponse
	GetBody() *SubmitConvertPdfToImageJobResponseBody
}

type SubmitConvertPdfToImageJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertPdfToImageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertPdfToImageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertPdfToImageJobResponse) GetBody() *SubmitConvertPdfToImageJobResponseBody {
	return s.Body
}

func (s *SubmitConvertPdfToImageJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToImageJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponse) SetBody(v *SubmitConvertPdfToImageJobResponseBody) *SubmitConvertPdfToImageJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertPdfToImageJobResponse) Validate() error {
	return dara.Validate(s)
}
