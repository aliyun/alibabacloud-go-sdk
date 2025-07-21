// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueuesResponse
	GetStatusCode() *int32
	SetBody(v *ListQueuesResponseBody) *ListQueuesResponse
	GetBody() *ListQueuesResponseBody
}

type ListQueuesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueuesResponse) GetBody() *ListQueuesResponseBody {
	return s.Body
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetStatusCode(v int32) *ListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

func (s *ListQueuesResponse) Validate() error {
	return dara.Validate(s)
}
