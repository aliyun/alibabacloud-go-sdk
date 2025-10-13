// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebalanceHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebalanceHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *RebalanceHoloWarehouseResponseBody) *RebalanceHoloWarehouseResponse
	GetBody() *RebalanceHoloWarehouseResponseBody
}

type RebalanceHoloWarehouseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebalanceHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebalanceHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s RebalanceHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebalanceHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebalanceHoloWarehouseResponse) GetBody() *RebalanceHoloWarehouseResponseBody {
	return s.Body
}

func (s *RebalanceHoloWarehouseResponse) SetHeaders(v map[string]*string) *RebalanceHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RebalanceHoloWarehouseResponse) SetStatusCode(v int32) *RebalanceHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceHoloWarehouseResponse) SetBody(v *RebalanceHoloWarehouseResponseBody) *RebalanceHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *RebalanceHoloWarehouseResponse) Validate() error {
	return dara.Validate(s)
}
