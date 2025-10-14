// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMetricRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMetricRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMetricRuleTemplateResponseBody) *ModifyMetricRuleTemplateResponse
	GetBody() *ModifyMetricRuleTemplateResponseBody
}

type ModifyMetricRuleTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMetricRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMetricRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMetricRuleTemplateResponse) GetBody() *ModifyMetricRuleTemplateResponseBody {
	return s.Body
}

func (s *ModifyMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *ModifyMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyMetricRuleTemplateResponse) SetStatusCode(v int32) *ModifyMetricRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponse) SetBody(v *ModifyMetricRuleTemplateResponseBody) *ModifyMetricRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyMetricRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
