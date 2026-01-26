// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDispatchRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDispatchRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDispatchRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDispatchRuleResponseBody) *CreateDispatchRuleResponse
	GetBody() *CreateDispatchRuleResponseBody
}

type CreateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDispatchRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDispatchRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDispatchRuleResponse) GetBody() *CreateDispatchRuleResponseBody {
	return s.Body
}

func (s *CreateDispatchRuleResponse) SetHeaders(v map[string]*string) *CreateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDispatchRuleResponse) SetStatusCode(v int32) *CreateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDispatchRuleResponse) SetBody(v *CreateDispatchRuleResponseBody) *CreateDispatchRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDispatchRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
