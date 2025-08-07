// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChunkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChunkResponseBody) *DeleteChunkResponse
	GetBody() *DeleteChunkResponseBody
}

type DeleteChunkResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunkResponse) GoString() string {
	return s.String()
}

func (s *DeleteChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChunkResponse) GetBody() *DeleteChunkResponseBody {
	return s.Body
}

func (s *DeleteChunkResponse) SetHeaders(v map[string]*string) *DeleteChunkResponse {
	s.Headers = v
	return s
}

func (s *DeleteChunkResponse) SetStatusCode(v int32) *DeleteChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChunkResponse) SetBody(v *DeleteChunkResponseBody) *DeleteChunkResponse {
	s.Body = v
	return s
}

func (s *DeleteChunkResponse) Validate() error {
	return dara.Validate(s)
}
