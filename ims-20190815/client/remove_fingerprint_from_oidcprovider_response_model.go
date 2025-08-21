// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFingerprintFromOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFingerprintFromOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFingerprintFromOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *RemoveFingerprintFromOIDCProviderResponseBody) *RemoveFingerprintFromOIDCProviderResponse
	GetBody() *RemoveFingerprintFromOIDCProviderResponseBody
}

type RemoveFingerprintFromOIDCProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFingerprintFromOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFingerprintFromOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFingerprintFromOIDCProviderResponse) GetBody() *RemoveFingerprintFromOIDCProviderResponseBody {
	return s.Body
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetHeaders(v map[string]*string) *RemoveFingerprintFromOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetStatusCode(v int32) *RemoveFingerprintFromOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetBody(v *RemoveFingerprintFromOIDCProviderResponseBody) *RemoveFingerprintFromOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponse) Validate() error {
	return dara.Validate(s)
}
