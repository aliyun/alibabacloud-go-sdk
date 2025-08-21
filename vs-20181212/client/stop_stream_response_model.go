// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopStreamResponse
	GetStatusCode() *int32
	SetBody(v *StopStreamResponseBody) *StopStreamResponse
	GetBody() *StopStreamResponseBody
}

type StopStreamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StopStreamResponse) GoString() string {
	return s.String()
}

func (s *StopStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopStreamResponse) GetBody() *StopStreamResponseBody {
	return s.Body
}

func (s *StopStreamResponse) SetHeaders(v map[string]*string) *StopStreamResponse {
	s.Headers = v
	return s
}

func (s *StopStreamResponse) SetStatusCode(v int32) *StopStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStreamResponse) SetBody(v *StopStreamResponseBody) *StopStreamResponse {
	s.Body = v
	return s
}

func (s *StopStreamResponse) Validate() error {
	return dara.Validate(s)
}
