// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHbaseInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *ListHbaseInstancesResponseBodyInstances) *ListHbaseInstancesResponseBody
	GetInstances() *ListHbaseInstancesResponseBodyInstances
	SetRequestId(v string) *ListHbaseInstancesResponseBody
	GetRequestId() *string
}

type ListHbaseInstancesResponseBody struct {
	Instances *ListHbaseInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHbaseInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHbaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBody) GetInstances() *ListHbaseInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListHbaseInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHbaseInstancesResponseBody) SetInstances(v *ListHbaseInstancesResponseBodyInstances) *ListHbaseInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListHbaseInstancesResponseBody) SetRequestId(v string) *ListHbaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHbaseInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHbaseInstancesResponseBodyInstances struct {
	Instance []*ListHbaseInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListHbaseInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListHbaseInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBodyInstances) GetInstance() []*ListHbaseInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *ListHbaseInstancesResponseBodyInstances) SetInstance(v []*ListHbaseInstancesResponseBodyInstancesInstance) *ListHbaseInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHbaseInstancesResponseBodyInstancesInstance struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s ListHbaseInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s ListHbaseInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) SetIsDefault(v bool) *ListHbaseInstancesResponseBodyInstancesInstance {
	s.IsDefault = &v
	return s
}

func (s *ListHbaseInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
