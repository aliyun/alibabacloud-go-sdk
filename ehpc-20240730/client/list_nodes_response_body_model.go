// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody
	GetNodes() []*ListNodesResponseBodyNodes
	SetPageNumber(v int32) *ListNodesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNodesResponseBody
	GetTotalCount() *int32
}

type ListNodesResponseBody struct {
	// The information about the nodes.
	Nodes []*ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89A1AC0F-4A6C-4F3D-98F9-BEF9A823****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetNodes() []*ListNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListNodesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodesResponseBody) SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetPageNumber(v int32) *ListNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBody) SetPageSize(v int32) *ListNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyNodes struct {
	// The time when the node was created.
	//
	// example:
	//
	// 2020-06-09T06:22:02.000Z
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-8vbe4av4gededlqg****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 2020-06-09T06:22:02.000Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The hostname of the node.
	//
	// example:
	//
	// edas.cn-shanghai.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Indicates whether hyper-threading is enabled.
	//
	// example:
	//
	// true
	HtEnabled *bool `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	// The instance ID of the node.
	//
	// example:
	//
	// i-bp15707mys2rsy0j****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image ID of the node.
	//
	// example:
	//
	// centos_7_06_64_20G_alibase_20190711.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type of the node.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The VPC IP address of the node.
	//
	// example:
	//
	// ``172.16.**.**``
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// Indicates whether deletion protection is enabled for the node. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The public IP address of the node.
	//
	// example:
	//
	// ``172.16.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The name of the queue to which the node belongs.
	//
	// example:
	//
	// autoque3
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The bidding policy of the node. Valid values:
	//
	// 	- NoSpot: The instances of the compute node are pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The node is a preemptible instance for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The node state in the scheduler.
	//
	// example:
	//
	// active
	StateInSched *string `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	// The node state. Valid values:
	//
	// 	- uninit: The node is being installed.
	//
	// 	- initing: The node is being initialized.
	//
	// 	- running: The node is running.
	//
	// 	- releasing: The node is being released.
	//
	// 	- stopped: The node is stopped.
	//
	// 	- exception: The node has run into an exception.
	//
	// 	- untracking: The node is not added to the cluster.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The hardware configurations of the node.
	TotalResources *ListNodesResponseBodyNodesTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	// The vSwitch ID of the node.
	//
	// example:
	//
	// vsw-bp1e47optm9g58zcu****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1gnu8br4ay7beb2w****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) GetAddTime() *string {
	return s.AddTime
}

func (s *ListNodesResponseBodyNodes) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ListNodesResponseBodyNodes) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListNodesResponseBodyNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListNodesResponseBodyNodes) GetHtEnabled() *bool {
	return s.HtEnabled
}

func (s *ListNodesResponseBodyNodes) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *ListNodesResponseBodyNodes) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListNodesResponseBodyNodes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListNodesResponseBodyNodes) GetKeepAlive() *bool {
	return s.KeepAlive
}

func (s *ListNodesResponseBodyNodes) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *ListNodesResponseBodyNodes) GetQueueName() *string {
	return s.QueueName
}

func (s *ListNodesResponseBodyNodes) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ListNodesResponseBodyNodes) GetStateInSched() *string {
	return s.StateInSched
}

func (s *ListNodesResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *ListNodesResponseBodyNodes) GetTotalResources() *ListNodesResponseBodyNodesTotalResources {
	return s.TotalResources
}

func (s *ListNodesResponseBodyNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListNodesResponseBodyNodes) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNodesResponseBodyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNodesResponseBodyNodes) SetAddTime(v string) *ListNodesResponseBodyNodes {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetDeploymentSetId(v string) *ListNodesResponseBodyNodes {
	s.DeploymentSetId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExpiredTime(v string) *ListNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHostname(v string) *ListNodesResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHtEnabled(v bool) *ListNodesResponseBodyNodes {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetId(v string) *ListNodesResponseBodyNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetImageId(v string) *ListNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetInstanceType(v string) *ListNodesResponseBodyNodes {
	s.InstanceType = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetIpAddress(v string) *ListNodesResponseBodyNodes {
	s.IpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetKeepAlive(v bool) *ListNodesResponseBodyNodes {
	s.KeepAlive = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetPublicIpAddress(v string) *ListNodesResponseBodyNodes {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetQueueName(v string) *ListNodesResponseBodyNodes {
	s.QueueName = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetSpotStrategy(v string) *ListNodesResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStateInSched(v string) *ListNodesResponseBodyNodes {
	s.StateInSched = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStatus(v string) *ListNodesResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetTotalResources(v *ListNodesResponseBodyNodesTotalResources) *ListNodesResponseBodyNodes {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVSwitchId(v string) *ListNodesResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVpcId(v string) *ListNodesResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetZoneId(v string) *ListNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) Validate() error {
	if s.TotalResources != nil {
		if err := s.TotalResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyNodesTotalResources struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 0
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The amount of memory. Unit: GiB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesTotalResources) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyNodesTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesTotalResources) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListNodesResponseBodyNodesTotalResources) GetGpu() *int32 {
	return s.Gpu
}

func (s *ListNodesResponseBodyNodesTotalResources) GetMemory() *int32 {
	return s.Memory
}

func (s *ListNodesResponseBodyNodesTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Memory = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) Validate() error {
	return dara.Validate(s)
}
