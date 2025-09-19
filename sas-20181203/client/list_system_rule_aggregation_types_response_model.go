// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemRuleAggregationTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemRuleAggregationTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemRuleAggregationTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemRuleAggregationTypesResponseBody) *ListSystemRuleAggregationTypesResponse
	GetBody() *ListSystemRuleAggregationTypesResponseBody
}

type ListSystemRuleAggregationTypesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemRuleAggregationTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemRuleAggregationTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemRuleAggregationTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemRuleAggregationTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemRuleAggregationTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemRuleAggregationTypesResponse) GetBody() *ListSystemRuleAggregationTypesResponseBody {
	return s.Body
}

func (s *ListSystemRuleAggregationTypesResponse) SetHeaders(v map[string]*string) *ListSystemRuleAggregationTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemRuleAggregationTypesResponse) SetStatusCode(v int32) *ListSystemRuleAggregationTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemRuleAggregationTypesResponse) SetBody(v *ListSystemRuleAggregationTypesResponseBody) *ListSystemRuleAggregationTypesResponse {
	s.Body = v
	return s
}

func (s *ListSystemRuleAggregationTypesResponse) Validate() error {
	return dara.Validate(s)
}
