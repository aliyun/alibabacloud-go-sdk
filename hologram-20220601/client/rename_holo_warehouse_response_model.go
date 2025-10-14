// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameHoloWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameHoloWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameHoloWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *RenameHoloWarehouseResponseBody) *RenameHoloWarehouseResponse
	GetBody() *RenameHoloWarehouseResponseBody
}

type RenameHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameHoloWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameHoloWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameHoloWarehouseResponse) GetBody() *RenameHoloWarehouseResponseBody {
	return s.Body
}

func (s *RenameHoloWarehouseResponse) SetHeaders(v map[string]*string) *RenameHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RenameHoloWarehouseResponse) SetStatusCode(v int32) *RenameHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameHoloWarehouseResponse) SetBody(v *RenameHoloWarehouseResponseBody) *RenameHoloWarehouseResponse {
	s.Body = v
	return s
}

func (s *RenameHoloWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
