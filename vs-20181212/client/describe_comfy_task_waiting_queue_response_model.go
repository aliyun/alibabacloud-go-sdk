// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTaskWaitingQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyTaskWaitingQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyTaskWaitingQueueResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyTaskWaitingQueueResponseBody) *DescribeComfyTaskWaitingQueueResponse
	GetBody() *DescribeComfyTaskWaitingQueueResponseBody
}

type DescribeComfyTaskWaitingQueueResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyTaskWaitingQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyTaskWaitingQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTaskWaitingQueueResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyTaskWaitingQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyTaskWaitingQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyTaskWaitingQueueResponse) GetBody() *DescribeComfyTaskWaitingQueueResponseBody {
	return s.Body
}

func (s *DescribeComfyTaskWaitingQueueResponse) SetHeaders(v map[string]*string) *DescribeComfyTaskWaitingQueueResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponse) SetStatusCode(v int32) *DescribeComfyTaskWaitingQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponse) SetBody(v *DescribeComfyTaskWaitingQueueResponseBody) *DescribeComfyTaskWaitingQueueResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyTaskWaitingQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
