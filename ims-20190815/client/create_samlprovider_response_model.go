// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSAMLProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSAMLProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSAMLProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateSAMLProviderResponseBody) *CreateSAMLProviderResponse
	GetBody() *CreateSAMLProviderResponseBody
}

type CreateSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSAMLProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSAMLProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSAMLProviderResponse) GetBody() *CreateSAMLProviderResponseBody {
	return s.Body
}

func (s *CreateSAMLProviderResponse) SetHeaders(v map[string]*string) *CreateSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateSAMLProviderResponse) SetStatusCode(v int32) *CreateSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSAMLProviderResponse) SetBody(v *CreateSAMLProviderResponseBody) *CreateSAMLProviderResponse {
	s.Body = v
	return s
}

func (s *CreateSAMLProviderResponse) Validate() error {
	return dara.Validate(s)
}
