// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryImageCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryImageCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryImageCacheResponse
	GetStatusCode() *int32
	SetBody(v *RetryImageCacheResponseBody) *RetryImageCacheResponse
	GetBody() *RetryImageCacheResponseBody
}

type RetryImageCacheResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryImageCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheResponse) GoString() string {
	return s.String()
}

func (s *RetryImageCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryImageCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryImageCacheResponse) GetBody() *RetryImageCacheResponseBody {
	return s.Body
}

func (s *RetryImageCacheResponse) SetHeaders(v map[string]*string) *RetryImageCacheResponse {
	s.Headers = v
	return s
}

func (s *RetryImageCacheResponse) SetStatusCode(v int32) *RetryImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryImageCacheResponse) SetBody(v *RetryImageCacheResponseBody) *RetryImageCacheResponse {
	s.Body = v
	return s
}

func (s *RetryImageCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
