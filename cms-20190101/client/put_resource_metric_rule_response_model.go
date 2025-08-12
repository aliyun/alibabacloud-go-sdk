// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutResourceMetricRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutResourceMetricRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutResourceMetricRuleResponseBody) *PutResourceMetricRuleResponse
	GetBody() *PutResourceMetricRuleResponseBody
}

type PutResourceMetricRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutResourceMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutResourceMetricRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutResourceMetricRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutResourceMetricRuleResponse) GetBody() *PutResourceMetricRuleResponseBody {
	return s.Body
}

func (s *PutResourceMetricRuleResponse) SetHeaders(v map[string]*string) *PutResourceMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutResourceMetricRuleResponse) SetStatusCode(v int32) *PutResourceMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutResourceMetricRuleResponse) SetBody(v *PutResourceMetricRuleResponseBody) *PutResourceMetricRuleResponse {
	s.Body = v
	return s
}

func (s *PutResourceMetricRuleResponse) Validate() error {
	return dara.Validate(s)
}
