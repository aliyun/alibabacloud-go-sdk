// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMixStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMixStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMixStreamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMixStreamResponseBody) *DeleteMixStreamResponse
	GetBody() *DeleteMixStreamResponseBody
}

type DeleteMixStreamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMixStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMixStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMixStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteMixStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMixStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMixStreamResponse) GetBody() *DeleteMixStreamResponseBody {
	return s.Body
}

func (s *DeleteMixStreamResponse) SetHeaders(v map[string]*string) *DeleteMixStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteMixStreamResponse) SetStatusCode(v int32) *DeleteMixStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMixStreamResponse) SetBody(v *DeleteMixStreamResponseBody) *DeleteMixStreamResponse {
	s.Body = v
	return s
}

func (s *DeleteMixStreamResponse) Validate() error {
	return dara.Validate(s)
}
