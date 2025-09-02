// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityResultsByRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityResultsByRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityResultsByRuleResponseBody) *ListQualityResultsByRuleResponse
	GetBody() *ListQualityResultsByRuleResponseBody
}

type ListQualityResultsByRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityResultsByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityResultsByRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleResponse) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityResultsByRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityResultsByRuleResponse) GetBody() *ListQualityResultsByRuleResponseBody {
	return s.Body
}

func (s *ListQualityResultsByRuleResponse) SetHeaders(v map[string]*string) *ListQualityResultsByRuleResponse {
	s.Headers = v
	return s
}

func (s *ListQualityResultsByRuleResponse) SetStatusCode(v int32) *ListQualityResultsByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityResultsByRuleResponse) SetBody(v *ListQualityResultsByRuleResponseBody) *ListQualityResultsByRuleResponse {
	s.Body = v
	return s
}

func (s *ListQualityResultsByRuleResponse) Validate() error {
	return dara.Validate(s)
}
