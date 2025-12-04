// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryVccResponse
	GetStatusCode() *int32
	SetBody(v *RetryVccResponseBody) *RetryVccResponse
	GetBody() *RetryVccResponseBody
}

type RetryVccResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryVccResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryVccResponse) GoString() string {
	return s.String()
}

func (s *RetryVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryVccResponse) GetBody() *RetryVccResponseBody {
	return s.Body
}

func (s *RetryVccResponse) SetHeaders(v map[string]*string) *RetryVccResponse {
	s.Headers = v
	return s
}

func (s *RetryVccResponse) SetStatusCode(v int32) *RetryVccResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryVccResponse) SetBody(v *RetryVccResponseBody) *RetryVccResponse {
	s.Body = v
	return s
}

func (s *RetryVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
