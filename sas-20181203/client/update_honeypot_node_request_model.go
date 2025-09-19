// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableProbeNum(v int32) *UpdateHoneypotNodeRequest
	GetAvailableProbeNum() *int32
	SetNodeId(v string) *UpdateHoneypotNodeRequest
	GetNodeId() *string
	SetNodeName(v string) *UpdateHoneypotNodeRequest
	GetNodeName() *string
	SetSecurityGroupProbeIpList(v []*string) *UpdateHoneypotNodeRequest
	GetSecurityGroupProbeIpList() []*string
}

type UpdateHoneypotNodeRequest struct {
	// The number of available probes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	AvailableProbeNum *int32 `json:"AvailableProbeNum,omitempty" xml:"AvailableProbeNum,omitempty"`
	// The ID of the management node.
	//
	// > You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to query the IDs of management nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67ab3f4c-3db5-4fc3-b51f-00f8bfabfa08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the management node.
	//
	// This parameter is required.
	//
	// example:
	//
	// HoneypotNodeTest
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The CIDR blocks that are allowed to access the management node.
	SecurityGroupProbeIpList []*string `json:"SecurityGroupProbeIpList,omitempty" xml:"SecurityGroupProbeIpList,omitempty" type:"Repeated"`
}

func (s UpdateHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotNodeRequest) GetAvailableProbeNum() *int32 {
	return s.AvailableProbeNum
}

func (s *UpdateHoneypotNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateHoneypotNodeRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *UpdateHoneypotNodeRequest) GetSecurityGroupProbeIpList() []*string {
	return s.SecurityGroupProbeIpList
}

func (s *UpdateHoneypotNodeRequest) SetAvailableProbeNum(v int32) *UpdateHoneypotNodeRequest {
	s.AvailableProbeNum = &v
	return s
}

func (s *UpdateHoneypotNodeRequest) SetNodeId(v string) *UpdateHoneypotNodeRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateHoneypotNodeRequest) SetNodeName(v string) *UpdateHoneypotNodeRequest {
	s.NodeName = &v
	return s
}

func (s *UpdateHoneypotNodeRequest) SetSecurityGroupProbeIpList(v []*string) *UpdateHoneypotNodeRequest {
	s.SecurityGroupProbeIpList = v
	return s
}

func (s *UpdateHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
