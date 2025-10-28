// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeOrderTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryChangeOrderTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryChangeOrderTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryChangeOrderTaskResponseBody) *RetryChangeOrderTaskResponse
	GetBody() *RetryChangeOrderTaskResponseBody
}

type RetryChangeOrderTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryChangeOrderTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryChangeOrderTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeOrderTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryChangeOrderTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryChangeOrderTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryChangeOrderTaskResponse) GetBody() *RetryChangeOrderTaskResponseBody {
	return s.Body
}

func (s *RetryChangeOrderTaskResponse) SetHeaders(v map[string]*string) *RetryChangeOrderTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryChangeOrderTaskResponse) SetStatusCode(v int32) *RetryChangeOrderTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryChangeOrderTaskResponse) SetBody(v *RetryChangeOrderTaskResponseBody) *RetryChangeOrderTaskResponse {
	s.Body = v
	return s
}

func (s *RetryChangeOrderTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
