// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sIngressRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateK8sIngressRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateK8sIngressRuleResponseBody
	GetMessage() *string
}

type CreateK8sIngressRuleResponseBody struct {
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

func (s CreateK8sIngressRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sIngressRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateK8sIngressRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateK8sIngressRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateK8sIngressRuleResponseBody) SetCode(v int32) *CreateK8sIngressRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateK8sIngressRuleResponseBody) SetMessage(v string) *CreateK8sIngressRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateK8sIngressRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
