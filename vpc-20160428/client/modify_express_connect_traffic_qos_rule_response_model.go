// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectTrafficQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectTrafficQosRuleResponseBody) *ModifyExpressConnectTrafficQosRuleResponse
	GetBody() *ModifyExpressConnectTrafficQosRuleResponseBody
}

type ModifyExpressConnectTrafficQosRuleResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectTrafficQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) GetBody() *ModifyExpressConnectTrafficQosRuleResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) SetStatusCode(v int32) *ModifyExpressConnectTrafficQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) SetBody(v *ModifyExpressConnectTrafficQosRuleResponseBody) *ModifyExpressConnectTrafficQosRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
