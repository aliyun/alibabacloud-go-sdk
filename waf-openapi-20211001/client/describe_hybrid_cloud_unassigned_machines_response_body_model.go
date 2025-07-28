// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnassignedMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHybridCloudUnassignedMachinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeHybridCloudUnassignedMachinesResponseBody
	GetTotalCount() *int64
	SetUnassignedMachines(v []*DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) *DescribeHybridCloudUnassignedMachinesResponseBody
	GetUnassignedMachines() []*DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines
}

type DescribeHybridCloudUnassignedMachinesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3EBCFCE9-4A3C-5E01-915D-691B****510A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 28
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The servers that are not assigned to the hybrid cloud cluster.
	UnassignedMachines []*DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines `json:"UnassignedMachines,omitempty" xml:"UnassignedMachines,omitempty" type:"Repeated"`
}

func (s DescribeHybridCloudUnassignedMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnassignedMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) GetUnassignedMachines() []*DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	return s.UnassignedMachines
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) SetRequestId(v string) *DescribeHybridCloudUnassignedMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) SetTotalCount(v int64) *DescribeHybridCloudUnassignedMachinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) SetUnassignedMachines(v []*DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) *DescribeHybridCloudUnassignedMachinesResponseBody {
	s.UnassignedMachines = v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 16
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// exampleName
	CustomName *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
	// The host name.
	//
	// example:
	//
	// online-xagent1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 1.X.X.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The media access control (MAC) address of the device.
	//
	// example:
	//
	// 00163e2686ac
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The memory size. Unit: KB. A conversion factor of 1,000 is used.
	//
	// example:
	//
	// 31580872
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 78db009ab6cf055a9085f9f4****ae3a
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
}

func (s DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetCpu() *int64 {
	return s.Cpu
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetCustomName() *string {
	return s.CustomName
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetMac() *string {
	return s.Mac
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) GetMid() *string {
	return s.Mid
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetCpu(v int64) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.Cpu = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetCustomName(v string) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.CustomName = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetHostName(v string) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetIp(v string) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetMac(v string) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.Mac = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetMemory(v int64) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.Memory = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) SetMid(v string) *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines {
	s.Mid = &v
	return s
}

func (s *DescribeHybridCloudUnassignedMachinesResponseBodyUnassignedMachines) Validate() error {
	return dara.Validate(s)
}
