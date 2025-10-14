// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActiveMetricRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableActiveMetricRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableActiveMetricRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableActiveMetricRuleResponseBody) *DisableActiveMetricRuleResponse
	GetBody() *DisableActiveMetricRuleResponseBody
}

type DisableActiveMetricRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableActiveMetricRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableActiveMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableActiveMetricRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableActiveMetricRuleResponse) GetBody() *DisableActiveMetricRuleResponseBody {
	return s.Body
}

func (s *DisableActiveMetricRuleResponse) SetHeaders(v map[string]*string) *DisableActiveMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableActiveMetricRuleResponse) SetStatusCode(v int32) *DisableActiveMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableActiveMetricRuleResponse) SetBody(v *DisableActiveMetricRuleResponseBody) *DisableActiveMetricRuleResponse {
	s.Body = v
	return s
}

func (s *DisableActiveMetricRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
