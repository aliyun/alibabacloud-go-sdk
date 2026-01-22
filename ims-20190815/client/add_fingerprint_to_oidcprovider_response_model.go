// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFingerprintToOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFingerprintToOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFingerprintToOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *AddFingerprintToOIDCProviderResponseBody) *AddFingerprintToOIDCProviderResponse
	GetBody() *AddFingerprintToOIDCProviderResponseBody
}

type AddFingerprintToOIDCProviderResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFingerprintToOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFingerprintToOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFingerprintToOIDCProviderResponse) GetBody() *AddFingerprintToOIDCProviderResponseBody {
	return s.Body
}

func (s *AddFingerprintToOIDCProviderResponse) SetHeaders(v map[string]*string) *AddFingerprintToOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *AddFingerprintToOIDCProviderResponse) SetStatusCode(v int32) *AddFingerprintToOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponse) SetBody(v *AddFingerprintToOIDCProviderResponseBody) *AddFingerprintToOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *AddFingerprintToOIDCProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
