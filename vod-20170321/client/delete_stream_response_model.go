// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamResponseBody) *DeleteStreamResponse
	GetBody() *DeleteStreamResponseBody
}

type DeleteStreamResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamResponse) GetBody() *DeleteStreamResponseBody {
	return s.Body
}

func (s *DeleteStreamResponse) SetHeaders(v map[string]*string) *DeleteStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamResponse) SetStatusCode(v int32) *DeleteStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamResponse) SetBody(v *DeleteStreamResponseBody) *DeleteStreamResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
