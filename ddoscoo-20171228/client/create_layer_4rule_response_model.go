// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer4RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLayer4RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLayer4RuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateLayer4RuleResponseBody) *CreateLayer4RuleResponse
	GetBody() *CreateLayer4RuleResponseBody
}

type CreateLayer4RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayer4RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLayer4RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLayer4RuleResponse) GetBody() *CreateLayer4RuleResponseBody {
	return s.Body
}

func (s *CreateLayer4RuleResponse) SetHeaders(v map[string]*string) *CreateLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer4RuleResponse) SetStatusCode(v int32) *CreateLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayer4RuleResponse) SetBody(v *CreateLayer4RuleResponseBody) *CreateLayer4RuleResponse {
	s.Body = v
	return s
}

func (s *CreateLayer4RuleResponse) Validate() error {
	return dara.Validate(s)
}
