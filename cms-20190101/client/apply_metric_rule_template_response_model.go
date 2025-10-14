// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyMetricRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyMetricRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyMetricRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ApplyMetricRuleTemplateResponseBody) *ApplyMetricRuleTemplateResponse
	GetBody() *ApplyMetricRuleTemplateResponseBody
}

type ApplyMetricRuleTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyMetricRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyMetricRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyMetricRuleTemplateResponse) GetBody() *ApplyMetricRuleTemplateResponseBody {
	return s.Body
}

func (s *ApplyMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *ApplyMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApplyMetricRuleTemplateResponse) SetStatusCode(v int32) *ApplyMetricRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponse) SetBody(v *ApplyMetricRuleTemplateResponseBody) *ApplyMetricRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *ApplyMetricRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
