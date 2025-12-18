// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateConfigRulesResponseBody) *ListAggregateConfigRulesResponse
	GetBody() *ListAggregateConfigRulesResponseBody
}

type ListAggregateConfigRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateConfigRulesResponse) GetBody() *ListAggregateConfigRulesResponseBody {
	return s.Body
}

func (s *ListAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *ListAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigRulesResponse) SetStatusCode(v int32) *ListAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigRulesResponse) SetBody(v *ListAggregateConfigRulesResponseBody) *ListAggregateConfigRulesResponse {
	s.Body = v
	return s
}

func (s *ListAggregateConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
