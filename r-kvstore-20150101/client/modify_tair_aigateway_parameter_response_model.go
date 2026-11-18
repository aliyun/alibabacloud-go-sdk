// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTairAIGatewayParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTairAIGatewayParameterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTairAIGatewayParameterResponseBody) *ModifyTairAIGatewayParameterResponse
	GetBody() *ModifyTairAIGatewayParameterResponseBody
}

type ModifyTairAIGatewayParameterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTairAIGatewayParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTairAIGatewayParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTairAIGatewayParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTairAIGatewayParameterResponse) GetBody() *ModifyTairAIGatewayParameterResponseBody {
	return s.Body
}

func (s *ModifyTairAIGatewayParameterResponse) SetHeaders(v map[string]*string) *ModifyTairAIGatewayParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyTairAIGatewayParameterResponse) SetStatusCode(v int32) *ModifyTairAIGatewayParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTairAIGatewayParameterResponse) SetBody(v *ModifyTairAIGatewayParameterResponseBody) *ModifyTairAIGatewayParameterResponse {
	s.Body = v
	return s
}

func (s *ModifyTairAIGatewayParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
