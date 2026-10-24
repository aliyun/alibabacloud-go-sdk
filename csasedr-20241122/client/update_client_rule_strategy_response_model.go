// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientRuleStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientRuleStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientRuleStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientRuleStrategyResponseBody) *UpdateClientRuleStrategyResponse
	GetBody() *UpdateClientRuleStrategyResponseBody
}

type UpdateClientRuleStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientRuleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientRuleStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientRuleStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientRuleStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientRuleStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientRuleStrategyResponse) GetBody() *UpdateClientRuleStrategyResponseBody {
	return s.Body
}

func (s *UpdateClientRuleStrategyResponse) SetHeaders(v map[string]*string) *UpdateClientRuleStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientRuleStrategyResponse) SetStatusCode(v int32) *UpdateClientRuleStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientRuleStrategyResponse) SetBody(v *UpdateClientRuleStrategyResponseBody) *UpdateClientRuleStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateClientRuleStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
