// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSasContainerWebDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSasContainerWebDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSasContainerWebDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListSasContainerWebDefenseRuleResponseBody) *ListSasContainerWebDefenseRuleResponse
	GetBody() *ListSasContainerWebDefenseRuleResponseBody
}

type ListSasContainerWebDefenseRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSasContainerWebDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSasContainerWebDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSasContainerWebDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSasContainerWebDefenseRuleResponse) GetBody() *ListSasContainerWebDefenseRuleResponseBody {
	return s.Body
}

func (s *ListSasContainerWebDefenseRuleResponse) SetHeaders(v map[string]*string) *ListSasContainerWebDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponse) SetStatusCode(v int32) *ListSasContainerWebDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponse) SetBody(v *ListSasContainerWebDefenseRuleResponseBody) *ListSasContainerWebDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *ListSasContainerWebDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
