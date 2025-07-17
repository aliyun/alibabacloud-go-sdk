// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateAlertRuleResponseBody) *CreateOrUpdateAlertRuleResponse
	GetBody() *CreateOrUpdateAlertRuleResponseBody
}

type CreateOrUpdateAlertRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateAlertRuleResponse) GetBody() *CreateOrUpdateAlertRuleResponseBody {
	return s.Body
}

func (s *CreateOrUpdateAlertRuleResponse) SetHeaders(v map[string]*string) *CreateOrUpdateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponse) SetStatusCode(v int32) *CreateOrUpdateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponse) SetBody(v *CreateOrUpdateAlertRuleResponseBody) *CreateOrUpdateAlertRuleResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
