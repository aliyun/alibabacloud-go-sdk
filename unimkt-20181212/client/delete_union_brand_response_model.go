// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnionBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUnionBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUnionBrandResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUnionBrandResponseBody) *DeleteUnionBrandResponse
	GetBody() *DeleteUnionBrandResponseBody
}

type DeleteUnionBrandResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUnionBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUnionBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnionBrandResponse) GoString() string {
	return s.String()
}

func (s *DeleteUnionBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUnionBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUnionBrandResponse) GetBody() *DeleteUnionBrandResponseBody {
	return s.Body
}

func (s *DeleteUnionBrandResponse) SetHeaders(v map[string]*string) *DeleteUnionBrandResponse {
	s.Headers = v
	return s
}

func (s *DeleteUnionBrandResponse) SetStatusCode(v int32) *DeleteUnionBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUnionBrandResponse) SetBody(v *DeleteUnionBrandResponseBody) *DeleteUnionBrandResponse {
	s.Body = v
	return s
}

func (s *DeleteUnionBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
