// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRecommendManagedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateRecommendManagedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateRecommendManagedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateRecommendManagedRulesResponseBody) *ListAggregateRecommendManagedRulesResponse
	GetBody() *ListAggregateRecommendManagedRulesResponseBody
}

type ListAggregateRecommendManagedRulesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateRecommendManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateRecommendManagedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRecommendManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateRecommendManagedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateRecommendManagedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateRecommendManagedRulesResponse) GetBody() *ListAggregateRecommendManagedRulesResponseBody {
	return s.Body
}

func (s *ListAggregateRecommendManagedRulesResponse) SetHeaders(v map[string]*string) *ListAggregateRecommendManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponse) SetStatusCode(v int32) *ListAggregateRecommendManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponse) SetBody(v *ListAggregateRecommendManagedRulesResponseBody) *ListAggregateRecommendManagedRulesResponse {
	s.Body = v
	return s
}

func (s *ListAggregateRecommendManagedRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
