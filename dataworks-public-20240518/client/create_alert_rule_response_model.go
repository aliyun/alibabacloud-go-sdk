// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertRuleResponseBody) *CreateAlertRuleResponse
	GetBody() *CreateAlertRuleResponseBody
}

type CreateAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertRuleResponse) GetBody() *CreateAlertRuleResponseBody {
	return s.Body
}

func (s *CreateAlertRuleResponse) SetHeaders(v map[string]*string) *CreateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertRuleResponse) SetStatusCode(v int32) *CreateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertRuleResponse) SetBody(v *CreateAlertRuleResponseBody) *CreateAlertRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
