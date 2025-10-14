// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetricRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetricRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetricRuleTemplateResponseBody) *DeleteMetricRuleTemplateResponse
	GetBody() *DeleteMetricRuleTemplateResponseBody
}

type DeleteMetricRuleTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetricRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetricRuleTemplateResponse) GetBody() *DeleteMetricRuleTemplateResponseBody {
	return s.Body
}

func (s *DeleteMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleTemplateResponse) SetStatusCode(v int32) *DeleteMetricRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponse) SetBody(v *DeleteMetricRuleTemplateResponseBody) *DeleteMetricRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteMetricRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
