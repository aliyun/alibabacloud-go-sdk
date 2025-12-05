// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkRuleResponseBody) *UpdateNetworkRuleResponse
	GetBody() *UpdateNetworkRuleResponseBody
}

type UpdateNetworkRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkRuleResponse) GetBody() *UpdateNetworkRuleResponseBody {
	return s.Body
}

func (s *UpdateNetworkRuleResponse) SetHeaders(v map[string]*string) *UpdateNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkRuleResponse) SetStatusCode(v int32) *UpdateNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkRuleResponse) SetBody(v *UpdateNetworkRuleResponseBody) *UpdateNetworkRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
