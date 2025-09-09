// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListServicesResponseBody) *ListServicesResponse
	GetBody() *ListServicesResponseBody
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServicesResponse) GetBody() *ListServicesResponseBody {
	return s.Body
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

func (s *ListServicesResponse) Validate() error {
	return dara.Validate(s)
}
