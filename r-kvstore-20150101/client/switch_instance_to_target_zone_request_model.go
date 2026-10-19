// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceToTargetZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchInstanceToTargetZoneRequest
	GetInstanceId() *string
	SetNodeId(v string) *SwitchInstanceToTargetZoneRequest
	GetNodeId() *string
	SetSwitchType(v string) *SwitchInstanceToTargetZoneRequest
	GetSwitchType() *string
	SetTargetZoneId(v string) *SwitchInstanceToTargetZoneRequest
	GetTargetZoneId() *string
}

type SwitchInstanceToTargetZoneRequest struct {
	// Instance ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the child node in a cluster instance.
	//
	// > If you do not specify this parameter, all shards are switched to the destination zone.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The switch mode. Valid values:
	//
	// 	- **ReliabilityPriority*	- (default): Reliability priority. A primary/secondary switchover is performed only when data synchronization between the primary and secondary nodes has no latency. This ensures that no data is lost. However, in scenarios with a large write volume and continuous synchronization latency, the switchover may fail.
	//
	// 	- **AvailablePriority**: Availability priority. A primary/secondary switchover is performed immediately regardless of the synchronization latency between the primary and secondary nodes. A small amount of data may be lost.
	//
	// > Select a switch mode based on your business requirements for data and services.
	//
	// example:
	//
	// ReliabilityPriority
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
	// The ID of the destination zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s SwitchInstanceToTargetZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceToTargetZoneRequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceToTargetZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchInstanceToTargetZoneRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *SwitchInstanceToTargetZoneRequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *SwitchInstanceToTargetZoneRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *SwitchInstanceToTargetZoneRequest) SetInstanceId(v string) *SwitchInstanceToTargetZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceToTargetZoneRequest) SetNodeId(v string) *SwitchInstanceToTargetZoneRequest {
	s.NodeId = &v
	return s
}

func (s *SwitchInstanceToTargetZoneRequest) SetSwitchType(v string) *SwitchInstanceToTargetZoneRequest {
	s.SwitchType = &v
	return s
}

func (s *SwitchInstanceToTargetZoneRequest) SetTargetZoneId(v string) *SwitchInstanceToTargetZoneRequest {
	s.TargetZoneId = &v
	return s
}

func (s *SwitchInstanceToTargetZoneRequest) Validate() error {
	return dara.Validate(s)
}
