// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWarehouseResponseBody) *DeleteWarehouseResponse
	GetBody() *DeleteWarehouseResponseBody
}

type DeleteWarehouseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWarehouseResponse) GetBody() *DeleteWarehouseResponseBody {
	return s.Body
}

func (s *DeleteWarehouseResponse) SetHeaders(v map[string]*string) *DeleteWarehouseResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarehouseResponse) SetStatusCode(v int32) *DeleteWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarehouseResponse) SetBody(v *DeleteWarehouseResponseBody) *DeleteWarehouseResponse {
	s.Body = v
	return s
}

func (s *DeleteWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
