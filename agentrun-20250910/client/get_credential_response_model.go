// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CredentialResult) *GetCredentialResponse
	GetBody() *CredentialResult
}

type GetCredentialResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCredentialResponse) GetBody() *CredentialResult {
	return s.Body
}

func (s *GetCredentialResponse) SetHeaders(v map[string]*string) *GetCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialResponse) SetStatusCode(v int32) *GetCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialResponse) SetBody(v *CredentialResult) *GetCredentialResponse {
	s.Body = v
	return s
}

func (s *GetCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
