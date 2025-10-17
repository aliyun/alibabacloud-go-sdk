// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryDocResponse
	GetStatusCode() *int32
	SetBody(v *RetryDocResponseBody) *RetryDocResponse
	GetBody() *RetryDocResponseBody
}

type RetryDocResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDocResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryDocResponse) GoString() string {
	return s.String()
}

func (s *RetryDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryDocResponse) GetBody() *RetryDocResponseBody {
	return s.Body
}

func (s *RetryDocResponse) SetHeaders(v map[string]*string) *RetryDocResponse {
	s.Headers = v
	return s
}

func (s *RetryDocResponse) SetStatusCode(v int32) *RetryDocResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDocResponse) SetBody(v *RetryDocResponseBody) *RetryDocResponse {
	s.Body = v
	return s
}

func (s *RetryDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
