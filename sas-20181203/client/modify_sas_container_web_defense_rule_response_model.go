// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySasContainerWebDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySasContainerWebDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySasContainerWebDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifySasContainerWebDefenseRuleResponseBody) *ModifySasContainerWebDefenseRuleResponse
	GetBody() *ModifySasContainerWebDefenseRuleResponseBody
}

type ModifySasContainerWebDefenseRuleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySasContainerWebDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySasContainerWebDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySasContainerWebDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySasContainerWebDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySasContainerWebDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySasContainerWebDefenseRuleResponse) GetBody() *ModifySasContainerWebDefenseRuleResponseBody {
	return s.Body
}

func (s *ModifySasContainerWebDefenseRuleResponse) SetHeaders(v map[string]*string) *ModifySasContainerWebDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySasContainerWebDefenseRuleResponse) SetStatusCode(v int32) *ModifySasContainerWebDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleResponse) SetBody(v *ModifySasContainerWebDefenseRuleResponseBody) *ModifySasContainerWebDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *ModifySasContainerWebDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
