// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationClientSecretExpirationTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationClientSecretExpirationTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationClientSecretExpirationTimeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationClientSecretExpirationTimeResponseBody) *UpdateApplicationClientSecretExpirationTimeResponse
	GetBody() *UpdateApplicationClientSecretExpirationTimeResponseBody
}

type UpdateApplicationClientSecretExpirationTimeResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationClientSecretExpirationTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationClientSecretExpirationTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationClientSecretExpirationTimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) GetBody() *UpdateApplicationClientSecretExpirationTimeResponseBody {
	return s.Body
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) SetHeaders(v map[string]*string) *UpdateApplicationClientSecretExpirationTimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) SetStatusCode(v int32) *UpdateApplicationClientSecretExpirationTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) SetBody(v *UpdateApplicationClientSecretExpirationTimeResponseBody) *UpdateApplicationClientSecretExpirationTimeResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
