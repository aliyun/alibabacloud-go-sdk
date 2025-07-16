// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceServicesResponseBody) *ListResourceServicesResponse
	GetBody() *ListResourceServicesResponseBody
}

type ListResourceServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServicesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceServicesResponse) GetBody() *ListResourceServicesResponseBody {
	return s.Body
}

func (s *ListResourceServicesResponse) SetHeaders(v map[string]*string) *ListResourceServicesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceServicesResponse) SetStatusCode(v int32) *ListResourceServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceServicesResponse) SetBody(v *ListResourceServicesResponseBody) *ListResourceServicesResponse {
	s.Body = v
	return s
}

func (s *ListResourceServicesResponse) Validate() error {
	return dara.Validate(s)
}
