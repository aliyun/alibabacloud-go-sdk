// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairClusterNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRestart(v bool) *RepairClusterNodePoolRequest
	GetAutoRestart() *bool
	SetNodes(v []*string) *RepairClusterNodePoolRequest
	GetNodes() []*string
	SetOperations(v []*RepairClusterNodePoolRequestOperations) *RepairClusterNodePoolRequest
	GetOperations() []*RepairClusterNodePoolRequestOperations
}

type RepairClusterNodePoolRequest struct {
	// Deprecated
	//
	// **[Deprecated]*	- Specifies whether to restart the instance.
	//
	// example:
	//
	// null
	AutoRestart *bool `json:"auto_restart,omitempty" xml:"auto_restart,omitempty"`
	// The list of nodes.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The repair operations to be performed. If you do not specify this parameter, all repair operations are performed by default. You do not need to specify this parameter in common scenarios.
	Operations []*RepairClusterNodePoolRequestOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
}

func (s RepairClusterNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s RepairClusterNodePoolRequest) GoString() string {
	return s.String()
}

func (s *RepairClusterNodePoolRequest) GetAutoRestart() *bool {
	return s.AutoRestart
}

func (s *RepairClusterNodePoolRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *RepairClusterNodePoolRequest) GetOperations() []*RepairClusterNodePoolRequestOperations {
	return s.Operations
}

func (s *RepairClusterNodePoolRequest) SetAutoRestart(v bool) *RepairClusterNodePoolRequest {
	s.AutoRestart = &v
	return s
}

func (s *RepairClusterNodePoolRequest) SetNodes(v []*string) *RepairClusterNodePoolRequest {
	s.Nodes = v
	return s
}

func (s *RepairClusterNodePoolRequest) SetOperations(v []*RepairClusterNodePoolRequestOperations) *RepairClusterNodePoolRequest {
	s.Operations = v
	return s
}

func (s *RepairClusterNodePoolRequest) Validate() error {
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RepairClusterNodePoolRequestOperations struct {
	// The list of repair operation parameters.
	Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
	// The repair operation ID. Valid values:
	//
	// - restart.kubelet: Restarts kubelet.
	//
	// - restart.docker: Restarts Docker.
	//
	// - restart.containerd: Restarts Containerd.
	//
	// - restart.ntp: Restarts ntpd or chronyd.
	//
	// - remove.containerdContainerInSandbox: Deletes the specified sandbox container in Containerd.
	//
	// - remove.dockerContainerInSandbox: Deletes the specified sandbox container in Docker.
	//
	// - remove.containerdContainer: Deletes the specified container in Containerd.
	//
	// - remove.dockerContainer: Deletes the specified container in Docker.
	//
	// example:
	//
	// remove.containerdContainer
	OperationId *string `json:"operation_id,omitempty" xml:"operation_id,omitempty"`
}

func (s RepairClusterNodePoolRequestOperations) String() string {
	return dara.Prettify(s)
}

func (s RepairClusterNodePoolRequestOperations) GoString() string {
	return s.String()
}

func (s *RepairClusterNodePoolRequestOperations) GetArgs() []*string {
	return s.Args
}

func (s *RepairClusterNodePoolRequestOperations) GetOperationId() *string {
	return s.OperationId
}

func (s *RepairClusterNodePoolRequestOperations) SetArgs(v []*string) *RepairClusterNodePoolRequestOperations {
	s.Args = v
	return s
}

func (s *RepairClusterNodePoolRequestOperations) SetOperationId(v string) *RepairClusterNodePoolRequestOperations {
	s.OperationId = &v
	return s
}

func (s *RepairClusterNodePoolRequestOperations) Validate() error {
	return dara.Validate(s)
}
