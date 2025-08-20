// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApsOptimizationStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApsOptimizationStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListApsOptimizationStrategyResponseBody) *ListApsOptimizationStrategyResponse
	GetBody() *ListApsOptimizationStrategyResponseBody
}

type ListApsOptimizationStrategyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApsOptimizationStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApsOptimizationStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApsOptimizationStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApsOptimizationStrategyResponse) GetBody() *ListApsOptimizationStrategyResponseBody {
	return s.Body
}

func (s *ListApsOptimizationStrategyResponse) SetHeaders(v map[string]*string) *ListApsOptimizationStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListApsOptimizationStrategyResponse) SetStatusCode(v int32) *ListApsOptimizationStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApsOptimizationStrategyResponse) SetBody(v *ListApsOptimizationStrategyResponseBody) *ListApsOptimizationStrategyResponse {
	s.Body = v
	return s
}

func (s *ListApsOptimizationStrategyResponse) Validate() error {
	return dara.Validate(s)
}
