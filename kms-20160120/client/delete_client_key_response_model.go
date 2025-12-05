// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientKeyResponseBody) *DeleteClientKeyResponse
	GetBody() *DeleteClientKeyResponseBody
}

type DeleteClientKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientKeyResponse) GetBody() *DeleteClientKeyResponseBody {
	return s.Body
}

func (s *DeleteClientKeyResponse) SetHeaders(v map[string]*string) *DeleteClientKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientKeyResponse) SetStatusCode(v int32) *DeleteClientKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientKeyResponse) SetBody(v *DeleteClientKeyResponseBody) *DeleteClientKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteClientKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
