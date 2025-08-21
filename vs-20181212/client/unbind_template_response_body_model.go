// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnbindTemplateResponseBody
	GetInstanceId() *string
	SetInstanceType(v string) *UnbindTemplateResponseBody
	GetInstanceType() *string
	SetRequestId(v string) *UnbindTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *UnbindTemplateResponseBody
	GetTemplateId() *string
	SetTemplateType(v string) *UnbindTemplateResponseBody
	GetTemplateType() *string
}

type UnbindTemplateResponseBody struct {
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
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// record
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s UnbindTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindTemplateResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindTemplateResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnbindTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UnbindTemplateResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UnbindTemplateResponseBody) SetInstanceId(v string) *UnbindTemplateResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetInstanceType(v string) *UnbindTemplateResponseBody {
	s.InstanceType = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetRequestId(v string) *UnbindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetTemplateId(v string) *UnbindTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *UnbindTemplateResponseBody) SetTemplateType(v string) *UnbindTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *UnbindTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
