// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRuleResponseBody) *CreateRuleResponse
	GetBody() *CreateRuleResponseBody
}

type CreateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRuleResponse) GetBody() *CreateRuleResponseBody {
	return s.Body
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
