// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAggregateConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveAggregateConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveAggregateConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *ActiveAggregateConfigRulesResponseBody) *ActiveAggregateConfigRulesResponse
	GetBody() *ActiveAggregateConfigRulesResponseBody
}

type ActiveAggregateConfigRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveAggregateConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveAggregateConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveAggregateConfigRulesResponse) GetBody() *ActiveAggregateConfigRulesResponseBody {
	return s.Body
}

func (s *ActiveAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *ActiveAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ActiveAggregateConfigRulesResponse) SetStatusCode(v int32) *ActiveAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponse) SetBody(v *ActiveAggregateConfigRulesResponseBody) *ActiveAggregateConfigRulesResponse {
	s.Body = v
	return s
}

func (s *ActiveAggregateConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
