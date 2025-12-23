// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertWarehouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCertWarehouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCertWarehouseResponse
	GetStatusCode() *int32
	SetBody(v *ListCertWarehouseResponseBody) *ListCertWarehouseResponse
	GetBody() *ListCertWarehouseResponseBody
}

type ListCertWarehouseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertWarehouseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCertWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ListCertWarehouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCertWarehouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCertWarehouseResponse) GetBody() *ListCertWarehouseResponseBody {
	return s.Body
}

func (s *ListCertWarehouseResponse) SetHeaders(v map[string]*string) *ListCertWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ListCertWarehouseResponse) SetStatusCode(v int32) *ListCertWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertWarehouseResponse) SetBody(v *ListCertWarehouseResponseBody) *ListCertWarehouseResponse {
	s.Body = v
	return s
}

func (s *ListCertWarehouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
