// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRuleTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityRuleTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityRuleTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityRuleTemplatesResponseBody) *ListDataQualityRuleTemplatesResponse
	GetBody() *ListDataQualityRuleTemplatesResponseBody
}

type ListDataQualityRuleTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityRuleTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityRuleTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityRuleTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityRuleTemplatesResponse) GetBody() *ListDataQualityRuleTemplatesResponseBody {
	return s.Body
}

func (s *ListDataQualityRuleTemplatesResponse) SetHeaders(v map[string]*string) *ListDataQualityRuleTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponse) SetStatusCode(v int32) *ListDataQualityRuleTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponse) SetBody(v *ListDataQualityRuleTemplatesResponseBody) *ListDataQualityRuleTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
