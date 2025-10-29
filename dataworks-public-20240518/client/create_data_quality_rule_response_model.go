// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityRuleResponseBody) *CreateDataQualityRuleResponse
	GetBody() *CreateDataQualityRuleResponseBody
}

type CreateDataQualityRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityRuleResponse) GetBody() *CreateDataQualityRuleResponseBody {
	return s.Body
}

func (s *CreateDataQualityRuleResponse) SetHeaders(v map[string]*string) *CreateDataQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityRuleResponse) SetStatusCode(v int32) *CreateDataQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityRuleResponse) SetBody(v *CreateDataQualityRuleResponseBody) *CreateDataQualityRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
