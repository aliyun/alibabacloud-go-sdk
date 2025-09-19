// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSecurityScoreRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeSecurityScoreRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeSecurityScoreRuleResponse
	GetStatusCode() *int32
	SetBody(v *ChangeSecurityScoreRuleResponseBody) *ChangeSecurityScoreRuleResponse
	GetBody() *ChangeSecurityScoreRuleResponseBody
}

type ChangeSecurityScoreRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeSecurityScoreRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeSecurityScoreRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeSecurityScoreRuleResponse) GoString() string {
	return s.String()
}

func (s *ChangeSecurityScoreRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeSecurityScoreRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeSecurityScoreRuleResponse) GetBody() *ChangeSecurityScoreRuleResponseBody {
	return s.Body
}

func (s *ChangeSecurityScoreRuleResponse) SetHeaders(v map[string]*string) *ChangeSecurityScoreRuleResponse {
	s.Headers = v
	return s
}

func (s *ChangeSecurityScoreRuleResponse) SetStatusCode(v int32) *ChangeSecurityScoreRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeSecurityScoreRuleResponse) SetBody(v *ChangeSecurityScoreRuleResponseBody) *ChangeSecurityScoreRuleResponse {
	s.Body = v
	return s
}

func (s *ChangeSecurityScoreRuleResponse) Validate() error {
	return dara.Validate(s)
}
