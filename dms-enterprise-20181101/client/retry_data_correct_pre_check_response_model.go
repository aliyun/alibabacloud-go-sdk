// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDataCorrectPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryDataCorrectPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryDataCorrectPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *RetryDataCorrectPreCheckResponseBody) *RetryDataCorrectPreCheckResponse
	GetBody() *RetryDataCorrectPreCheckResponseBody
}

type RetryDataCorrectPreCheckResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDataCorrectPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDataCorrectPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryDataCorrectPreCheckResponse) GoString() string {
	return s.String()
}

func (s *RetryDataCorrectPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryDataCorrectPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryDataCorrectPreCheckResponse) GetBody() *RetryDataCorrectPreCheckResponseBody {
	return s.Body
}

func (s *RetryDataCorrectPreCheckResponse) SetHeaders(v map[string]*string) *RetryDataCorrectPreCheckResponse {
	s.Headers = v
	return s
}

func (s *RetryDataCorrectPreCheckResponse) SetStatusCode(v int32) *RetryDataCorrectPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDataCorrectPreCheckResponse) SetBody(v *RetryDataCorrectPreCheckResponseBody) *RetryDataCorrectPreCheckResponse {
	s.Body = v
	return s
}

func (s *RetryDataCorrectPreCheckResponse) Validate() error {
	return dara.Validate(s)
}
