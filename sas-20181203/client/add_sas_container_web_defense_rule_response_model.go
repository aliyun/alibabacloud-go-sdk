// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasContainerWebDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSasContainerWebDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSasContainerWebDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddSasContainerWebDefenseRuleResponseBody) *AddSasContainerWebDefenseRuleResponse
	GetBody() *AddSasContainerWebDefenseRuleResponseBody
}

type AddSasContainerWebDefenseRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSasContainerWebDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSasContainerWebDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSasContainerWebDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *AddSasContainerWebDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSasContainerWebDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSasContainerWebDefenseRuleResponse) GetBody() *AddSasContainerWebDefenseRuleResponseBody {
	return s.Body
}

func (s *AddSasContainerWebDefenseRuleResponse) SetHeaders(v map[string]*string) *AddSasContainerWebDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *AddSasContainerWebDefenseRuleResponse) SetStatusCode(v int32) *AddSasContainerWebDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleResponse) SetBody(v *AddSasContainerWebDefenseRuleResponseBody) *AddSasContainerWebDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *AddSasContainerWebDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
