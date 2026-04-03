// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChunksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChunksResponseBody) *DeleteChunksResponse
	GetBody() *DeleteChunksResponseBody
}

type DeleteChunksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunksResponse) GoString() string {
	return s.String()
}

func (s *DeleteChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChunksResponse) GetBody() *DeleteChunksResponseBody {
	return s.Body
}

func (s *DeleteChunksResponse) SetHeaders(v map[string]*string) *DeleteChunksResponse {
	s.Headers = v
	return s
}

func (s *DeleteChunksResponse) SetStatusCode(v int32) *DeleteChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChunksResponse) SetBody(v *DeleteChunksResponseBody) *DeleteChunksResponse {
	s.Body = v
	return s
}

func (s *DeleteChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
