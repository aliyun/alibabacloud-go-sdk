// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWafRuleResponseBody) *CreateWafRuleResponse
	GetBody() *CreateWafRuleResponseBody
}

type CreateWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWafRuleResponse) GetBody() *CreateWafRuleResponseBody {
	return s.Body
}

func (s *CreateWafRuleResponse) SetHeaders(v map[string]*string) *CreateWafRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWafRuleResponse) SetStatusCode(v int32) *CreateWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWafRuleResponse) SetBody(v *CreateWafRuleResponseBody) *CreateWafRuleResponse {
	s.Body = v
	return s
}

func (s *CreateWafRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
