// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityRuleTemplateResponseBody) *CreateDataQualityRuleTemplateResponse
	GetBody() *CreateDataQualityRuleTemplateResponseBody
}

type CreateDataQualityRuleTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityRuleTemplateResponse) GetBody() *CreateDataQualityRuleTemplateResponseBody {
	return s.Body
}

func (s *CreateDataQualityRuleTemplateResponse) SetHeaders(v map[string]*string) *CreateDataQualityRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityRuleTemplateResponse) SetStatusCode(v int32) *CreateDataQualityRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityRuleTemplateResponse) SetBody(v *CreateDataQualityRuleTemplateResponseBody) *CreateDataQualityRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
