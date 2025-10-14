// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHtmlTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHtmlTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHtmlTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHtmlTranslateTaskResponseBody) *SubmitHtmlTranslateTaskResponse
	GetBody() *SubmitHtmlTranslateTaskResponseBody
}

type SubmitHtmlTranslateTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHtmlTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHtmlTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHtmlTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHtmlTranslateTaskResponse) GetBody() *SubmitHtmlTranslateTaskResponseBody {
	return s.Body
}

func (s *SubmitHtmlTranslateTaskResponse) SetHeaders(v map[string]*string) *SubmitHtmlTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHtmlTranslateTaskResponse) SetStatusCode(v int32) *SubmitHtmlTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHtmlTranslateTaskResponse) SetBody(v *SubmitHtmlTranslateTaskResponseBody) *SubmitHtmlTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitHtmlTranslateTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
