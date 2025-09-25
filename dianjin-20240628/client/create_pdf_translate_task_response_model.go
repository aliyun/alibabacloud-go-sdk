// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdfTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdfTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdfTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatePdfTranslateTaskResponseBody) *CreatePdfTranslateTaskResponse
	GetBody() *CreatePdfTranslateTaskResponseBody
}

type CreatePdfTranslateTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePdfTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdfTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdfTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdfTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdfTranslateTaskResponse) GetBody() *CreatePdfTranslateTaskResponseBody {
	return s.Body
}

func (s *CreatePdfTranslateTaskResponse) SetHeaders(v map[string]*string) *CreatePdfTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePdfTranslateTaskResponse) SetStatusCode(v int32) *CreatePdfTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdfTranslateTaskResponse) SetBody(v *CreatePdfTranslateTaskResponseBody) *CreatePdfTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *CreatePdfTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
