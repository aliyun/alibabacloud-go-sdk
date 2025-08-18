// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOriginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOriginRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateOriginRuleResponseBody) *CreateOriginRuleResponse
	GetBody() *CreateOriginRuleResponseBody
}

type CreateOriginRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOriginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOriginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateOriginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOriginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOriginRuleResponse) GetBody() *CreateOriginRuleResponseBody {
	return s.Body
}

func (s *CreateOriginRuleResponse) SetHeaders(v map[string]*string) *CreateOriginRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateOriginRuleResponse) SetStatusCode(v int32) *CreateOriginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOriginRuleResponse) SetBody(v *CreateOriginRuleResponseBody) *CreateOriginRuleResponse {
	s.Body = v
	return s
}

func (s *CreateOriginRuleResponse) Validate() error {
	return dara.Validate(s)
}
