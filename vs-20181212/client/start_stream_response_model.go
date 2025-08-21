// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartStreamResponse
	GetStatusCode() *int32
	SetBody(v *StartStreamResponseBody) *StartStreamResponse
	GetBody() *StartStreamResponseBody
}

type StartStreamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StartStreamResponse) GoString() string {
	return s.String()
}

func (s *StartStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartStreamResponse) GetBody() *StartStreamResponseBody {
	return s.Body
}

func (s *StartStreamResponse) SetHeaders(v map[string]*string) *StartStreamResponse {
	s.Headers = v
	return s
}

func (s *StartStreamResponse) SetStatusCode(v int32) *StartStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StartStreamResponse) SetBody(v *StartStreamResponseBody) *StartStreamResponse {
	s.Body = v
	return s
}

func (s *StartStreamResponse) Validate() error {
	return dara.Validate(s)
}
