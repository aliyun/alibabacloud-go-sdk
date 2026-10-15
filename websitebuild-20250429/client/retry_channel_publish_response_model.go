// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChannelPublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryChannelPublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryChannelPublishResponse
	GetStatusCode() *int32
	SetBody(v *RetryChannelPublishResponseBody) *RetryChannelPublishResponse
	GetBody() *RetryChannelPublishResponseBody
}

type RetryChannelPublishResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryChannelPublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryChannelPublishResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryChannelPublishResponse) GoString() string {
	return s.String()
}

func (s *RetryChannelPublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryChannelPublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryChannelPublishResponse) GetBody() *RetryChannelPublishResponseBody {
	return s.Body
}

func (s *RetryChannelPublishResponse) SetHeaders(v map[string]*string) *RetryChannelPublishResponse {
	s.Headers = v
	return s
}

func (s *RetryChannelPublishResponse) SetStatusCode(v int32) *RetryChannelPublishResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryChannelPublishResponse) SetBody(v *RetryChannelPublishResponseBody) *RetryChannelPublishResponse {
	s.Body = v
	return s
}

func (s *RetryChannelPublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
