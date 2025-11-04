// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateScalingRuleResponseBody) *CreateScalingRuleResponse
	GetBody() *CreateScalingRuleResponseBody
}

type CreateScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScalingRuleResponse) GetBody() *CreateScalingRuleResponseBody {
	return s.Body
}

func (s *CreateScalingRuleResponse) SetHeaders(v map[string]*string) *CreateScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingRuleResponse) SetStatusCode(v int32) *CreateScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingRuleResponse) SetBody(v *CreateScalingRuleResponseBody) *CreateScalingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateScalingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
