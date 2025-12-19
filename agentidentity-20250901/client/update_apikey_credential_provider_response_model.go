// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAPIKeyCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAPIKeyCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAPIKeyCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAPIKeyCredentialProviderResponseBody) *UpdateAPIKeyCredentialProviderResponse
	GetBody() *UpdateAPIKeyCredentialProviderResponseBody
}

type UpdateAPIKeyCredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAPIKeyCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAPIKeyCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAPIKeyCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateAPIKeyCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAPIKeyCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAPIKeyCredentialProviderResponse) GetBody() *UpdateAPIKeyCredentialProviderResponseBody {
	return s.Body
}

func (s *UpdateAPIKeyCredentialProviderResponse) SetHeaders(v map[string]*string) *UpdateAPIKeyCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateAPIKeyCredentialProviderResponse) SetStatusCode(v int32) *UpdateAPIKeyCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAPIKeyCredentialProviderResponse) SetBody(v *UpdateAPIKeyCredentialProviderResponseBody) *UpdateAPIKeyCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateAPIKeyCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
