// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacUserCertStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacUserCertStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacUserCertStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacUserCertStatusResponseBody) *UpdateNacUserCertStatusResponse
	GetBody() *UpdateNacUserCertStatusResponseBody
}

type UpdateNacUserCertStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacUserCertStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacUserCertStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacUserCertStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacUserCertStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacUserCertStatusResponse) GetBody() *UpdateNacUserCertStatusResponseBody {
	return s.Body
}

func (s *UpdateNacUserCertStatusResponse) SetHeaders(v map[string]*string) *UpdateNacUserCertStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacUserCertStatusResponse) SetStatusCode(v int32) *UpdateNacUserCertStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacUserCertStatusResponse) SetBody(v *UpdateNacUserCertStatusResponseBody) *UpdateNacUserCertStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateNacUserCertStatusResponse) Validate() error {
	return dara.Validate(s)
}
