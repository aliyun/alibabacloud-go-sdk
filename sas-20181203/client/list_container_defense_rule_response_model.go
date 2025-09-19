// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContainerDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContainerDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContainerDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListContainerDefenseRuleResponseBody) *ListContainerDefenseRuleResponse
	GetBody() *ListContainerDefenseRuleResponseBody
}

type ListContainerDefenseRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContainerDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContainerDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContainerDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContainerDefenseRuleResponse) GetBody() *ListContainerDefenseRuleResponseBody {
	return s.Body
}

func (s *ListContainerDefenseRuleResponse) SetHeaders(v map[string]*string) *ListContainerDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ListContainerDefenseRuleResponse) SetStatusCode(v int32) *ListContainerDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContainerDefenseRuleResponse) SetBody(v *ListContainerDefenseRuleResponseBody) *ListContainerDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *ListContainerDefenseRuleResponse) Validate() error {
	return dara.Validate(s)
}
