// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomResponseCodeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomResponseCodeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomResponseCodeRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomResponseCodeRuleResponseBody) *CreateCustomResponseCodeRuleResponse
	GetBody() *CreateCustomResponseCodeRuleResponseBody
}

type CreateCustomResponseCodeRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomResponseCodeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomResponseCodeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomResponseCodeRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomResponseCodeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomResponseCodeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomResponseCodeRuleResponse) GetBody() *CreateCustomResponseCodeRuleResponseBody {
	return s.Body
}

func (s *CreateCustomResponseCodeRuleResponse) SetHeaders(v map[string]*string) *CreateCustomResponseCodeRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomResponseCodeRuleResponse) SetStatusCode(v int32) *CreateCustomResponseCodeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomResponseCodeRuleResponse) SetBody(v *CreateCustomResponseCodeRuleResponseBody) *CreateCustomResponseCodeRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCustomResponseCodeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
