// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutGroupMetricRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutGroupMetricRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutGroupMetricRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutGroupMetricRuleResponseBody) *PutGroupMetricRuleResponse
	GetBody() *PutGroupMetricRuleResponseBody
}

type PutGroupMetricRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutGroupMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutGroupMetricRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutGroupMetricRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutGroupMetricRuleResponse) GetBody() *PutGroupMetricRuleResponseBody {
	return s.Body
}

func (s *PutGroupMetricRuleResponse) SetHeaders(v map[string]*string) *PutGroupMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutGroupMetricRuleResponse) SetStatusCode(v int32) *PutGroupMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutGroupMetricRuleResponse) SetBody(v *PutGroupMetricRuleResponseBody) *PutGroupMetricRuleResponse {
	s.Body = v
	return s
}

func (s *PutGroupMetricRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
