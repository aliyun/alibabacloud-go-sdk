// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVccGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVccGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateVccGrantRuleResponseBody) *CreateVccGrantRuleResponse
	GetBody() *CreateVccGrantRuleResponseBody
}

type CreateVccGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVccGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVccGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVccGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVccGrantRuleResponse) GetBody() *CreateVccGrantRuleResponseBody {
	return s.Body
}

func (s *CreateVccGrantRuleResponse) SetHeaders(v map[string]*string) *CreateVccGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateVccGrantRuleResponse) SetStatusCode(v int32) *CreateVccGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccGrantRuleResponse) SetBody(v *CreateVccGrantRuleResponseBody) *CreateVccGrantRuleResponse {
	s.Body = v
	return s
}

func (s *CreateVccGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
