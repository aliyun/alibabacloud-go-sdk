// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CredentialResult) *CreateCredentialResponse
	GetBody() *CredentialResult
}

type CreateCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCredentialResponse) GetBody() *CredentialResult {
	return s.Body
}

func (s *CreateCredentialResponse) SetHeaders(v map[string]*string) *CreateCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateCredentialResponse) SetStatusCode(v int32) *CreateCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCredentialResponse) SetBody(v *CredentialResult) *CreateCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
