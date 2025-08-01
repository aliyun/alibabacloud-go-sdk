// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateIsolationRuleResponseBody) *CreateIsolationRuleResponse
	GetBody() *CreateIsolationRuleResponseBody
}

type CreateIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIsolationRuleResponse) GetBody() *CreateIsolationRuleResponseBody {
	return s.Body
}

func (s *CreateIsolationRuleResponse) SetHeaders(v map[string]*string) *CreateIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateIsolationRuleResponse) SetStatusCode(v int32) *CreateIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIsolationRuleResponse) SetBody(v *CreateIsolationRuleResponseBody) *CreateIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateIsolationRuleResponse) Validate() error {
	return dara.Validate(s)
}
