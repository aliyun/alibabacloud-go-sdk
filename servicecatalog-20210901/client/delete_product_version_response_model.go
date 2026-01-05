// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProductVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProductVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProductVersionResponseBody) *DeleteProductVersionResponse
	GetBody() *DeleteProductVersionResponseBody
}

type DeleteProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProductVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProductVersionResponse) GetBody() *DeleteProductVersionResponseBody {
	return s.Body
}

func (s *DeleteProductVersionResponse) SetHeaders(v map[string]*string) *DeleteProductVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductVersionResponse) SetStatusCode(v int32) *DeleteProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductVersionResponse) SetBody(v *DeleteProductVersionResponseBody) *DeleteProductVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteProductVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
