// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetOIDCProviderResponseBody) *GetOIDCProviderResponse
	GetBody() *GetOIDCProviderResponseBody
}

type GetOIDCProviderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOIDCProviderResponse) GetBody() *GetOIDCProviderResponseBody {
	return s.Body
}

func (s *GetOIDCProviderResponse) SetHeaders(v map[string]*string) *GetOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *GetOIDCProviderResponse) SetStatusCode(v int32) *GetOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOIDCProviderResponse) SetBody(v *GetOIDCProviderResponseBody) *GetOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *GetOIDCProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
