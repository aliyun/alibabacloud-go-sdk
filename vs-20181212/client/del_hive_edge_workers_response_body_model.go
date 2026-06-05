// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelHiveEdgeWorkersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *DelHiveEdgeWorkersResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*DelHiveEdgeWorkersResponseBodyFailedInstances) *DelHiveEdgeWorkersResponseBody
	GetFailedInstances() []*DelHiveEdgeWorkersResponseBodyFailedInstances
	SetRequestId(v string) *DelHiveEdgeWorkersResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *DelHiveEdgeWorkersResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*DelHiveEdgeWorkersResponseBodySuccessInstances) *DelHiveEdgeWorkersResponseBody
	GetSuccessInstances() []*DelHiveEdgeWorkersResponseBodySuccessInstances
}

type DelHiveEdgeWorkersResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *int32                                           `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*DelHiveEdgeWorkersResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *int32                                            `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*DelHiveEdgeWorkersResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s DelHiveEdgeWorkersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersResponseBody) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *DelHiveEdgeWorkersResponseBody) GetFailedInstances() []*DelHiveEdgeWorkersResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *DelHiveEdgeWorkersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelHiveEdgeWorkersResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *DelHiveEdgeWorkersResponseBody) GetSuccessInstances() []*DelHiveEdgeWorkersResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *DelHiveEdgeWorkersResponseBody) SetFailedInstanceCount(v int32) *DelHiveEdgeWorkersResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBody) SetFailedInstances(v []*DelHiveEdgeWorkersResponseBodyFailedInstances) *DelHiveEdgeWorkersResponseBody {
	s.FailedInstances = v
	return s
}

func (s *DelHiveEdgeWorkersResponseBody) SetRequestId(v string) *DelHiveEdgeWorkersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBody) SetSuccessInstanceCount(v int32) *DelHiveEdgeWorkersResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBody) SetSuccessInstances(v []*DelHiveEdgeWorkersResponseBodySuccessInstances) *DelHiveEdgeWorkersResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *DelHiveEdgeWorkersResponseBody) Validate() error {
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

type DelHiveEdgeWorkersResponseBodyFailedInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Error 1062 (23000): Duplicate entry \\"hive-4fbf3928d40e43948b98acdb4fb5aaed-ew-1226d588c69449209ee9631\\" for key \\"PRIMARY\\"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DelHiveEdgeWorkersResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersResponseBodyFailedInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DelHiveEdgeWorkersResponseBodyFailedInstances) GetMessage() *string {
	return s.Message
}

func (s *DelHiveEdgeWorkersResponseBodyFailedInstances) SetInstanceId(v string) *DelHiveEdgeWorkersResponseBodyFailedInstances {
	s.InstanceId = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBodyFailedInstances) SetMessage(v string) *DelHiveEdgeWorkersResponseBodyFailedInstances {
	s.Message = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type DelHiveEdgeWorkersResponseBodySuccessInstances struct {
	// example:
	//
	// ew-1226d588c69449209ee963161c067b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DelHiveEdgeWorkersResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersResponseBodySuccessInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DelHiveEdgeWorkersResponseBodySuccessInstances) GetMessage() *string {
	return s.Message
}

func (s *DelHiveEdgeWorkersResponseBodySuccessInstances) SetInstanceId(v string) *DelHiveEdgeWorkersResponseBodySuccessInstances {
	s.InstanceId = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBodySuccessInstances) SetMessage(v string) *DelHiveEdgeWorkersResponseBodySuccessInstances {
	s.Message = &v
	return s
}

func (s *DelHiveEdgeWorkersResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
