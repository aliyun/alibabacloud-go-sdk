// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDirectoryFailedSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryDirectoryFailedSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryDirectoryFailedSourcesResponse
	GetStatusCode() *int32
	SetBody(v *RetryDirectoryFailedSourcesResponseBody) *RetryDirectoryFailedSourcesResponse
	GetBody() *RetryDirectoryFailedSourcesResponseBody
}

type RetryDirectoryFailedSourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryDirectoryFailedSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryDirectoryFailedSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryDirectoryFailedSourcesResponse) GoString() string {
	return s.String()
}

func (s *RetryDirectoryFailedSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryDirectoryFailedSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryDirectoryFailedSourcesResponse) GetBody() *RetryDirectoryFailedSourcesResponseBody {
	return s.Body
}

func (s *RetryDirectoryFailedSourcesResponse) SetHeaders(v map[string]*string) *RetryDirectoryFailedSourcesResponse {
	s.Headers = v
	return s
}

func (s *RetryDirectoryFailedSourcesResponse) SetStatusCode(v int32) *RetryDirectoryFailedSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryDirectoryFailedSourcesResponse) SetBody(v *RetryDirectoryFailedSourcesResponseBody) *RetryDirectoryFailedSourcesResponse {
	s.Body = v
	return s
}

func (s *RetryDirectoryFailedSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
