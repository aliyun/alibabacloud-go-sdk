// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigNetworkRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigNetworkRulesResponse
	GetStatusCode() *int32
	SetBody(v *ConfigNetworkRulesResponseBody) *ConfigNetworkRulesResponse
	GetBody() *ConfigNetworkRulesResponseBody
}

type ConfigNetworkRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigNetworkRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigNetworkRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigNetworkRulesResponse) GetBody() *ConfigNetworkRulesResponseBody {
	return s.Body
}

func (s *ConfigNetworkRulesResponse) SetHeaders(v map[string]*string) *ConfigNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetworkRulesResponse) SetStatusCode(v int32) *ConfigNetworkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigNetworkRulesResponse) SetBody(v *ConfigNetworkRulesResponseBody) *ConfigNetworkRulesResponse {
	s.Body = v
	return s
}

func (s *ConfigNetworkRulesResponse) Validate() error {
	return dara.Validate(s)
}
