// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse
	GetBody() *UpdateServiceInstanceSpecResponseBody
}

type UpdateServiceInstanceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceInstanceSpecResponse) GetBody() *UpdateServiceInstanceSpecResponseBody {
	return s.Body
}

func (s *UpdateServiceInstanceSpecResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) Validate() error {
	return dara.Validate(s)
}
