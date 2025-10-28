// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sIngressRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sIngressRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sIngressRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sIngressRuleResponseBody) *DeleteK8sIngressRuleResponse
	GetBody() *DeleteK8sIngressRuleResponseBody
}

type DeleteK8sIngressRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sIngressRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sIngressRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sIngressRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sIngressRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sIngressRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sIngressRuleResponse) GetBody() *DeleteK8sIngressRuleResponseBody {
	return s.Body
}

func (s *DeleteK8sIngressRuleResponse) SetHeaders(v map[string]*string) *DeleteK8sIngressRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sIngressRuleResponse) SetStatusCode(v int32) *DeleteK8sIngressRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sIngressRuleResponse) SetBody(v *DeleteK8sIngressRuleResponseBody) *DeleteK8sIngressRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sIngressRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
