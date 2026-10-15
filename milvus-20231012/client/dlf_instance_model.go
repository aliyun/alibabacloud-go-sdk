// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDlfInstance interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DlfInstance
	GetInstanceId() *string
	SetInstanceName(v string) *DlfInstance
	GetInstanceName() *string
	SetInstanceType(v string) *DlfInstance
	GetInstanceType() *string
	SetRegionId(v string) *DlfInstance
	GetRegionId() *string
	SetStatus(v string) *DlfInstance
	GetStatus() *string
}

type DlfInstance struct {
	// The DLF instance ID.
	//
	// example:
	//
	// dlf-cn-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The DLF instance name.
	//
	// example:
	//
	// my-dlf-instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// DLF_2_0
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DlfInstance) String() string {
	return dara.Prettify(s)
}

func (s DlfInstance) GoString() string {
	return s.String()
}

func (s *DlfInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DlfInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DlfInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DlfInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DlfInstance) GetStatus() *string {
	return s.Status
}

func (s *DlfInstance) SetInstanceId(v string) *DlfInstance {
	s.InstanceId = &v
	return s
}

func (s *DlfInstance) SetInstanceName(v string) *DlfInstance {
	s.InstanceName = &v
	return s
}

func (s *DlfInstance) SetInstanceType(v string) *DlfInstance {
	s.InstanceType = &v
	return s
}

func (s *DlfInstance) SetRegionId(v string) *DlfInstance {
	s.RegionId = &v
	return s
}

func (s *DlfInstance) SetStatus(v string) *DlfInstance {
	s.Status = &v
	return s
}

func (s *DlfInstance) Validate() error {
	return dara.Validate(s)
}
