// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmoothExpandPreCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmoothExpandPreCheckTaskResponseBody) *SubmitSmoothExpandPreCheckTaskResponse
	GetBody() *SubmitSmoothExpandPreCheckTaskResponseBody
}

type SubmitSmoothExpandPreCheckTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmoothExpandPreCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) GetBody() *SubmitSmoothExpandPreCheckTaskResponseBody {
	return s.Body
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetHeaders(v map[string]*string) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetStatusCode(v int32) *SubmitSmoothExpandPreCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) SetBody(v *SubmitSmoothExpandPreCheckTaskResponseBody) *SubmitSmoothExpandPreCheckTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponse) Validate() error {
	return dara.Validate(s)
}
