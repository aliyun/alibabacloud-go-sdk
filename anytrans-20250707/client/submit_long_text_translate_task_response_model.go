// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLongTextTranslateTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLongTextTranslateTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLongTextTranslateTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLongTextTranslateTaskResponseBody) *SubmitLongTextTranslateTaskResponse
	GetBody() *SubmitLongTextTranslateTaskResponseBody
}

type SubmitLongTextTranslateTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLongTextTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLongTextTranslateTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLongTextTranslateTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLongTextTranslateTaskResponse) GetBody() *SubmitLongTextTranslateTaskResponseBody {
	return s.Body
}

func (s *SubmitLongTextTranslateTaskResponse) SetHeaders(v map[string]*string) *SubmitLongTextTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitLongTextTranslateTaskResponse) SetStatusCode(v int32) *SubmitLongTextTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponse) SetBody(v *SubmitLongTextTranslateTaskResponseBody) *SubmitLongTextTranslateTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitLongTextTranslateTaskResponse) Validate() error {
	return dara.Validate(s)
}
