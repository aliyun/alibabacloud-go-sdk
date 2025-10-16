// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigADConnectorTrustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigADConnectorTrustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigADConnectorTrustResponse
	GetStatusCode() *int32
	SetBody(v *ConfigADConnectorTrustResponseBody) *ConfigADConnectorTrustResponse
	GetBody() *ConfigADConnectorTrustResponseBody
}

type ConfigADConnectorTrustResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigADConnectorTrustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigADConnectorTrustResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigADConnectorTrustResponse) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigADConnectorTrustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigADConnectorTrustResponse) GetBody() *ConfigADConnectorTrustResponseBody {
	return s.Body
}

func (s *ConfigADConnectorTrustResponse) SetHeaders(v map[string]*string) *ConfigADConnectorTrustResponse {
	s.Headers = v
	return s
}

func (s *ConfigADConnectorTrustResponse) SetStatusCode(v int32) *ConfigADConnectorTrustResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigADConnectorTrustResponse) SetBody(v *ConfigADConnectorTrustResponseBody) *ConfigADConnectorTrustResponse {
	s.Body = v
	return s
}

func (s *ConfigADConnectorTrustResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
