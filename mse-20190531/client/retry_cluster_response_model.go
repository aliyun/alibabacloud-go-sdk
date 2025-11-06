// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryClusterResponse
	GetStatusCode() *int32
	SetBody(v *RetryClusterResponseBody) *RetryClusterResponse
	GetBody() *RetryClusterResponseBody
}

type RetryClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryClusterResponse) GoString() string {
	return s.String()
}

func (s *RetryClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryClusterResponse) GetBody() *RetryClusterResponseBody {
	return s.Body
}

func (s *RetryClusterResponse) SetHeaders(v map[string]*string) *RetryClusterResponse {
	s.Headers = v
	return s
}

func (s *RetryClusterResponse) SetStatusCode(v int32) *RetryClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryClusterResponse) SetBody(v *RetryClusterResponseBody) *RetryClusterResponse {
	s.Body = v
	return s
}

func (s *RetryClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
