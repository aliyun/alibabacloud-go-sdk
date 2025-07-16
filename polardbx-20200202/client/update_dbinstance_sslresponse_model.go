// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDBInstanceSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDBInstanceSSLResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDBInstanceSSLResponseBody) *UpdateDBInstanceSSLResponse
	GetBody() *UpdateDBInstanceSSLResponseBody
}

type UpdateDBInstanceSSLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDBInstanceSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDBInstanceSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDBInstanceSSLResponse) GetBody() *UpdateDBInstanceSSLResponseBody {
	return s.Body
}

func (s *UpdateDBInstanceSSLResponse) SetHeaders(v map[string]*string) *UpdateDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstanceSSLResponse) SetStatusCode(v int32) *UpdateDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstanceSSLResponse) SetBody(v *UpdateDBInstanceSSLResponseBody) *UpdateDBInstanceSSLResponse {
	s.Body = v
	return s
}

func (s *UpdateDBInstanceSSLResponse) Validate() error {
	return dara.Validate(s)
}
