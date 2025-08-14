// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficMatchRuleToTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse
	GetBody() *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody
}

type AddTrafficMatchRuleToTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) GetBody() *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) SetStatusCode(v int32) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) SetBody(v *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) *AddTrafficMatchRuleToTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
