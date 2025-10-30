// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCopyRuleVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCopyRuleVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCopyRuleVariableResponse
	GetStatusCode() *int32
	SetBody(v *CheckCopyRuleVariableResponseBody) *CheckCopyRuleVariableResponse
	GetBody() *CheckCopyRuleVariableResponseBody
}

type CheckCopyRuleVariableResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCopyRuleVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCopyRuleVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCopyRuleVariableResponse) GoString() string {
	return s.String()
}

func (s *CheckCopyRuleVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCopyRuleVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCopyRuleVariableResponse) GetBody() *CheckCopyRuleVariableResponseBody {
	return s.Body
}

func (s *CheckCopyRuleVariableResponse) SetHeaders(v map[string]*string) *CheckCopyRuleVariableResponse {
	s.Headers = v
	return s
}

func (s *CheckCopyRuleVariableResponse) SetStatusCode(v int32) *CheckCopyRuleVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCopyRuleVariableResponse) SetBody(v *CheckCopyRuleVariableResponseBody) *CheckCopyRuleVariableResponse {
	s.Body = v
	return s
}

func (s *CheckCopyRuleVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
