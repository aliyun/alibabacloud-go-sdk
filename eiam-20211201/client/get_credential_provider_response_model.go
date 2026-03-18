// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetCredentialProviderResponseBody) *GetCredentialProviderResponse
	GetBody() *GetCredentialProviderResponseBody
}

type GetCredentialProviderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCredentialProviderResponse) GetBody() *GetCredentialProviderResponseBody {
	return s.Body
}

func (s *GetCredentialProviderResponse) SetHeaders(v map[string]*string) *GetCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialProviderResponse) SetStatusCode(v int32) *GetCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialProviderResponse) SetBody(v *GetCredentialProviderResponseBody) *GetCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *GetCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
