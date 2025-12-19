// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPIKeyCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAPIKeyCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAPIKeyCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateAPIKeyCredentialProviderResponseBody) *CreateAPIKeyCredentialProviderResponse
	GetBody() *CreateAPIKeyCredentialProviderResponseBody
}

type CreateAPIKeyCredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAPIKeyCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAPIKeyCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAPIKeyCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateAPIKeyCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAPIKeyCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAPIKeyCredentialProviderResponse) GetBody() *CreateAPIKeyCredentialProviderResponseBody {
	return s.Body
}

func (s *CreateAPIKeyCredentialProviderResponse) SetHeaders(v map[string]*string) *CreateAPIKeyCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponse) SetStatusCode(v int32) *CreateAPIKeyCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponse) SetBody(v *CreateAPIKeyCredentialProviderResponseBody) *CreateAPIKeyCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
