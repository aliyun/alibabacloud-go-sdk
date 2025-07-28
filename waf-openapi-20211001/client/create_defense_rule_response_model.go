// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefenseRuleResponseBody) *CreateDefenseRuleResponse
	GetBody() *CreateDefenseRuleResponseBody
}

type CreateDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefenseRuleResponse) GetBody() *CreateDefenseRuleResponseBody {
	return s.Body
}

func (s *CreateDefenseRuleResponse) SetHeaders(v map[string]*string) *CreateDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseRuleResponse) SetStatusCode(v int32) *CreateDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseRuleResponse) SetBody(v *CreateDefenseRuleResponseBody) *CreateDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
