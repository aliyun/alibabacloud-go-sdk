// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BindTemplateResponseBody
	GetInstanceId() *string
	SetInstanceType(v string) *BindTemplateResponseBody
	GetInstanceType() *string
	SetRequestId(v string) *BindTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *BindTemplateResponseBody
	GetTemplateId() *string
}

type BindTemplateResponseBody struct {
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
}

func (s BindTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *BindTemplateResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindTemplateResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BindTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BindTemplateResponseBody) SetInstanceId(v string) *BindTemplateResponseBody {
	s.InstanceId = &v
	return s
}

func (s *BindTemplateResponseBody) SetInstanceType(v string) *BindTemplateResponseBody {
	s.InstanceType = &v
	return s
}

func (s *BindTemplateResponseBody) SetRequestId(v string) *BindTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindTemplateResponseBody) SetTemplateId(v string) *BindTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *BindTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
