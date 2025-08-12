// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCustomMetricRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCustomMetricRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutCustomMetricRuleResponseBody) *PutCustomMetricRuleResponse
	GetBody() *PutCustomMetricRuleResponseBody
}

type PutCustomMetricRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCustomMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCustomMetricRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCustomMetricRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCustomMetricRuleResponse) GetBody() *PutCustomMetricRuleResponseBody {
	return s.Body
}

func (s *PutCustomMetricRuleResponse) SetHeaders(v map[string]*string) *PutCustomMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutCustomMetricRuleResponse) SetStatusCode(v int32) *PutCustomMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCustomMetricRuleResponse) SetBody(v *PutCustomMetricRuleResponseBody) *PutCustomMetricRuleResponse {
	s.Body = v
	return s
}

func (s *PutCustomMetricRuleResponse) Validate() error {
	return dara.Validate(s)
}
