// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddContainerDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddContainerDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddContainerDefenseRuleResponseBody) *AddContainerDefenseRuleResponse
	GetBody() *AddContainerDefenseRuleResponseBody
}

type AddContainerDefenseRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddContainerDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddContainerDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddContainerDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddContainerDefenseRuleResponse) GetBody() *AddContainerDefenseRuleResponseBody {
	return s.Body
}

func (s *AddContainerDefenseRuleResponse) SetHeaders(v map[string]*string) *AddContainerDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *AddContainerDefenseRuleResponse) SetStatusCode(v int32) *AddContainerDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddContainerDefenseRuleResponse) SetBody(v *AddContainerDefenseRuleResponseBody) *AddContainerDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *AddContainerDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
