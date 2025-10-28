// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sIngressRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sIngressRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sIngressRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sIngressRulesResponseBody) *ListK8sIngressRulesResponse
	GetBody() *ListK8sIngressRulesResponseBody
}

type ListK8sIngressRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sIngressRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sIngressRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponse) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sIngressRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sIngressRulesResponse) GetBody() *ListK8sIngressRulesResponseBody {
	return s.Body
}

func (s *ListK8sIngressRulesResponse) SetHeaders(v map[string]*string) *ListK8sIngressRulesResponse {
	s.Headers = v
	return s
}

func (s *ListK8sIngressRulesResponse) SetStatusCode(v int32) *ListK8sIngressRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sIngressRulesResponse) SetBody(v *ListK8sIngressRulesResponseBody) *ListK8sIngressRulesResponse {
	s.Body = v
	return s
}

func (s *ListK8sIngressRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
