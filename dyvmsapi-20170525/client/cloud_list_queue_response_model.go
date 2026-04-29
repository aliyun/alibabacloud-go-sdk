// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListQueueResponse
	GetStatusCode() *int32
	SetBody(v *CloudListQueueResponseBody) *CloudListQueueResponse
	GetBody() *CloudListQueueResponseBody
}

type CloudListQueueResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueResponse) GoString() string {
	return s.String()
}

func (s *CloudListQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListQueueResponse) GetBody() *CloudListQueueResponseBody {
	return s.Body
}

func (s *CloudListQueueResponse) SetHeaders(v map[string]*string) *CloudListQueueResponse {
	s.Headers = v
	return s
}

func (s *CloudListQueueResponse) SetStatusCode(v int32) *CloudListQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListQueueResponse) SetBody(v *CloudListQueueResponseBody) *CloudListQueueResponse {
	s.Body = v
	return s
}

func (s *CloudListQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
