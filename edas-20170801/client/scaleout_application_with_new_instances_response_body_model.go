// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleoutApplicationWithNewInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *ScaleoutApplicationWithNewInstancesResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *ScaleoutApplicationWithNewInstancesResponseBody
	GetCode() *int32
	SetInstanceIds(v []*string) *ScaleoutApplicationWithNewInstancesResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *ScaleoutApplicationWithNewInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ScaleoutApplicationWithNewInstancesResponseBody
	GetRequestId() *string
}

type ScaleoutApplicationWithNewInstancesResponseBody struct {
	// The ID of the change process for the scale-out.
	//
	// example:
	//
	// e370c17f-*****-3df0721a327
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of ECS instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// e370c17f-*****-3df0721a327
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleoutApplicationWithNewInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetChangeOrderId(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetCode(v int32) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetInstanceIds(v []*string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetMessage(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetRequestId(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
