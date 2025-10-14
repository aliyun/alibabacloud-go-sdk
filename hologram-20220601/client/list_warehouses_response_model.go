// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehousesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarehousesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarehousesResponse
	GetStatusCode() *int32
	SetBody(v *ListWarehousesResponseBody) *ListWarehousesResponse
	GetBody() *ListWarehousesResponseBody
}

type ListWarehousesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarehousesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarehousesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarehousesResponse) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarehousesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarehousesResponse) GetBody() *ListWarehousesResponseBody {
	return s.Body
}

func (s *ListWarehousesResponse) SetHeaders(v map[string]*string) *ListWarehousesResponse {
	s.Headers = v
	return s
}

func (s *ListWarehousesResponse) SetStatusCode(v int32) *ListWarehousesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarehousesResponse) SetBody(v *ListWarehousesResponseBody) *ListWarehousesResponse {
	s.Body = v
	return s
}

func (s *ListWarehousesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
