// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataEventServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataEventServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataEventServicesResponseBody) *ListDataEventServicesResponse
	GetBody() *ListDataEventServicesResponseBody
}

type ListDataEventServicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataEventServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataEventServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventServicesResponse) GoString() string {
	return s.String()
}

func (s *ListDataEventServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataEventServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataEventServicesResponse) GetBody() *ListDataEventServicesResponseBody {
	return s.Body
}

func (s *ListDataEventServicesResponse) SetHeaders(v map[string]*string) *ListDataEventServicesResponse {
	s.Headers = v
	return s
}

func (s *ListDataEventServicesResponse) SetStatusCode(v int32) *ListDataEventServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataEventServicesResponse) SetBody(v *ListDataEventServicesResponseBody) *ListDataEventServicesResponse {
	s.Body = v
	return s
}

func (s *ListDataEventServicesResponse) Validate() error {
	return dara.Validate(s)
}
