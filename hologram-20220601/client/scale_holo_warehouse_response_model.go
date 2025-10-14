// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScaleHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScaleHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *ScaleHoloWarehouseResponseBody) *ScaleHoloWarehouseResponse
	GetBody() *ScaleHoloWarehouseResponseBody
}

type ScaleHoloWarehouseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s ScaleHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScaleHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScaleHoloWarehouseResponse) GetBody() *ScaleHoloWarehouseResponseBody {
	return s.Body
}

func (s *ScaleHoloWarehouseResponse) SetHeaders(v map[string]*string) *ScaleHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ScaleHoloWarehouseResponse) SetStatusCode(v int32) *ScaleHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleHoloWarehouseResponse) SetBody(v *ScaleHoloWarehouseResponseBody) *ScaleHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *ScaleHoloWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
