// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventStreamingResponseBody) *DeleteEventStreamingResponse
	GetBody() *DeleteEventStreamingResponseBody
}

type DeleteEventStreamingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventStreamingResponse) GetBody() *DeleteEventStreamingResponseBody {
	return s.Body
}

func (s *DeleteEventStreamingResponse) SetHeaders(v map[string]*string) *DeleteEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventStreamingResponse) SetStatusCode(v int32) *DeleteEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventStreamingResponse) SetBody(v *DeleteEventStreamingResponseBody) *DeleteEventStreamingResponse {
	s.Body = v
	return s
}

func (s *DeleteEventStreamingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
