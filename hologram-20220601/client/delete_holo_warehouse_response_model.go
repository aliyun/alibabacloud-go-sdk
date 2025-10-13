// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoloWarehouseResponseBody) *DeleteHoloWarehouseResponse
	GetBody() *DeleteHoloWarehouseResponseBody
}

type DeleteHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoloWarehouseResponse) GetBody() *DeleteHoloWarehouseResponseBody {
	return s.Body
}

func (s *DeleteHoloWarehouseResponse) SetHeaders(v map[string]*string) *DeleteHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoloWarehouseResponse) SetStatusCode(v int32) *DeleteHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoloWarehouseResponse) SetBody(v *DeleteHoloWarehouseResponseBody) *DeleteHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *DeleteHoloWarehouseResponse) Validate() error {
	return dara.Validate(s)
}
