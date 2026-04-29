// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteQueueResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteQueueResponseBody) *CloudDeleteQueueResponse
	GetBody() *CloudDeleteQueueResponseBody
}

type CloudDeleteQueueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteQueueResponse) GetBody() *CloudDeleteQueueResponseBody {
	return s.Body
}

func (s *CloudDeleteQueueResponse) SetHeaders(v map[string]*string) *CloudDeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteQueueResponse) SetStatusCode(v int32) *CloudDeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteQueueResponse) SetBody(v *CloudDeleteQueueResponseBody) *CloudDeleteQueueResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
