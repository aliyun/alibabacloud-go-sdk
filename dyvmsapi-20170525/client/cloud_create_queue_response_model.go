// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateQueueResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateQueueResponseBody) *CloudCreateQueueResponse
	GetBody() *CloudCreateQueueResponseBody
}

type CloudCreateQueueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateQueueResponse) GetBody() *CloudCreateQueueResponseBody {
	return s.Body
}

func (s *CloudCreateQueueResponse) SetHeaders(v map[string]*string) *CloudCreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateQueueResponse) SetStatusCode(v int32) *CloudCreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateQueueResponse) SetBody(v *CloudCreateQueueResponseBody) *CloudCreateQueueResponse {
	s.Body = v
	return s
}

func (s *CloudCreateQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
