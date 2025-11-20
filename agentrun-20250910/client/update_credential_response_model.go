// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CredentialResult) *UpdateCredentialResponse
	GetBody() *CredentialResult
}

type UpdateCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialResponse) GetBody() *CredentialResult {
	return s.Body
}

func (s *UpdateCredentialResponse) SetHeaders(v map[string]*string) *UpdateCredentialResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialResponse) SetStatusCode(v int32) *UpdateCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialResponse) SetBody(v *CredentialResult) *UpdateCredentialResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
