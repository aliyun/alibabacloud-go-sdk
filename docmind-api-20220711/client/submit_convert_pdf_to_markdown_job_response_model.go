// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToMarkdownJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertPdfToMarkdownJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertPdfToMarkdownJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertPdfToMarkdownJobResponseBody) *SubmitConvertPdfToMarkdownJobResponse
	GetBody() *SubmitConvertPdfToMarkdownJobResponseBody
}

type SubmitConvertPdfToMarkdownJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertPdfToMarkdownJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertPdfToMarkdownJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertPdfToMarkdownJobResponse) GetBody() *SubmitConvertPdfToMarkdownJobResponseBody {
	return s.Body
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetHeaders(v map[string]*string) *SubmitConvertPdfToMarkdownJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetStatusCode(v int32) *SubmitConvertPdfToMarkdownJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponse) SetBody(v *SubmitConvertPdfToMarkdownJobResponseBody) *SubmitConvertPdfToMarkdownJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponse) Validate() error {
	return dara.Validate(s)
}
