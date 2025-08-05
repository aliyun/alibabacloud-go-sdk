// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventStreamingResponseBody) *UpdateEventStreamingResponse
	GetBody() *UpdateEventStreamingResponseBody
}

type UpdateEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventStreamingResponse) GetBody() *UpdateEventStreamingResponseBody {
	return s.Body
}

func (s *UpdateEventStreamingResponse) SetHeaders(v map[string]*string) *UpdateEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventStreamingResponse) SetStatusCode(v int32) *UpdateEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventStreamingResponse) SetBody(v *UpdateEventStreamingResponseBody) *UpdateEventStreamingResponse {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingResponse) Validate() error {
	return dara.Validate(s)
}
