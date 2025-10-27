// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCheckRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCheckRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCheckRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCheckRuleResponseBody) *ModifyCheckRuleResponse
	GetBody() *ModifyCheckRuleResponseBody
}

type ModifyCheckRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCheckRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCheckRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCheckRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyCheckRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCheckRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCheckRuleResponse) GetBody() *ModifyCheckRuleResponseBody {
	return s.Body
}

func (s *ModifyCheckRuleResponse) SetHeaders(v map[string]*string) *ModifyCheckRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyCheckRuleResponse) SetStatusCode(v int32) *ModifyCheckRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCheckRuleResponse) SetBody(v *ModifyCheckRuleResponseBody) *ModifyCheckRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyCheckRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
