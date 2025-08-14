// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
	GetBody() *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody
}

type RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) GetBody() *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) SetStatusCode(v int32) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) SetBody(v *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
