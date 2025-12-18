// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveAggregateConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactiveAggregateConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactiveAggregateConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeactiveAggregateConfigRulesResponseBody) *DeactiveAggregateConfigRulesResponse
	GetBody() *DeactiveAggregateConfigRulesResponseBody
}

type DeactiveAggregateConfigRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactiveAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactiveAggregateConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactiveAggregateConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactiveAggregateConfigRulesResponse) GetBody() *DeactiveAggregateConfigRulesResponseBody {
	return s.Body
}

func (s *DeactiveAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *DeactiveAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponse) SetStatusCode(v int32) *DeactiveAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponse) SetBody(v *DeactiveAggregateConfigRulesResponseBody) *DeactiveAggregateConfigRulesResponse {
	s.Body = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
