// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityScoreRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecurityScoreRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecurityScoreRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetSecurityScoreRuleResponseBody) *GetSecurityScoreRuleResponse
	GetBody() *GetSecurityScoreRuleResponseBody
}

type GetSecurityScoreRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityScoreRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityScoreRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityScoreRuleResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityScoreRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecurityScoreRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecurityScoreRuleResponse) GetBody() *GetSecurityScoreRuleResponseBody {
	return s.Body
}

func (s *GetSecurityScoreRuleResponse) SetHeaders(v map[string]*string) *GetSecurityScoreRuleResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityScoreRuleResponse) SetStatusCode(v int32) *GetSecurityScoreRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityScoreRuleResponse) SetBody(v *GetSecurityScoreRuleResponseBody) *GetSecurityScoreRuleResponse {
	s.Body = v
	return s
}

func (s *GetSecurityScoreRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
