// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateRenderingProjectInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v string) *DisassociateRenderingProjectInstancesResponseBody
	GetFailedInstanceCount() *string
	SetFailedInstances(v []*DisassociateRenderingProjectInstancesResponseBodyFailedInstances) *DisassociateRenderingProjectInstancesResponseBody
	GetFailedInstances() []*DisassociateRenderingProjectInstancesResponseBodyFailedInstances
	SetRequestId(v string) *DisassociateRenderingProjectInstancesResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v string) *DisassociateRenderingProjectInstancesResponseBody
	GetSuccessInstanceCount() *string
	SetSuccessInstances(v []*DisassociateRenderingProjectInstancesResponseBodySuccessInstances) *DisassociateRenderingProjectInstancesResponseBody
	GetSuccessInstances() []*DisassociateRenderingProjectInstancesResponseBodySuccessInstances
}

type DisassociateRenderingProjectInstancesResponseBody struct {
	// The number of cloud application service instances that failed to dissociate.
	//
	// example:
	//
	// 0
	FailedInstanceCount *string `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	// A list of instances that failed to dissociate.
	FailedInstances []*DisassociateRenderingProjectInstancesResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of cloud application service instances that were successfully dissociated.
	//
	// example:
	//
	// 5
	SuccessInstanceCount *string `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	// A list of instances that were successfully dissociated.
	SuccessInstances []*DisassociateRenderingProjectInstancesResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s DisassociateRenderingProjectInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesResponseBody) GetFailedInstanceCount() *string {
	return s.FailedInstanceCount
}

func (s *DisassociateRenderingProjectInstancesResponseBody) GetFailedInstances() []*DisassociateRenderingProjectInstancesResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *DisassociateRenderingProjectInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateRenderingProjectInstancesResponseBody) GetSuccessInstanceCount() *string {
	return s.SuccessInstanceCount
}

func (s *DisassociateRenderingProjectInstancesResponseBody) GetSuccessInstances() []*DisassociateRenderingProjectInstancesResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *DisassociateRenderingProjectInstancesResponseBody) SetFailedInstanceCount(v string) *DisassociateRenderingProjectInstancesResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBody) SetFailedInstances(v []*DisassociateRenderingProjectInstancesResponseBodyFailedInstances) *DisassociateRenderingProjectInstancesResponseBody {
	s.FailedInstances = v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBody) SetRequestId(v string) *DisassociateRenderingProjectInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBody) SetSuccessInstanceCount(v string) *DisassociateRenderingProjectInstancesResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBody) SetSuccessInstances(v []*DisassociateRenderingProjectInstancesResponseBodySuccessInstances) *DisassociateRenderingProjectInstancesResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBody) Validate() error {
	if s.FailedInstances != nil {
		for _, item := range s.FailedInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessInstances != nil {
		for _, item := range s.SuccessInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DisassociateRenderingProjectInstancesResponseBodyFailedInstances struct {
	// The reason for failure.
	//
	// example:
	//
	// 会话中
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A cloud application service instance ID.
	//
	// example:
	//
	// render-421cd2a1125947c19fcd5c7dd2c7d31e
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DisassociateRenderingProjectInstancesResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesResponseBodyFailedInstances) GetMessage() *string {
	return s.Message
}

func (s *DisassociateRenderingProjectInstancesResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DisassociateRenderingProjectInstancesResponseBodyFailedInstances) SetMessage(v string) *DisassociateRenderingProjectInstancesResponseBodyFailedInstances {
	s.Message = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBodyFailedInstances) SetRenderingInstanceId(v string) *DisassociateRenderingProjectInstancesResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type DisassociateRenderingProjectInstancesResponseBodySuccessInstances struct {
	// The result message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A cloud application service instance ID.
	//
	// example:
	//
	// render-e6cf423c787e4e43b460a788da254fe3
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DisassociateRenderingProjectInstancesResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s DisassociateRenderingProjectInstancesResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *DisassociateRenderingProjectInstancesResponseBodySuccessInstances) GetMessage() *string {
	return s.Message
}

func (s *DisassociateRenderingProjectInstancesResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DisassociateRenderingProjectInstancesResponseBodySuccessInstances) SetMessage(v string) *DisassociateRenderingProjectInstancesResponseBodySuccessInstances {
	s.Message = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBodySuccessInstances) SetRenderingInstanceId(v string) *DisassociateRenderingProjectInstancesResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *DisassociateRenderingProjectInstancesResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
