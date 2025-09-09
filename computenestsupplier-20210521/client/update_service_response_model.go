// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse
	GetBody() *UpdateServiceResponseBody
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceResponse) GetBody() *UpdateServiceResponseBody {
	return s.Body
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceResponse) Validate() error {
	return dara.Validate(s)
}
