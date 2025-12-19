// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAPIKeyCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAPIKeyCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAPIKeyCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAPIKeyCredentialProviderResponseBody) *DeleteAPIKeyCredentialProviderResponse
	GetBody() *DeleteAPIKeyCredentialProviderResponseBody
}

type DeleteAPIKeyCredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAPIKeyCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAPIKeyCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAPIKeyCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteAPIKeyCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAPIKeyCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAPIKeyCredentialProviderResponse) GetBody() *DeleteAPIKeyCredentialProviderResponseBody {
	return s.Body
}

func (s *DeleteAPIKeyCredentialProviderResponse) SetHeaders(v map[string]*string) *DeleteAPIKeyCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteAPIKeyCredentialProviderResponse) SetStatusCode(v int32) *DeleteAPIKeyCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAPIKeyCredentialProviderResponse) SetBody(v *DeleteAPIKeyCredentialProviderResponseBody) *DeleteAPIKeyCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteAPIKeyCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
