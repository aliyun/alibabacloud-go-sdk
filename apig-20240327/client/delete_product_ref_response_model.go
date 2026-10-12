// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProductRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProductRefResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProductRefResponseBody) *DeleteProductRefResponse
	GetBody() *DeleteProductRefResponseBody
}

type DeleteProductRefResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductRefResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRefResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProductRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProductRefResponse) GetBody() *DeleteProductRefResponseBody {
	return s.Body
}

func (s *DeleteProductRefResponse) SetHeaders(v map[string]*string) *DeleteProductRefResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductRefResponse) SetStatusCode(v int32) *DeleteProductRefResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductRefResponse) SetBody(v *DeleteProductRefResponseBody) *DeleteProductRefResponse {
	s.Body = v
	return s
}

func (s *DeleteProductRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
