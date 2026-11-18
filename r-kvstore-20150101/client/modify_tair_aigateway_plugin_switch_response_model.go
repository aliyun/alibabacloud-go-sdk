// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayPluginSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTairAIGatewayPluginSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTairAIGatewayPluginSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTairAIGatewayPluginSwitchResponseBody) *ModifyTairAIGatewayPluginSwitchResponse
	GetBody() *ModifyTairAIGatewayPluginSwitchResponseBody
}

type ModifyTairAIGatewayPluginSwitchResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTairAIGatewayPluginSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTairAIGatewayPluginSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayPluginSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) GetBody() *ModifyTairAIGatewayPluginSwitchResponseBody {
	return s.Body
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) SetHeaders(v map[string]*string) *ModifyTairAIGatewayPluginSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) SetStatusCode(v int32) *ModifyTairAIGatewayPluginSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) SetBody(v *ModifyTairAIGatewayPluginSwitchResponseBody) *ModifyTairAIGatewayPluginSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
