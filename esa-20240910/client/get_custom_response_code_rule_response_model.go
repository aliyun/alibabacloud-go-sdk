// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomResponseCodeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomResponseCodeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomResponseCodeRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomResponseCodeRuleResponseBody) *GetCustomResponseCodeRuleResponse
	GetBody() *GetCustomResponseCodeRuleResponseBody
}

type GetCustomResponseCodeRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomResponseCodeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomResponseCodeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomResponseCodeRuleResponse) GoString() string {
	return s.String()
}

func (s *GetCustomResponseCodeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomResponseCodeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomResponseCodeRuleResponse) GetBody() *GetCustomResponseCodeRuleResponseBody {
	return s.Body
}

func (s *GetCustomResponseCodeRuleResponse) SetHeaders(v map[string]*string) *GetCustomResponseCodeRuleResponse {
	s.Headers = v
	return s
}

func (s *GetCustomResponseCodeRuleResponse) SetStatusCode(v int32) *GetCustomResponseCodeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomResponseCodeRuleResponse) SetBody(v *GetCustomResponseCodeRuleResponseBody) *GetCustomResponseCodeRuleResponse {
	s.Body = v
	return s
}

func (s *GetCustomResponseCodeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
