// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindings(v []*BatchBindTemplateResponseBodyBindings) *BatchBindTemplateResponseBody
	GetBindings() []*BatchBindTemplateResponseBodyBindings
	SetRequestId(v string) *BatchBindTemplateResponseBody
	GetRequestId() *string
}

type BatchBindTemplateResponseBody struct {
	Bindings []*BatchBindTemplateResponseBodyBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchBindTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponseBody) GetBindings() []*BatchBindTemplateResponseBodyBindings {
	return s.Bindings
}

func (s *BatchBindTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchBindTemplateResponseBody) SetBindings(v []*BatchBindTemplateResponseBodyBindings) *BatchBindTemplateResponseBody {
	s.Bindings = v
	return s
}

func (s *BatchBindTemplateResponseBody) SetRequestId(v string) *BatchBindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchBindTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchBindTemplateResponseBodyBindings struct {
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

func (s BatchBindTemplateResponseBodyBindings) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplateResponseBodyBindings) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateResponseBodyBindings) GetError() *string {
	return s.Error
}

func (s *BatchBindTemplateResponseBodyBindings) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchBindTemplateResponseBodyBindings) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchBindTemplateResponseBodyBindings) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchBindTemplateResponseBodyBindings) SetError(v string) *BatchBindTemplateResponseBodyBindings {
	s.Error = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetInstanceId(v string) *BatchBindTemplateResponseBodyBindings {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetInstanceType(v string) *BatchBindTemplateResponseBodyBindings {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) SetTemplateId(v string) *BatchBindTemplateResponseBodyBindings {
	s.TemplateId = &v
	return s
}

func (s *BatchBindTemplateResponseBodyBindings) Validate() error {
	return dara.Validate(s)
}
