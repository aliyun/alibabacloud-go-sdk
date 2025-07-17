// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertRuleResponseBody) *DeleteAlertRuleResponse
	GetBody() *DeleteAlertRuleResponseBody
}

type DeleteAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertRuleResponse) GetBody() *DeleteAlertRuleResponseBody {
	return s.Body
}

func (s *DeleteAlertRuleResponse) SetHeaders(v map[string]*string) *DeleteAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRuleResponse) SetStatusCode(v int32) *DeleteAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRuleResponse) SetBody(v *DeleteAlertRuleResponseBody) *DeleteAlertRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
