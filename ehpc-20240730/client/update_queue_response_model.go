// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQueueResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQueueResponseBody) *UpdateQueueResponse
	GetBody() *UpdateQueueResponseBody
}

type UpdateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueueResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQueueResponse) GetBody() *UpdateQueueResponseBody {
	return s.Body
}

func (s *UpdateQueueResponse) SetHeaders(v map[string]*string) *UpdateQueueResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueueResponse) SetStatusCode(v int32) *UpdateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQueueResponse) SetBody(v *UpdateQueueResponseBody) *UpdateQueueResponse {
	s.Body = v
	return s
}

func (s *UpdateQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
