// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRulesResponse
	GetStatusCode() *int32
	SetBody(v *CheckRulesResponseBody) *CheckRulesResponse
	GetBody() *CheckRulesResponseBody
}

type CheckRulesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRulesResponse) GoString() string {
	return s.String()
}

func (s *CheckRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRulesResponse) GetBody() *CheckRulesResponseBody {
	return s.Body
}

func (s *CheckRulesResponse) SetHeaders(v map[string]*string) *CheckRulesResponse {
	s.Headers = v
	return s
}

func (s *CheckRulesResponse) SetStatusCode(v int32) *CheckRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRulesResponse) SetBody(v *CheckRulesResponseBody) *CheckRulesResponse {
	s.Body = v
	return s
}

func (s *CheckRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
