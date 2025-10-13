// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *RestartHoloWarehouseResponseBody) *RestartHoloWarehouseResponse
	GetBody() *RestartHoloWarehouseResponseBody
}

type RestartHoloWarehouseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartHoloWarehouseResponse) GetBody() *RestartHoloWarehouseResponseBody {
	return s.Body
}

func (s *RestartHoloWarehouseResponse) SetHeaders(v map[string]*string) *RestartHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RestartHoloWarehouseResponse) SetStatusCode(v int32) *RestartHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartHoloWarehouseResponse) SetBody(v *RestartHoloWarehouseResponseBody) *RestartHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *RestartHoloWarehouseResponse) Validate() error {
	return dara.Validate(s)
}
