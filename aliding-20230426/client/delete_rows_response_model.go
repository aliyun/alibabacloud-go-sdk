// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRowsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRowsResponseBody) *DeleteRowsResponse
	GetBody() *DeleteRowsResponseBody
}

type DeleteRowsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRowsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRowsResponse) GetBody() *DeleteRowsResponseBody {
	return s.Body
}

func (s *DeleteRowsResponse) SetHeaders(v map[string]*string) *DeleteRowsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRowsResponse) SetStatusCode(v int32) *DeleteRowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRowsResponse) SetBody(v *DeleteRowsResponseBody) *DeleteRowsResponse {
	s.Body = v
	return s
}

func (s *DeleteRowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
