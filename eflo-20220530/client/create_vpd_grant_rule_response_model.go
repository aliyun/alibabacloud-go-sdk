// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdGrantRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpdGrantRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpdGrantRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpdGrantRuleResponseBody) *CreateVpdGrantRuleResponse
	GetBody() *CreateVpdGrantRuleResponseBody
}

type CreateVpdGrantRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpdGrantRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpdGrantRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdGrantRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpdGrantRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpdGrantRuleResponse) GetBody() *CreateVpdGrantRuleResponseBody {
	return s.Body
}

func (s *CreateVpdGrantRuleResponse) SetHeaders(v map[string]*string) *CreateVpdGrantRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateVpdGrantRuleResponse) SetStatusCode(v int32) *CreateVpdGrantRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpdGrantRuleResponse) SetBody(v *CreateVpdGrantRuleResponseBody) *CreateVpdGrantRuleResponse {
	s.Body = v
	return s
}

func (s *CreateVpdGrantRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
