// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindings(v []*BatchUnbindTemplateResponseBodyBindings) *BatchUnbindTemplateResponseBody
	GetBindings() []*BatchUnbindTemplateResponseBodyBindings
	SetRequestId(v string) *BatchUnbindTemplateResponseBody
	GetRequestId() *string
}

type BatchUnbindTemplateResponseBody struct {
	Bindings []*BatchUnbindTemplateResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUnbindTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponseBody) GetBindings() []*BatchUnbindTemplateResponseBodyBindings {
	return s.Bindings
}

func (s *BatchUnbindTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUnbindTemplateResponseBody) SetBindings(v []*BatchUnbindTemplateResponseBodyBindings) *BatchUnbindTemplateResponseBody {
	s.Bindings = v
	return s
}

func (s *BatchUnbindTemplateResponseBody) SetRequestId(v string) *BatchUnbindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUnbindTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchUnbindTemplateResponseBodyBindings struct {
	// example:
	//
	// some error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 323*****994-cn-qingdao
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// group
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s BatchUnbindTemplateResponseBodyBindings) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplateResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateResponseBodyBindings) GetError() *string {
	return s.Error
}

func (s *BatchUnbindTemplateResponseBodyBindings) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchUnbindTemplateResponseBodyBindings) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchUnbindTemplateResponseBodyBindings) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetError(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.Error = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetInstanceId(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetInstanceType(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) SetTemplateId(v string) *BatchUnbindTemplateResponseBodyBindings {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplateResponseBodyBindings) Validate() error {
	return dara.Validate(s)
}
