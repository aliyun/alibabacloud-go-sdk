// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocTranslateTaskResponseBody) *SubmitDocTranslateTaskResponse
	GetBody() *SubmitDocTranslateTaskResponseBody
}

type SubmitDocTranslateTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocTranslateTaskResponse) GetBody() *SubmitDocTranslateTaskResponseBody {
	return s.Body
}

func (s *SubmitDocTranslateTaskResponse) SetHeaders(v map[string]*string) *SubmitDocTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocTranslateTaskResponse) SetStatusCode(v int32) *SubmitDocTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocTranslateTaskResponse) SetBody(v *SubmitDocTranslateTaskResponseBody) *SubmitDocTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitDocTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
