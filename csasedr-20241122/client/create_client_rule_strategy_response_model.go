// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientRuleStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientRuleStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientRuleStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientRuleStrategyResponseBody) *CreateClientRuleStrategyResponse
	GetBody() *CreateClientRuleStrategyResponseBody
}

type CreateClientRuleStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientRuleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientRuleStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientRuleStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateClientRuleStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientRuleStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientRuleStrategyResponse) GetBody() *CreateClientRuleStrategyResponseBody {
	return s.Body
}

func (s *CreateClientRuleStrategyResponse) SetHeaders(v map[string]*string) *CreateClientRuleStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateClientRuleStrategyResponse) SetStatusCode(v int32) *CreateClientRuleStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientRuleStrategyResponse) SetBody(v *CreateClientRuleStrategyResponseBody) *CreateClientRuleStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateClientRuleStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
