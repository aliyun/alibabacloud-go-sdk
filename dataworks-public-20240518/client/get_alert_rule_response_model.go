// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertRuleResponseBody) *GetAlertRuleResponse
	GetBody() *GetAlertRuleResponseBody
}

type GetAlertRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertRuleResponse) GetBody() *GetAlertRuleResponseBody {
	return s.Body
}

func (s *GetAlertRuleResponse) SetHeaders(v map[string]*string) *GetAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAlertRuleResponse) SetStatusCode(v int32) *GetAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertRuleResponse) SetBody(v *GetAlertRuleResponseBody) *GetAlertRuleResponse {
	s.Body = v
	return s
}

func (s *GetAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
