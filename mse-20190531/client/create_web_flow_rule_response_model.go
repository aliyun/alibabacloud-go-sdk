// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebFlowRuleResponseBody) *CreateWebFlowRuleResponse
	GetBody() *CreateWebFlowRuleResponseBody
}

type CreateWebFlowRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWebFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebFlowRuleResponse) GetBody() *CreateWebFlowRuleResponseBody {
	return s.Body
}

func (s *CreateWebFlowRuleResponse) SetHeaders(v map[string]*string) *CreateWebFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWebFlowRuleResponse) SetStatusCode(v int32) *CreateWebFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebFlowRuleResponse) SetBody(v *CreateWebFlowRuleResponseBody) *CreateWebFlowRuleResponse {
	s.Body = v
	return s
}

func (s *CreateWebFlowRuleResponse) Validate() error {
	return dara.Validate(s)
}
