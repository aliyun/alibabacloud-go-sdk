// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityAlertRuleResponseBody) *GetDataQualityAlertRuleResponse
	GetBody() *GetDataQualityAlertRuleResponseBody
}

type GetDataQualityAlertRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityAlertRuleResponse) GetBody() *GetDataQualityAlertRuleResponseBody {
	return s.Body
}

func (s *GetDataQualityAlertRuleResponse) SetHeaders(v map[string]*string) *GetDataQualityAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityAlertRuleResponse) SetStatusCode(v int32) *GetDataQualityAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityAlertRuleResponse) SetBody(v *GetDataQualityAlertRuleResponseBody) *GetDataQualityAlertRuleResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
