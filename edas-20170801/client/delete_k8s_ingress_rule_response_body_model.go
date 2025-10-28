// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sIngressRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteK8sIngressRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteK8sIngressRuleResponseBody
	GetMessage() *string
}

type DeleteK8sIngressRuleResponseBody struct {
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

func (s DeleteK8sIngressRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sIngressRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sIngressRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteK8sIngressRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteK8sIngressRuleResponseBody) SetCode(v int32) *DeleteK8sIngressRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteK8sIngressRuleResponseBody) SetMessage(v string) *DeleteK8sIngressRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteK8sIngressRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
