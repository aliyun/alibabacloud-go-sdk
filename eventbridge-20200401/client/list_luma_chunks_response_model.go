// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaChunksResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaChunksResponseBody) *ListLumaChunksResponse
	GetBody() *ListLumaChunksResponseBody
}

type ListLumaChunksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaChunksResponse) GoString() string {
	return s.String()
}

func (s *ListLumaChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaChunksResponse) GetBody() *ListLumaChunksResponseBody {
	return s.Body
}

func (s *ListLumaChunksResponse) SetHeaders(v map[string]*string) *ListLumaChunksResponse {
	s.Headers = v
	return s
}

func (s *ListLumaChunksResponse) SetStatusCode(v int32) *ListLumaChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaChunksResponse) SetBody(v *ListLumaChunksResponseBody) *ListLumaChunksResponse {
	s.Body = v
	return s
}

func (s *ListLumaChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
