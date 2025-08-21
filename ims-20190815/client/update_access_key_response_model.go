// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccessKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccessKeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccessKeyResponseBody) *UpdateAccessKeyResponse
	GetBody() *UpdateAccessKeyResponseBody
}

type UpdateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccessKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccessKeyResponse) GetBody() *UpdateAccessKeyResponseBody {
	return s.Body
}

func (s *UpdateAccessKeyResponse) SetHeaders(v map[string]*string) *UpdateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessKeyResponse) SetStatusCode(v int32) *UpdateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessKeyResponse) SetBody(v *UpdateAccessKeyResponseBody) *UpdateAccessKeyResponse {
	s.Body = v
	return s
}

func (s *UpdateAccessKeyResponse) Validate() error {
	return dara.Validate(s)
}
