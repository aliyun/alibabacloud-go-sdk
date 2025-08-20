// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoBuildRuleResponseBody) *CreateRepoBuildRuleResponse
	GetBody() *CreateRepoBuildRuleResponseBody
}

type CreateRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoBuildRuleResponse) GetBody() *CreateRepoBuildRuleResponseBody {
	return s.Body
}

func (s *CreateRepoBuildRuleResponse) SetHeaders(v map[string]*string) *CreateRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoBuildRuleResponse) SetStatusCode(v int32) *CreateRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoBuildRuleResponse) SetBody(v *CreateRepoBuildRuleResponseBody) *CreateRepoBuildRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRepoBuildRuleResponse) Validate() error {
	return dara.Validate(s)
}
