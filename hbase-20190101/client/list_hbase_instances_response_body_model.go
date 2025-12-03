// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *ListHBaseInstancesResponseBodyInstances) *ListHBaseInstancesResponseBody
	GetInstances() *ListHBaseInstancesResponseBodyInstances
	SetRequestId(v string) *ListHBaseInstancesResponseBody
	GetRequestId() *string
}

type ListHBaseInstancesResponseBody struct {
	Instances *ListHBaseInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 89F81C30-320B-4550-91DB-C37C81D2358F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHBaseInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBody) GetInstances() *ListHBaseInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListHBaseInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHBaseInstancesResponseBody) SetInstances(v *ListHBaseInstancesResponseBodyInstances) *ListHBaseInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListHBaseInstancesResponseBody) SetRequestId(v string) *ListHBaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHBaseInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHBaseInstancesResponseBodyInstances struct {
	Instance []*ListHBaseInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListHBaseInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstances) GetInstance() []*ListHBaseInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *ListHBaseInstancesResponseBodyInstances) SetInstance(v []*ListHBaseInstancesResponseBodyInstancesInstance) *ListHBaseInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstances) Validate() error {
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

type ListHBaseInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// hb-t4naqsay5gn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// name_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) SetIsDefault(v bool) *ListHBaseInstancesResponseBodyInstancesInstance {
	s.IsDefault = &v
	return s
}

func (s *ListHBaseInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
