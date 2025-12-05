// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableRuleResponseBody) *DisableRuleResponse
	GetBody() *DisableRuleResponseBody
}

type DisableRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableRuleResponse) GetBody() *DisableRuleResponseBody {
	return s.Body
}

func (s *DisableRuleResponse) SetHeaders(v map[string]*string) *DisableRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableRuleResponse) SetStatusCode(v int32) *DisableRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRuleResponse) SetBody(v *DisableRuleResponseBody) *DisableRuleResponse {
	s.Body = v
	return s
}

func (s *DisableRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
