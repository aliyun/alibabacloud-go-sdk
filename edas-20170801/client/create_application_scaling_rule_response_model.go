// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationScalingRuleResponseBody) *CreateApplicationScalingRuleResponse
	GetBody() *CreateApplicationScalingRuleResponseBody
}

type CreateApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationScalingRuleResponse) GetBody() *CreateApplicationScalingRuleResponseBody {
	return s.Body
}

func (s *CreateApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *CreateApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationScalingRuleResponse) SetStatusCode(v int32) *CreateApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationScalingRuleResponse) SetBody(v *CreateApplicationScalingRuleResponseBody) *CreateApplicationScalingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationScalingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
