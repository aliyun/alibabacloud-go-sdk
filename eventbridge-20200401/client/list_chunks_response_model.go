// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChunksResponse
	GetStatusCode() *int32
	SetBody(v *ListChunksResponseBody) *ListChunksResponse
	GetBody() *ListChunksResponseBody
}

type ListChunksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponse) GoString() string {
	return s.String()
}

func (s *ListChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChunksResponse) GetBody() *ListChunksResponseBody {
	return s.Body
}

func (s *ListChunksResponse) SetHeaders(v map[string]*string) *ListChunksResponse {
	s.Headers = v
	return s
}

func (s *ListChunksResponse) SetStatusCode(v int32) *ListChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChunksResponse) SetBody(v *ListChunksResponseBody) *ListChunksResponse {
	s.Body = v
	return s
}

func (s *ListChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
