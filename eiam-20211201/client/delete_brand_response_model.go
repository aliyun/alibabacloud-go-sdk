// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBrandResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBrandResponseBody) *DeleteBrandResponse
	GetBody() *DeleteBrandResponseBody
}

type DeleteBrandResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrandResponse) GoString() string {
	return s.String()
}

func (s *DeleteBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBrandResponse) GetBody() *DeleteBrandResponseBody {
	return s.Body
}

func (s *DeleteBrandResponse) SetHeaders(v map[string]*string) *DeleteBrandResponse {
	s.Headers = v
	return s
}

func (s *DeleteBrandResponse) SetStatusCode(v int32) *DeleteBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBrandResponse) SetBody(v *DeleteBrandResponseBody) *DeleteBrandResponse {
	s.Body = v
	return s
}

func (s *DeleteBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
