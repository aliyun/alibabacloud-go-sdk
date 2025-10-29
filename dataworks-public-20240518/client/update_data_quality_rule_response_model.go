// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityRuleResponseBody) *UpdateDataQualityRuleResponse
	GetBody() *UpdateDataQualityRuleResponseBody
}

type UpdateDataQualityRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityRuleResponse) GetBody() *UpdateDataQualityRuleResponseBody {
	return s.Body
}

func (s *UpdateDataQualityRuleResponse) SetHeaders(v map[string]*string) *UpdateDataQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityRuleResponse) SetStatusCode(v int32) *UpdateDataQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityRuleResponse) SetBody(v *UpdateDataQualityRuleResponseBody) *UpdateDataQualityRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
