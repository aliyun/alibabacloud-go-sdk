// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKyuubiServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKyuubiServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListKyuubiServicesResponseBody) *ListKyuubiServicesResponse
	GetBody() *ListKyuubiServicesResponseBody
}

type ListKyuubiServicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiServicesResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKyuubiServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKyuubiServicesResponse) GetBody() *ListKyuubiServicesResponseBody {
	return s.Body
}

func (s *ListKyuubiServicesResponse) SetHeaders(v map[string]*string) *ListKyuubiServicesResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiServicesResponse) SetStatusCode(v int32) *ListKyuubiServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiServicesResponse) SetBody(v *ListKyuubiServicesResponseBody) *ListKyuubiServicesResponse {
	s.Body = v
	return s
}

func (s *ListKyuubiServicesResponse) Validate() error {
	return dara.Validate(s)
}
