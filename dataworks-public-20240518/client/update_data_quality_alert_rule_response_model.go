// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityAlertRuleResponseBody) *UpdateDataQualityAlertRuleResponse
	GetBody() *UpdateDataQualityAlertRuleResponseBody
}

type UpdateDataQualityAlertRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityAlertRuleResponse) GetBody() *UpdateDataQualityAlertRuleResponseBody {
	return s.Body
}

func (s *UpdateDataQualityAlertRuleResponse) SetHeaders(v map[string]*string) *UpdateDataQualityAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityAlertRuleResponse) SetStatusCode(v int32) *UpdateDataQualityAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityAlertRuleResponse) SetBody(v *UpdateDataQualityAlertRuleResponseBody) *UpdateDataQualityAlertRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
