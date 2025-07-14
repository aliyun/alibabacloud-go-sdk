// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppServicesResponseBody) *ListAppServicesResponse
	GetBody() *ListAppServicesResponseBody
}

type ListAppServicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesResponse) GoString() string {
	return s.String()
}

func (s *ListAppServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppServicesResponse) GetBody() *ListAppServicesResponseBody {
	return s.Body
}

func (s *ListAppServicesResponse) SetHeaders(v map[string]*string) *ListAppServicesResponse {
	s.Headers = v
	return s
}

func (s *ListAppServicesResponse) SetStatusCode(v int32) *ListAppServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppServicesResponse) SetBody(v *ListAppServicesResponseBody) *ListAppServicesResponse {
	s.Body = v
	return s
}

func (s *ListAppServicesResponse) Validate() error {
	return dara.Validate(s)
}
