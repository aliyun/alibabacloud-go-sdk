// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressConnectTrafficQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressConnectTrafficQosRuleResponseBody) *CreateExpressConnectTrafficQosRuleResponse
	GetBody() *CreateExpressConnectTrafficQosRuleResponseBody
}

type CreateExpressConnectTrafficQosRuleResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectTrafficQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectTrafficQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressConnectTrafficQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressConnectTrafficQosRuleResponse) GetBody() *CreateExpressConnectTrafficQosRuleResponseBody {
	return s.Body
}

func (s *CreateExpressConnectTrafficQosRuleResponse) SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponse) SetStatusCode(v int32) *CreateExpressConnectTrafficQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponse) SetBody(v *CreateExpressConnectTrafficQosRuleResponseBody) *CreateExpressConnectTrafficQosRuleResponse {
	s.Body = v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponse) Validate() error {
	return dara.Validate(s)
}
