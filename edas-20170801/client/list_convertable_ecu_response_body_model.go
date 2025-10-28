// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConvertableEcuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListConvertableEcuResponseBody
	GetCode() *int32
	SetInstanceList(v *ListConvertableEcuResponseBodyInstanceList) *ListConvertableEcuResponseBody
	GetInstanceList() *ListConvertableEcuResponseBodyInstanceList
	SetMessage(v string) *ListConvertableEcuResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConvertableEcuResponseBody
	GetRequestId() *string
}

type ListConvertableEcuResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ECS instances that can be imported to the cluster.
	InstanceList *ListConvertableEcuResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
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
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConvertableEcuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConvertableEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListConvertableEcuResponseBody) GetInstanceList() *ListConvertableEcuResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListConvertableEcuResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConvertableEcuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConvertableEcuResponseBody) SetCode(v int32) *ListConvertableEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListConvertableEcuResponseBody) SetInstanceList(v *ListConvertableEcuResponseBodyInstanceList) *ListConvertableEcuResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListConvertableEcuResponseBody) SetMessage(v string) *ListConvertableEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListConvertableEcuResponseBody) SetRequestId(v string) *ListConvertableEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConvertableEcuResponseBody) Validate() error {
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConvertableEcuResponseBodyInstanceList struct {
	Instance []*ListConvertableEcuResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListConvertableEcuResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListConvertableEcuResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBodyInstanceList) GetInstance() []*ListConvertableEcuResponseBodyInstanceListInstance {
	return s.Instance
}

func (s *ListConvertableEcuResponseBodyInstanceList) SetInstance(v []*ListConvertableEcuResponseBodyInstanceListInstance) *ListConvertableEcuResponseBodyInstanceList {
	s.Instance = v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceList) Validate() error {
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

type ListConvertableEcuResponseBodyInstanceListInstance struct {
	// The number of CPU cores of the ECS instance.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID of the elastic compute units (ECU).
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The elastic IP address (EIP) associated with the ECS instance. The EIP can be changed.
	//
	// example:
	//
	// 13.xx.xxx.xx
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// Indicates whether the ECS instance has expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The private IP address of the ECS instance. This parameter is valid only when the ECS instance is deployed in a VPC.
	//
	// example:
	//
	// 192.168.13.xx
	InnerIp *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2ze7s2v0b***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// worker-k8s
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The size of memory for the ECS instance.
	//
	// example:
	//
	// 4096
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The private IP address of the ECS instance. This parameter is valid only when the ECS instance is deployed in a VPC.
	//
	// example:
	//
	// 192.XX.XX.123
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The public IP address of the ECS instance. This IP address can be used only by the ECS instance.
	//
	// example:
	//
	// 13.xx.xx.xxx
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The ID of the region where the ECS instance is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- Pending: The instance is being created.
	//
	// 	- Running: The instance is running.
	//
	// 	- Starting: The instance is being started.
	//
	// 	- Stopping: The instance is being stopped.
	//
	// 	- Stopped: The instance is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zef6ob8m************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// jianwei-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListConvertableEcuResponseBodyInstanceListInstance) String() string {
	return dara.Prettify(s)
}

func (s ListConvertableEcuResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetEcuId() *string {
	return s.EcuId
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetEip() *string {
	return s.Eip
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetExpired() *bool {
	return s.Expired
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetInnerIp() *string {
	return s.InnerIp
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetMem() *int32 {
	return s.Mem
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetStatus() *string {
	return s.Status
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) GetVpcName() *string {
	return s.VpcName
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetCpu(v int32) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Cpu = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetEcuId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.EcuId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetEip(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Eip = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetExpired(v bool) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Expired = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInnerIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InnerIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInstanceId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInstanceName(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetMem(v int32) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Mem = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetPrivateIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.PrivateIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetPublicIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.PublicIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetRegionId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.RegionId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetStatus(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Status = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetVpcId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.VpcId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetVpcName(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.VpcName = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) Validate() error {
	return dara.Validate(s)
}
