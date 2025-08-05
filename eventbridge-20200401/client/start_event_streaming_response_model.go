// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *StartEventStreamingResponseBody) *StartEventStreamingResponse
	GetBody() *StartEventStreamingResponseBody
}

type StartEventStreamingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *StartEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEventStreamingResponse) GetBody() *StartEventStreamingResponseBody {
	return s.Body
}

func (s *StartEventStreamingResponse) SetHeaders(v map[string]*string) *StartEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *StartEventStreamingResponse) SetStatusCode(v int32) *StartEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEventStreamingResponse) SetBody(v *StartEventStreamingResponseBody) *StartEventStreamingResponse {
	s.Body = v
	return s
}

func (s *StartEventStreamingResponse) Validate() error {
	return dara.Validate(s)
}
