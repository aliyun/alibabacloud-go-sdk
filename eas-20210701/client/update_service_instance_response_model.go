// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceInstanceResponseBody) *UpdateServiceInstanceResponse
	GetBody() *UpdateServiceInstanceResponseBody
}

type UpdateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceInstanceResponse) GetBody() *UpdateServiceInstanceResponseBody {
	return s.Body
}

func (s *UpdateServiceInstanceResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceResponse) SetStatusCode(v int32) *UpdateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceResponse) SetBody(v *UpdateServiceInstanceResponseBody) *UpdateServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
