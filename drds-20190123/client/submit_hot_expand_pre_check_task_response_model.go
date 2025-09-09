// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandPreCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHotExpandPreCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHotExpandPreCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHotExpandPreCheckTaskResponseBody) *SubmitHotExpandPreCheckTaskResponse
	GetBody() *SubmitHotExpandPreCheckTaskResponseBody
}

type SubmitHotExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotExpandPreCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHotExpandPreCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHotExpandPreCheckTaskResponse) GetBody() *SubmitHotExpandPreCheckTaskResponseBody {
	return s.Body
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitHotExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitHotExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) SetBody(v *SubmitHotExpandPreCheckTaskResponseBody) *SubmitHotExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponse) Validate() error {
	return dara.Validate(s)
}
