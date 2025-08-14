// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficMatchRuleToTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetBody() *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody
}

type ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) GetBody() *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) SetStatusCode(v int32) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) SetBody(v *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
