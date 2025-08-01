// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEurekaServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEurekaServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListEurekaServicesResponseBody) *ListEurekaServicesResponse
	GetBody() *ListEurekaServicesResponseBody
}

type ListEurekaServicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEurekaServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEurekaServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaServicesResponse) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEurekaServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEurekaServicesResponse) GetBody() *ListEurekaServicesResponseBody {
	return s.Body
}

func (s *ListEurekaServicesResponse) SetHeaders(v map[string]*string) *ListEurekaServicesResponse {
	s.Headers = v
	return s
}

func (s *ListEurekaServicesResponse) SetStatusCode(v int32) *ListEurekaServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEurekaServicesResponse) SetBody(v *ListEurekaServicesResponseBody) *ListEurekaServicesResponse {
	s.Body = v
	return s
}

func (s *ListEurekaServicesResponse) Validate() error {
	return dara.Validate(s)
}
