// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNumbersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNumbersResponseBody) *DeleteNumbersResponse
	GetBody() *DeleteNumbersResponseBody
}

type DeleteNumbersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNumbersResponse) GoString() string {
	return s.String()
}

func (s *DeleteNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNumbersResponse) GetBody() *DeleteNumbersResponseBody {
	return s.Body
}

func (s *DeleteNumbersResponse) SetHeaders(v map[string]*string) *DeleteNumbersResponse {
	s.Headers = v
	return s
}

func (s *DeleteNumbersResponse) SetStatusCode(v int32) *DeleteNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNumbersResponse) SetBody(v *DeleteNumbersResponseBody) *DeleteNumbersResponse {
	s.Body = v
	return s
}

func (s *DeleteNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
