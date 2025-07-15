// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigADConnectorUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigADConnectorUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigADConnectorUserResponse
	GetStatusCode() *int32
	SetBody(v *ConfigADConnectorUserResponseBody) *ConfigADConnectorUserResponse
	GetBody() *ConfigADConnectorUserResponseBody
}

type ConfigADConnectorUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigADConnectorUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigADConnectorUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigADConnectorUserResponse) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigADConnectorUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigADConnectorUserResponse) GetBody() *ConfigADConnectorUserResponseBody {
	return s.Body
}

func (s *ConfigADConnectorUserResponse) SetHeaders(v map[string]*string) *ConfigADConnectorUserResponse {
	s.Headers = v
	return s
}

func (s *ConfigADConnectorUserResponse) SetStatusCode(v int32) *ConfigADConnectorUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigADConnectorUserResponse) SetBody(v *ConfigADConnectorUserResponseBody) *ConfigADConnectorUserResponse {
	s.Body = v
	return s
}

func (s *ConfigADConnectorUserResponse) Validate() error {
	return dara.Validate(s)
}
