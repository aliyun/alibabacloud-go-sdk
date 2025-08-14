// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTraficMatchRuleFromTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
	GetBody() *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody
}

type RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) GetBody() *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) SetStatusCode(v int32) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) SetBody(v *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
