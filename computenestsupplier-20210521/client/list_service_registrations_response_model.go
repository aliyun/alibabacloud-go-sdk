// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRegistrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceRegistrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceRegistrationsResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceRegistrationsResponseBody) *ListServiceRegistrationsResponse
	GetBody() *ListServiceRegistrationsResponseBody
}

type ListServiceRegistrationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceRegistrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceRegistrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceRegistrationsResponse) GetBody() *ListServiceRegistrationsResponseBody {
	return s.Body
}

func (s *ListServiceRegistrationsResponse) SetHeaders(v map[string]*string) *ListServiceRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceRegistrationsResponse) SetStatusCode(v int32) *ListServiceRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceRegistrationsResponse) SetBody(v *ListServiceRegistrationsResponseBody) *ListServiceRegistrationsResponse {
	s.Body = v
	return s
}

func (s *ListServiceRegistrationsResponse) Validate() error {
	return dara.Validate(s)
}
