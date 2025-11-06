// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLosslessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLosslessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLosslessRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLosslessRuleResponseBody) *ModifyLosslessRuleResponse
	GetBody() *ModifyLosslessRuleResponseBody
}

type ModifyLosslessRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLosslessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLosslessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLosslessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyLosslessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLosslessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLosslessRuleResponse) GetBody() *ModifyLosslessRuleResponseBody {
	return s.Body
}

func (s *ModifyLosslessRuleResponse) SetHeaders(v map[string]*string) *ModifyLosslessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyLosslessRuleResponse) SetStatusCode(v int32) *ModifyLosslessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLosslessRuleResponse) SetBody(v *ModifyLosslessRuleResponseBody) *ModifyLosslessRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyLosslessRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
