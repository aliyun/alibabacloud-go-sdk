// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoloWarehouseResponseBody) *CreateHoloWarehouseResponse
	GetBody() *CreateHoloWarehouseResponseBody
}

type CreateHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoloWarehouseResponse) GetBody() *CreateHoloWarehouseResponseBody {
	return s.Body
}

func (s *CreateHoloWarehouseResponse) SetHeaders(v map[string]*string) *CreateHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *CreateHoloWarehouseResponse) SetStatusCode(v int32) *CreateHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoloWarehouseResponse) SetBody(v *CreateHoloWarehouseResponseBody) *CreateHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *CreateHoloWarehouseResponse) Validate() error {
	return dara.Validate(s)
}
