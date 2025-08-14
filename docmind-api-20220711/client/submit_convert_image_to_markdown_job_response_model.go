// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToMarkdownJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertImageToMarkdownJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertImageToMarkdownJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertImageToMarkdownJobResponseBody) *SubmitConvertImageToMarkdownJobResponse
	GetBody() *SubmitConvertImageToMarkdownJobResponseBody
}

type SubmitConvertImageToMarkdownJobResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToMarkdownJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertImageToMarkdownJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertImageToMarkdownJobResponse) GetBody() *SubmitConvertImageToMarkdownJobResponseBody {
	return s.Body
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToMarkdownJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetStatusCode(v int32) *SubmitConvertImageToMarkdownJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponse) SetBody(v *SubmitConvertImageToMarkdownJobResponseBody) *SubmitConvertImageToMarkdownJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponse) Validate() error {
	return dara.Validate(s)
}
