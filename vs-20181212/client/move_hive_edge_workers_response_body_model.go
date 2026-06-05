// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHiveEdgeWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *MoveHiveEdgeWorkersResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*MoveHiveEdgeWorkersResponseBodyFailedInstances) *MoveHiveEdgeWorkersResponseBody
	GetFailedInstances() []*MoveHiveEdgeWorkersResponseBodyFailedInstances
	SetRequestId(v string) *MoveHiveEdgeWorkersResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *MoveHiveEdgeWorkersResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*MoveHiveEdgeWorkersResponseBodySuccessInstances) *MoveHiveEdgeWorkersResponseBody
	GetSuccessInstances() []*MoveHiveEdgeWorkersResponseBodySuccessInstances
}

type MoveHiveEdgeWorkersResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *int32                                            `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*MoveHiveEdgeWorkersResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *int32                                             `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*MoveHiveEdgeWorkersResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s MoveHiveEdgeWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *MoveHiveEdgeWorkersResponseBody) GetFailedInstances() []*MoveHiveEdgeWorkersResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *MoveHiveEdgeWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveHiveEdgeWorkersResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *MoveHiveEdgeWorkersResponseBody) GetSuccessInstances() []*MoveHiveEdgeWorkersResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *MoveHiveEdgeWorkersResponseBody) SetFailedInstanceCount(v int32) *MoveHiveEdgeWorkersResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBody) SetFailedInstances(v []*MoveHiveEdgeWorkersResponseBodyFailedInstances) *MoveHiveEdgeWorkersResponseBody {
	s.FailedInstances = v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBody) SetRequestId(v string) *MoveHiveEdgeWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBody) SetSuccessInstanceCount(v int32) *MoveHiveEdgeWorkersResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBody) SetSuccessInstances(v []*MoveHiveEdgeWorkersResponseBodySuccessInstances) *MoveHiveEdgeWorkersResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBody) Validate() error {
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

type MoveHiveEdgeWorkersResponseBodyFailedInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Error 1062 (23000): Duplicate entry \\"hive-4fbf3928d40e43948b98acdb4fb5aaed-ew-1226d588c69449209ee9631\\" for key \\"PRIMARY\\"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MoveHiveEdgeWorkersResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersResponseBodyFailedInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveHiveEdgeWorkersResponseBodyFailedInstances) GetMessage() *string {
	return s.Message
}

func (s *MoveHiveEdgeWorkersResponseBodyFailedInstances) SetInstanceId(v string) *MoveHiveEdgeWorkersResponseBodyFailedInstances {
	s.InstanceId = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBodyFailedInstances) SetMessage(v string) *MoveHiveEdgeWorkersResponseBodyFailedInstances {
	s.Message = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type MoveHiveEdgeWorkersResponseBodySuccessInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MoveHiveEdgeWorkersResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersResponseBodySuccessInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveHiveEdgeWorkersResponseBodySuccessInstances) GetMessage() *string {
	return s.Message
}

func (s *MoveHiveEdgeWorkersResponseBodySuccessInstances) SetInstanceId(v string) *MoveHiveEdgeWorkersResponseBodySuccessInstances {
	s.InstanceId = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBodySuccessInstances) SetMessage(v string) *MoveHiveEdgeWorkersResponseBodySuccessInstances {
	s.Message = &v
	return s
}

func (s *MoveHiveEdgeWorkersResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
