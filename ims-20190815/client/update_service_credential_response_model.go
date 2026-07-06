// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceCredentialResponseBody) *UpdateServiceCredentialResponse
	GetBody() *UpdateServiceCredentialResponseBody
}

type UpdateServiceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCredentialResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceCredentialResponse) GetBody() *UpdateServiceCredentialResponseBody {
	return s.Body
}

func (s *UpdateServiceCredentialResponse) SetHeaders(v map[string]*string) *UpdateServiceCredentialResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceCredentialResponse) SetStatusCode(v int32) *UpdateServiceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceCredentialResponse) SetBody(v *UpdateServiceCredentialResponseBody) *UpdateServiceCredentialResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
