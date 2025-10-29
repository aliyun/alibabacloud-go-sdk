// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityRuleTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityRuleTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityRuleTemplateResponseBody) *DeleteDataQualityRuleTemplateResponse
	GetBody() *DeleteDataQualityRuleTemplateResponseBody
}

type DeleteDataQualityRuleTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityRuleTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityRuleTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityRuleTemplateResponse) GetBody() *DeleteDataQualityRuleTemplateResponseBody {
	return s.Body
}

func (s *DeleteDataQualityRuleTemplateResponse) SetHeaders(v map[string]*string) *DeleteDataQualityRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityRuleTemplateResponse) SetStatusCode(v int32) *DeleteDataQualityRuleTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityRuleTemplateResponse) SetBody(v *DeleteDataQualityRuleTemplateResponseBody) *DeleteDataQualityRuleTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityRuleTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
