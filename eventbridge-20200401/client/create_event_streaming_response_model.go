// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventStreamingResponseBody) *CreateEventStreamingResponse
	GetBody() *CreateEventStreamingResponseBody
}

type CreateEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventStreamingResponse) GetBody() *CreateEventStreamingResponseBody {
	return s.Body
}

func (s *CreateEventStreamingResponse) SetHeaders(v map[string]*string) *CreateEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *CreateEventStreamingResponse) SetStatusCode(v int32) *CreateEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventStreamingResponse) SetBody(v *CreateEventStreamingResponseBody) *CreateEventStreamingResponse {
	s.Body = v
	return s
}

func (s *CreateEventStreamingResponse) Validate() error {
	return dara.Validate(s)
}
