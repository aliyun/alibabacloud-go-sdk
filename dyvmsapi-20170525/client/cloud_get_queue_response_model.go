// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetQueueResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetQueueResponseBody) *CloudGetQueueResponse
	GetBody() *CloudGetQueueResponseBody
}

type CloudGetQueueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueResponse) GoString() string {
	return s.String()
}

func (s *CloudGetQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetQueueResponse) GetBody() *CloudGetQueueResponseBody {
	return s.Body
}

func (s *CloudGetQueueResponse) SetHeaders(v map[string]*string) *CloudGetQueueResponse {
	s.Headers = v
	return s
}

func (s *CloudGetQueueResponse) SetStatusCode(v int32) *CloudGetQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetQueueResponse) SetBody(v *CloudGetQueueResponseBody) *CloudGetQueueResponse {
	s.Body = v
	return s
}

func (s *CloudGetQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
