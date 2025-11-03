// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectTrafficQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectTrafficQosRuleResponseBody) *DeleteExpressConnectTrafficQosRuleResponse
	GetBody() *DeleteExpressConnectTrafficQosRuleResponseBody
}

type DeleteExpressConnectTrafficQosRuleResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectTrafficQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectTrafficQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) GetBody() *DeleteExpressConnectTrafficQosRuleResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) SetStatusCode(v int32) *DeleteExpressConnectTrafficQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) SetBody(v *DeleteExpressConnectTrafficQosRuleResponseBody) *DeleteExpressConnectTrafficQosRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
