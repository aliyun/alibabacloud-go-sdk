// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sIngressRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sIngressRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sIngressRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sIngressRuleResponseBody) *UpdateK8sIngressRuleResponse
	GetBody() *UpdateK8sIngressRuleResponseBody
}

type UpdateK8sIngressRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sIngressRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sIngressRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sIngressRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sIngressRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sIngressRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sIngressRuleResponse) GetBody() *UpdateK8sIngressRuleResponseBody {
	return s.Body
}

func (s *UpdateK8sIngressRuleResponse) SetHeaders(v map[string]*string) *UpdateK8sIngressRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sIngressRuleResponse) SetStatusCode(v int32) *UpdateK8sIngressRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sIngressRuleResponse) SetBody(v *UpdateK8sIngressRuleResponseBody) *UpdateK8sIngressRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sIngressRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
