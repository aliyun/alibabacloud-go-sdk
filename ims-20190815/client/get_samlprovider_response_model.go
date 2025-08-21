// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSAMLProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSAMLProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetSAMLProviderResponseBody) *GetSAMLProviderResponse
	GetBody() *GetSAMLProviderResponseBody
}

type GetSAMLProviderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSAMLProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSAMLProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSAMLProviderResponse) GetBody() *GetSAMLProviderResponseBody {
	return s.Body
}

func (s *GetSAMLProviderResponse) SetHeaders(v map[string]*string) *GetSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *GetSAMLProviderResponse) SetStatusCode(v int32) *GetSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSAMLProviderResponse) SetBody(v *GetSAMLProviderResponseBody) *GetSAMLProviderResponse {
	s.Body = v
	return s
}

func (s *GetSAMLProviderResponse) Validate() error {
	return dara.Validate(s)
}
