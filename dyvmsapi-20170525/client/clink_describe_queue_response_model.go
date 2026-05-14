// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDescribeQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDescribeQueueResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDescribeQueueResponseBody) *ClinkDescribeQueueResponse
	GetBody() *ClinkDescribeQueueResponseBody
}

type ClinkDescribeQueueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDescribeQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDescribeQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueResponse) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDescribeQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDescribeQueueResponse) GetBody() *ClinkDescribeQueueResponseBody {
	return s.Body
}

func (s *ClinkDescribeQueueResponse) SetHeaders(v map[string]*string) *ClinkDescribeQueueResponse {
	s.Headers = v
	return s
}

func (s *ClinkDescribeQueueResponse) SetStatusCode(v int32) *ClinkDescribeQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDescribeQueueResponse) SetBody(v *ClinkDescribeQueueResponseBody) *ClinkDescribeQueueResponse {
	s.Body = v
	return s
}

func (s *ClinkDescribeQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
