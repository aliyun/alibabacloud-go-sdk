// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQueuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQueuesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQueuesResponseBody) *DeleteQueuesResponse
	GetBody() *DeleteQueuesResponseBody
}

type DeleteQueuesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueuesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueuesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQueuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQueuesResponse) GetBody() *DeleteQueuesResponseBody {
	return s.Body
}

func (s *DeleteQueuesResponse) SetHeaders(v map[string]*string) *DeleteQueuesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueuesResponse) SetStatusCode(v int32) *DeleteQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueuesResponse) SetBody(v *DeleteQueuesResponseBody) *DeleteQueuesResponse {
	s.Body = v
	return s
}

func (s *DeleteQueuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
