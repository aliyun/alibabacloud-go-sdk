// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSAMLProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSAMLProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSAMLProviderResponseBody) *DeleteSAMLProviderResponse
	GetBody() *DeleteSAMLProviderResponseBody
}

type DeleteSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSAMLProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSAMLProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSAMLProviderResponse) GetBody() *DeleteSAMLProviderResponseBody {
	return s.Body
}

func (s *DeleteSAMLProviderResponse) SetHeaders(v map[string]*string) *DeleteSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteSAMLProviderResponse) SetStatusCode(v int32) *DeleteSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSAMLProviderResponse) SetBody(v *DeleteSAMLProviderResponseBody) *DeleteSAMLProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteSAMLProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
