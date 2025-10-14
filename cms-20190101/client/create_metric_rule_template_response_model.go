// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMetricRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMetricRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateMetricRuleTemplateResponseBody) *CreateMetricRuleTemplateResponse
	GetBody() *CreateMetricRuleTemplateResponseBody
}

type CreateMetricRuleTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMetricRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMetricRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMetricRuleTemplateResponse) GetBody() *CreateMetricRuleTemplateResponseBody {
	return s.Body
}

func (s *CreateMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *CreateMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricRuleTemplateResponse) SetStatusCode(v int32) *CreateMetricRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMetricRuleTemplateResponse) SetBody(v *CreateMetricRuleTemplateResponseBody) *CreateMetricRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateMetricRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
