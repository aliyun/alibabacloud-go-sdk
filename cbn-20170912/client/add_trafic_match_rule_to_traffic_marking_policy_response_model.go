// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTraficMatchRuleToTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTraficMatchRuleToTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTraficMatchRuleToTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) *AddTraficMatchRuleToTrafficMarkingPolicyResponse
	GetBody() *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody
}

type AddTraficMatchRuleToTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) GetBody() *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *AddTraficMatchRuleToTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) SetStatusCode(v int32) *AddTraficMatchRuleToTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) SetBody(v *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) *AddTraficMatchRuleToTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
