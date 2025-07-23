// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueueResponse
	GetStatusCode() *int32
	SetBody(v *ListQueueResponseBody) *ListQueueResponse
	GetBody() *ListQueueResponseBody
}

type ListQueueResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueueResponse) GoString() string {
	return s.String()
}

func (s *ListQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueueResponse) GetBody() *ListQueueResponseBody {
	return s.Body
}

func (s *ListQueueResponse) SetHeaders(v map[string]*string) *ListQueueResponse {
	s.Headers = v
	return s
}

func (s *ListQueueResponse) SetStatusCode(v int32) *ListQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueResponse) SetBody(v *ListQueueResponseBody) *ListQueueResponse {
	s.Body = v
	return s
}

func (s *ListQueueResponse) Validate() error {
	return dara.Validate(s)
}
