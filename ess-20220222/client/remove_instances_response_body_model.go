// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredInstances(v []*RemoveInstancesResponseBodyIgnoredInstances) *RemoveInstancesResponseBody
	GetIgnoredInstances() []*RemoveInstancesResponseBodyIgnoredInstances
	SetRequestId(v string) *RemoveInstancesResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *RemoveInstancesResponseBody
	GetScalingActivityId() *string
}

type RemoveInstancesResponseBody struct {
	IgnoredInstances []*RemoveInstancesResponseBodyIgnoredInstances `json:"IgnoredInstances,omitempty" xml:"IgnoredInstances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp175o6f6ego3r2j****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RemoveInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBody) GetIgnoredInstances() []*RemoveInstancesResponseBodyIgnoredInstances {
	return s.IgnoredInstances
}

func (s *RemoveInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstancesResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *RemoveInstancesResponseBody) SetIgnoredInstances(v []*RemoveInstancesResponseBodyIgnoredInstances) *RemoveInstancesResponseBody {
	s.IgnoredInstances = v
	return s
}

func (s *RemoveInstancesResponseBody) SetRequestId(v string) *RemoveInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstancesResponseBody) SetScalingActivityId(v string) *RemoveInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *RemoveInstancesResponseBody) Validate() error {
	if s.IgnoredInstances != nil {
		for _, item := range s.IgnoredInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveInstancesResponseBodyIgnoredInstances struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveInstancesResponseBodyIgnoredInstances) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesResponseBodyIgnoredInstances) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) GetCode() *string {
	return s.Code
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) GetMessage() *string {
	return s.Message
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) SetCode(v string) *RemoveInstancesResponseBodyIgnoredInstances {
	s.Code = &v
	return s
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) SetInstanceId(v string) *RemoveInstancesResponseBodyIgnoredInstances {
	s.InstanceId = &v
	return s
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) SetMessage(v string) *RemoveInstancesResponseBodyIgnoredInstances {
	s.Message = &v
	return s
}

func (s *RemoveInstancesResponseBodyIgnoredInstances) Validate() error {
	return dara.Validate(s)
}
