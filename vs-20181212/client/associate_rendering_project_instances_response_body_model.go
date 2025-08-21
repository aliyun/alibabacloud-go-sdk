// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRenderingProjectInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v string) *AssociateRenderingProjectInstancesResponseBody
	GetFailedInstanceCount() *string
	SetFailedInstances(v []*AssociateRenderingProjectInstancesResponseBodyFailedInstances) *AssociateRenderingProjectInstancesResponseBody
	GetFailedInstances() []*AssociateRenderingProjectInstancesResponseBodyFailedInstances
	SetRequestId(v string) *AssociateRenderingProjectInstancesResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v string) *AssociateRenderingProjectInstancesResponseBody
	GetSuccessInstanceCount() *string
	SetSuccessInstances(v []*AssociateRenderingProjectInstancesResponseBodySuccessInstances) *AssociateRenderingProjectInstancesResponseBody
	GetSuccessInstances() []*AssociateRenderingProjectInstancesResponseBodySuccessInstances
}

type AssociateRenderingProjectInstancesResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *string                                                          `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*AssociateRenderingProjectInstancesResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *string                                                           `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*AssociateRenderingProjectInstancesResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s AssociateRenderingProjectInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesResponseBody) GetFailedInstanceCount() *string {
	return s.FailedInstanceCount
}

func (s *AssociateRenderingProjectInstancesResponseBody) GetFailedInstances() []*AssociateRenderingProjectInstancesResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *AssociateRenderingProjectInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateRenderingProjectInstancesResponseBody) GetSuccessInstanceCount() *string {
	return s.SuccessInstanceCount
}

func (s *AssociateRenderingProjectInstancesResponseBody) GetSuccessInstances() []*AssociateRenderingProjectInstancesResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *AssociateRenderingProjectInstancesResponseBody) SetFailedInstanceCount(v string) *AssociateRenderingProjectInstancesResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBody) SetFailedInstances(v []*AssociateRenderingProjectInstancesResponseBodyFailedInstances) *AssociateRenderingProjectInstancesResponseBody {
	s.FailedInstances = v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBody) SetRequestId(v string) *AssociateRenderingProjectInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBody) SetSuccessInstanceCount(v string) *AssociateRenderingProjectInstancesResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBody) SetSuccessInstances(v []*AssociateRenderingProjectInstancesResponseBodySuccessInstances) *AssociateRenderingProjectInstancesResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type AssociateRenderingProjectInstancesResponseBodyFailedInstances struct {
	// example:
	//
	// invalid id
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// render-b45f28650ffe4591bf4c5c95996a428c
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s AssociateRenderingProjectInstancesResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesResponseBodyFailedInstances) GetMessage() *string {
	return s.Message
}

func (s *AssociateRenderingProjectInstancesResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *AssociateRenderingProjectInstancesResponseBodyFailedInstances) SetMessage(v string) *AssociateRenderingProjectInstancesResponseBodyFailedInstances {
	s.Message = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBodyFailedInstances) SetRenderingInstanceId(v string) *AssociateRenderingProjectInstancesResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type AssociateRenderingProjectInstancesResponseBodySuccessInstances struct {
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// render-5130e2feb23f442fb9456a3d977f03d4
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s AssociateRenderingProjectInstancesResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s AssociateRenderingProjectInstancesResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *AssociateRenderingProjectInstancesResponseBodySuccessInstances) GetMessage() *string {
	return s.Message
}

func (s *AssociateRenderingProjectInstancesResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *AssociateRenderingProjectInstancesResponseBodySuccessInstances) SetMessage(v string) *AssociateRenderingProjectInstancesResponseBodySuccessInstances {
	s.Message = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBodySuccessInstances) SetRenderingInstanceId(v string) *AssociateRenderingProjectInstancesResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *AssociateRenderingProjectInstancesResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
