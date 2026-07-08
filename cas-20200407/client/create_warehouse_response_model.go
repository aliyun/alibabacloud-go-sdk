// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *CreateWarehouseResponseBody) *CreateWarehouseResponse
	GetBody() *CreateWarehouseResponseBody
}

type CreateWarehouseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseResponse) GoString() string {
	return s.String()
}

func (s *CreateWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWarehouseResponse) GetBody() *CreateWarehouseResponseBody {
	return s.Body
}

func (s *CreateWarehouseResponse) SetHeaders(v map[string]*string) *CreateWarehouseResponse {
	s.Headers = v
	return s
}

func (s *CreateWarehouseResponse) SetStatusCode(v int32) *CreateWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarehouseResponse) SetBody(v *CreateWarehouseResponseBody) *CreateWarehouseResponse {
	s.Body = v
	return s
}

func (s *CreateWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
