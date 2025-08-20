// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypeRegistrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceTypeRegistrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceTypeRegistrationsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceTypeRegistrationsResponseBody) *ListResourceTypeRegistrationsResponse
	GetBody() *ListResourceTypeRegistrationsResponseBody
}

type ListResourceTypeRegistrationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypeRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypeRegistrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceTypeRegistrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceTypeRegistrationsResponse) GetBody() *ListResourceTypeRegistrationsResponseBody {
	return s.Body
}

func (s *ListResourceTypeRegistrationsResponse) SetHeaders(v map[string]*string) *ListResourceTypeRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypeRegistrationsResponse) SetStatusCode(v int32) *ListResourceTypeRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponse) SetBody(v *ListResourceTypeRegistrationsResponseBody) *ListResourceTypeRegistrationsResponse {
	s.Body = v
	return s
}

func (s *ListResourceTypeRegistrationsResponse) Validate() error {
	return dara.Validate(s)
}
