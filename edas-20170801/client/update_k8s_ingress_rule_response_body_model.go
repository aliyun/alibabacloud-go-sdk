// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sIngressRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateK8sIngressRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sIngressRuleResponseBody
	GetMessage() *string
}

type UpdateK8sIngressRuleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateK8sIngressRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sIngressRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sIngressRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sIngressRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sIngressRuleResponseBody) SetCode(v int32) *UpdateK8sIngressRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sIngressRuleResponseBody) SetMessage(v string) *UpdateK8sIngressRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sIngressRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
