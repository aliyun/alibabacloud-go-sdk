// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityRuleTemplateResponseBody) *UpdateDataQualityRuleTemplateResponse
	GetBody() *UpdateDataQualityRuleTemplateResponseBody
}

type UpdateDataQualityRuleTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityRuleTemplateResponse) GetBody() *UpdateDataQualityRuleTemplateResponseBody {
	return s.Body
}

func (s *UpdateDataQualityRuleTemplateResponse) SetHeaders(v map[string]*string) *UpdateDataQualityRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityRuleTemplateResponse) SetStatusCode(v int32) *UpdateDataQualityRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateResponse) SetBody(v *UpdateDataQualityRuleTemplateResponseBody) *UpdateDataQualityRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
