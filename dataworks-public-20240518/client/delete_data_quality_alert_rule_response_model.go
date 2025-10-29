// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityAlertRuleResponseBody) *DeleteDataQualityAlertRuleResponse
	GetBody() *DeleteDataQualityAlertRuleResponseBody
}

type DeleteDataQualityAlertRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityAlertRuleResponse) GetBody() *DeleteDataQualityAlertRuleResponseBody {
	return s.Body
}

func (s *DeleteDataQualityAlertRuleResponse) SetHeaders(v map[string]*string) *DeleteDataQualityAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityAlertRuleResponse) SetStatusCode(v int32) *DeleteDataQualityAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityAlertRuleResponse) SetBody(v *DeleteDataQualityAlertRuleResponseBody) *DeleteDataQualityAlertRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
