// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertRuleResponseBody) *UpdateAlertRuleResponse
	GetBody() *UpdateAlertRuleResponseBody
}

type UpdateAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertRuleResponse) GetBody() *UpdateAlertRuleResponseBody {
	return s.Body
}

func (s *UpdateAlertRuleResponse) SetHeaders(v map[string]*string) *UpdateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertRuleResponse) SetStatusCode(v int32) *UpdateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertRuleResponse) SetBody(v *UpdateAlertRuleResponseBody) *UpdateAlertRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
