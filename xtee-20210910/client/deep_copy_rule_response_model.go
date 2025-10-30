// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepCopyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeepCopyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeepCopyRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeepCopyRuleResponseBody) *DeepCopyRuleResponse
	GetBody() *DeepCopyRuleResponseBody
}

type DeepCopyRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeepCopyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeepCopyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeepCopyRuleResponse) GoString() string {
	return s.String()
}

func (s *DeepCopyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeepCopyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeepCopyRuleResponse) GetBody() *DeepCopyRuleResponseBody {
	return s.Body
}

func (s *DeepCopyRuleResponse) SetHeaders(v map[string]*string) *DeepCopyRuleResponse {
	s.Headers = v
	return s
}

func (s *DeepCopyRuleResponse) SetStatusCode(v int32) *DeepCopyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeepCopyRuleResponse) SetBody(v *DeepCopyRuleResponseBody) *DeepCopyRuleResponse {
	s.Body = v
	return s
}

func (s *DeepCopyRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
