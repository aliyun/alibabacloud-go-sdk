// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHiveEdgeWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *AddHiveEdgeWorkersResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*AddHiveEdgeWorkersResponseBodyFailedInstances) *AddHiveEdgeWorkersResponseBody
	GetFailedInstances() []*AddHiveEdgeWorkersResponseBodyFailedInstances
	SetRequestId(v string) *AddHiveEdgeWorkersResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *AddHiveEdgeWorkersResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*AddHiveEdgeWorkersResponseBodySuccessInstances) *AddHiveEdgeWorkersResponseBody
	GetSuccessInstances() []*AddHiveEdgeWorkersResponseBodySuccessInstances
}

type AddHiveEdgeWorkersResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *int32                                           `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*AddHiveEdgeWorkersResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *int32                                            `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*AddHiveEdgeWorkersResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s AddHiveEdgeWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *AddHiveEdgeWorkersResponseBody) GetFailedInstances() []*AddHiveEdgeWorkersResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *AddHiveEdgeWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddHiveEdgeWorkersResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *AddHiveEdgeWorkersResponseBody) GetSuccessInstances() []*AddHiveEdgeWorkersResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *AddHiveEdgeWorkersResponseBody) SetFailedInstanceCount(v int32) *AddHiveEdgeWorkersResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBody) SetFailedInstances(v []*AddHiveEdgeWorkersResponseBodyFailedInstances) *AddHiveEdgeWorkersResponseBody {
	s.FailedInstances = v
	return s
}

func (s *AddHiveEdgeWorkersResponseBody) SetRequestId(v string) *AddHiveEdgeWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBody) SetSuccessInstanceCount(v int32) *AddHiveEdgeWorkersResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBody) SetSuccessInstances(v []*AddHiveEdgeWorkersResponseBodySuccessInstances) *AddHiveEdgeWorkersResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *AddHiveEdgeWorkersResponseBody) Validate() error {
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

type AddHiveEdgeWorkersResponseBodyFailedInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Error 1062 (23000): Duplicate entry \\"hive-4fbf3928d40e43948b98acdb4fb5aaed-ew-1226d588c69449209ee9631\\" for key \\"PRIMARY\\"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddHiveEdgeWorkersResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersResponseBodyFailedInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHiveEdgeWorkersResponseBodyFailedInstances) GetMessage() *string {
	return s.Message
}

func (s *AddHiveEdgeWorkersResponseBodyFailedInstances) SetInstanceId(v string) *AddHiveEdgeWorkersResponseBodyFailedInstances {
	s.InstanceId = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBodyFailedInstances) SetMessage(v string) *AddHiveEdgeWorkersResponseBodyFailedInstances {
	s.Message = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type AddHiveEdgeWorkersResponseBodySuccessInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddHiveEdgeWorkersResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersResponseBodySuccessInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHiveEdgeWorkersResponseBodySuccessInstances) GetMessage() *string {
	return s.Message
}

func (s *AddHiveEdgeWorkersResponseBodySuccessInstances) SetInstanceId(v string) *AddHiveEdgeWorkersResponseBodySuccessInstances {
	s.InstanceId = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBodySuccessInstances) SetMessage(v string) *AddHiveEdgeWorkersResponseBodySuccessInstances {
	s.Message = &v
	return s
}

func (s *AddHiveEdgeWorkersResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
