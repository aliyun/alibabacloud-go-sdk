// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSAMLProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSAMLProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSAMLProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSAMLProviderResponseBody) *UpdateSAMLProviderResponse
	GetBody() *UpdateSAMLProviderResponseBody
}

type UpdateSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSAMLProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSAMLProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSAMLProviderResponse) GetBody() *UpdateSAMLProviderResponseBody {
	return s.Body
}

func (s *UpdateSAMLProviderResponse) SetHeaders(v map[string]*string) *UpdateSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateSAMLProviderResponse) SetStatusCode(v int32) *UpdateSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSAMLProviderResponse) SetBody(v *UpdateSAMLProviderResponseBody) *UpdateSAMLProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateSAMLProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
