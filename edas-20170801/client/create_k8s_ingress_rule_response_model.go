// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sIngressRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateK8sIngressRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateK8sIngressRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateK8sIngressRuleResponseBody) *CreateK8sIngressRuleResponse
	GetBody() *CreateK8sIngressRuleResponseBody
}

type CreateK8sIngressRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateK8sIngressRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateK8sIngressRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sIngressRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateK8sIngressRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateK8sIngressRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateK8sIngressRuleResponse) GetBody() *CreateK8sIngressRuleResponseBody {
	return s.Body
}

func (s *CreateK8sIngressRuleResponse) SetHeaders(v map[string]*string) *CreateK8sIngressRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateK8sIngressRuleResponse) SetStatusCode(v int32) *CreateK8sIngressRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateK8sIngressRuleResponse) SetBody(v *CreateK8sIngressRuleResponseBody) *CreateK8sIngressRuleResponse {
	s.Body = v
	return s
}

func (s *CreateK8sIngressRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
