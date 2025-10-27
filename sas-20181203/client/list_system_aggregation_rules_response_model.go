// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAggregationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemAggregationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemAggregationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemAggregationRulesResponseBody) *ListSystemAggregationRulesResponse
	GetBody() *ListSystemAggregationRulesResponseBody
}

type ListSystemAggregationRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemAggregationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemAggregationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAggregationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemAggregationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemAggregationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemAggregationRulesResponse) GetBody() *ListSystemAggregationRulesResponseBody {
	return s.Body
}

func (s *ListSystemAggregationRulesResponse) SetHeaders(v map[string]*string) *ListSystemAggregationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemAggregationRulesResponse) SetStatusCode(v int32) *ListSystemAggregationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemAggregationRulesResponse) SetBody(v *ListSystemAggregationRulesResponseBody) *ListSystemAggregationRulesResponse {
	s.Body = v
	return s
}

func (s *ListSystemAggregationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
