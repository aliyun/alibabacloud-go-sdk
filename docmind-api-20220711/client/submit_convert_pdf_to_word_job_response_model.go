// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToWordJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertPdfToWordJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertPdfToWordJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertPdfToWordJobResponseBody) *SubmitConvertPdfToWordJobResponse
	GetBody() *SubmitConvertPdfToWordJobResponseBody
}

type SubmitConvertPdfToWordJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertPdfToWordJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertPdfToWordJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertPdfToWordJobResponse) GetBody() *SubmitConvertPdfToWordJobResponseBody {
	return s.Body
}

func (s *SubmitConvertPdfToWordJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToWordJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToWordJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToWordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponse) SetBody(v *SubmitConvertPdfToWordJobResponseBody) *SubmitConvertPdfToWordJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertPdfToWordJobResponse) Validate() error {
	return dara.Validate(s)
}
