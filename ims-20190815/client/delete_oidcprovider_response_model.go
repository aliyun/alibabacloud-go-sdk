// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOIDCProviderResponseBody) *DeleteOIDCProviderResponse
	GetBody() *DeleteOIDCProviderResponseBody
}

type DeleteOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOIDCProviderResponse) GetBody() *DeleteOIDCProviderResponseBody {
	return s.Body
}

func (s *DeleteOIDCProviderResponse) SetHeaders(v map[string]*string) *DeleteOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteOIDCProviderResponse) SetStatusCode(v int32) *DeleteOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOIDCProviderResponse) SetBody(v *DeleteOIDCProviderResponseBody) *DeleteOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteOIDCProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
