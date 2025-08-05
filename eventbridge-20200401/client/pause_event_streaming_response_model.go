// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *PauseEventStreamingResponseBody) *PauseEventStreamingResponse
	GetBody() *PauseEventStreamingResponseBody
}

type PauseEventStreamingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseEventStreamingResponse) GetBody() *PauseEventStreamingResponseBody {
	return s.Body
}

func (s *PauseEventStreamingResponse) SetHeaders(v map[string]*string) *PauseEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *PauseEventStreamingResponse) SetStatusCode(v int32) *PauseEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseEventStreamingResponse) SetBody(v *PauseEventStreamingResponseBody) *PauseEventStreamingResponse {
	s.Body = v
	return s
}

func (s *PauseEventStreamingResponse) Validate() error {
	return dara.Validate(s)
}
