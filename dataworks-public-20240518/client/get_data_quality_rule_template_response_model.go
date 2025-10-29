// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityRuleTemplateResponseBody) *GetDataQualityRuleTemplateResponse
	GetBody() *GetDataQualityRuleTemplateResponseBody
}

type GetDataQualityRuleTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityRuleTemplateResponse) GetBody() *GetDataQualityRuleTemplateResponseBody {
	return s.Body
}

func (s *GetDataQualityRuleTemplateResponse) SetHeaders(v map[string]*string) *GetDataQualityRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityRuleTemplateResponse) SetStatusCode(v int32) *GetDataQualityRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponse) SetBody(v *GetDataQualityRuleTemplateResponseBody) *GetDataQualityRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
