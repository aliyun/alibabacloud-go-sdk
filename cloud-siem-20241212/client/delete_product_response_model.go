// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProductResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProductResponseBody) *DeleteProductResponse
	GetBody() *DeleteProductResponseBody
}

type DeleteProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProductResponse) GetBody() *DeleteProductResponseBody {
	return s.Body
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetStatusCode(v int32) *DeleteProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

func (s *DeleteProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
