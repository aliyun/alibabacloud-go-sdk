// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkCreateQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkCreateQueueResponse
	GetStatusCode() *int32
	SetBody(v *ClinkCreateQueueResponseBody) *ClinkCreateQueueResponse
	GetBody() *ClinkCreateQueueResponseBody
}

type ClinkCreateQueueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkCreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkCreateQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueResponse) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkCreateQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkCreateQueueResponse) GetBody() *ClinkCreateQueueResponseBody {
	return s.Body
}

func (s *ClinkCreateQueueResponse) SetHeaders(v map[string]*string) *ClinkCreateQueueResponse {
	s.Headers = v
	return s
}

func (s *ClinkCreateQueueResponse) SetStatusCode(v int32) *ClinkCreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkCreateQueueResponse) SetBody(v *ClinkCreateQueueResponseBody) *ClinkCreateQueueResponse {
	s.Body = v
	return s
}

func (s *ClinkCreateQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
