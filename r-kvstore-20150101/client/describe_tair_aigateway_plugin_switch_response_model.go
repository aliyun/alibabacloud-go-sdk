// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayPluginSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairAIGatewayPluginSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairAIGatewayPluginSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairAIGatewayPluginSwitchResponseBody) *DescribeTairAIGatewayPluginSwitchResponse
	GetBody() *DescribeTairAIGatewayPluginSwitchResponseBody
}

type DescribeTairAIGatewayPluginSwitchResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairAIGatewayPluginSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairAIGatewayPluginSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayPluginSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) GetBody() *DescribeTairAIGatewayPluginSwitchResponseBody {
	return s.Body
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) SetHeaders(v map[string]*string) *DescribeTairAIGatewayPluginSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) SetStatusCode(v int32) *DescribeTairAIGatewayPluginSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) SetBody(v *DescribeTairAIGatewayPluginSwitchResponseBody) *DescribeTairAIGatewayPluginSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
