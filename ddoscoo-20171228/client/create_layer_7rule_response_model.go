// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer7RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLayer7RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLayer7RuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateLayer7RuleResponseBody) *CreateLayer7RuleResponse
	GetBody() *CreateLayer7RuleResponseBody
}

type CreateLayer7RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayer7RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLayer7RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLayer7RuleResponse) GetBody() *CreateLayer7RuleResponseBody {
	return s.Body
}

func (s *CreateLayer7RuleResponse) SetHeaders(v map[string]*string) *CreateLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer7RuleResponse) SetStatusCode(v int32) *CreateLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayer7RuleResponse) SetBody(v *CreateLayer7RuleResponseBody) *CreateLayer7RuleResponse {
	s.Body = v
	return s
}

func (s *CreateLayer7RuleResponse) Validate() error {
	return dara.Validate(s)
}
