// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCredentialProviderResponseBody) *DeleteCredentialProviderResponse
	GetBody() *DeleteCredentialProviderResponseBody
}

type DeleteCredentialProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCredentialProviderResponse) GetBody() *DeleteCredentialProviderResponseBody {
	return s.Body
}

func (s *DeleteCredentialProviderResponse) SetHeaders(v map[string]*string) *DeleteCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteCredentialProviderResponse) SetStatusCode(v int32) *DeleteCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCredentialProviderResponse) SetBody(v *DeleteCredentialProviderResponseBody) *DeleteCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
