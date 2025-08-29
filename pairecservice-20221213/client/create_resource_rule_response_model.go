// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceRuleResponseBody) *CreateResourceRuleResponse
	GetBody() *CreateResourceRuleResponseBody
}

type CreateResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceRuleResponse) GetBody() *CreateResourceRuleResponseBody {
	return s.Body
}

func (s *CreateResourceRuleResponse) SetHeaders(v map[string]*string) *CreateResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceRuleResponse) SetStatusCode(v int32) *CreateResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceRuleResponse) SetBody(v *CreateResourceRuleResponseBody) *CreateResourceRuleResponse {
	s.Body = v
	return s
}

func (s *CreateResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
