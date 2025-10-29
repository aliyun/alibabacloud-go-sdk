// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityAlertRuleResponseBody) *CreateDataQualityAlertRuleResponse
	GetBody() *CreateDataQualityAlertRuleResponseBody
}

type CreateDataQualityAlertRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityAlertRuleResponse) GetBody() *CreateDataQualityAlertRuleResponseBody {
	return s.Body
}

func (s *CreateDataQualityAlertRuleResponse) SetHeaders(v map[string]*string) *CreateDataQualityAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityAlertRuleResponse) SetStatusCode(v int32) *CreateDataQualityAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityAlertRuleResponse) SetBody(v *CreateDataQualityAlertRuleResponseBody) *CreateDataQualityAlertRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
