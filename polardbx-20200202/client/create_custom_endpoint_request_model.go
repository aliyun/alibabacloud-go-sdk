// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateCustomEndpointRequest
	GetDBInstanceName() *string
	SetName(v string) *CreateCustomEndpointRequest
	GetName() *string
	SetNodeAutoEnter(v bool) *CreateCustomEndpointRequest
	GetNodeAutoEnter() *bool
	SetNodeIds(v string) *CreateCustomEndpointRequest
	GetNodeIds() *string
	SetNodeRole(v string) *CreateCustomEndpointRequest
	GetNodeRole() *string
	SetRegionId(v string) *CreateCustomEndpointRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateCustomEndpointRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateCustomEndpointRequest
	GetVpcId() *string
}

type CreateCustomEndpointRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the access control instance. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. The name can contain digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ext-win-live-17
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether a node automatically joins the cluster and starts providing services after the node is added or recovered.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	NodeAutoEnter *bool `json:"NodeAutoEnter,omitempty" xml:"NodeAutoEnter,omitempty"`
	// The IDs of the monitored nodes when RemindUnit (object type) is set to NODE (node). Separate multiple IDs with commas (,). A maximum of 50 nodes can be monitored per rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1l5kfgw8****c3iv
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// To query the metrics of a read-only node in a cloud-native read/write splitting architecture instance, set this parameter to **READONLY*	- and specify the **NodeId*	- parameter.
	//
	// >  In other cases, you do not need to specify this parameter or you can set it to **MASTER**.
	//
	// example:
	//
	// READONLY
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the endpoint resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ndoug37dtwoedlmru0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateCustomEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomEndpointRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCustomEndpointRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomEndpointRequest) GetNodeAutoEnter() *bool {
	return s.NodeAutoEnter
}

func (s *CreateCustomEndpointRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *CreateCustomEndpointRequest) GetNodeRole() *string {
	return s.NodeRole
}

func (s *CreateCustomEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomEndpointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateCustomEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateCustomEndpointRequest) SetDBInstanceName(v string) *CreateCustomEndpointRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetName(v string) *CreateCustomEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetNodeAutoEnter(v bool) *CreateCustomEndpointRequest {
	s.NodeAutoEnter = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetNodeIds(v string) *CreateCustomEndpointRequest {
	s.NodeIds = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetNodeRole(v string) *CreateCustomEndpointRequest {
	s.NodeRole = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetRegionId(v string) *CreateCustomEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetVSwitchId(v string) *CreateCustomEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateCustomEndpointRequest) SetVpcId(v string) *CreateCustomEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateCustomEndpointRequest) Validate() error {
	return dara.Validate(s)
}
