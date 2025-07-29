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
	// Specifies whether to enable automatic instance restart.
	//
	// **
	//
	// **Warning*	- This parameter is deprecated. Any configured values will be ignored.
	//
	// example:
	//
	// true
	AutoRestart *bool `json:"auto_restart,omitempty" xml:"auto_restart,omitempty"`
	// The list of nodes. If not specified, all nodes in the node pool are selected.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The list of repair operations to execute. If not specified, all repair operations are executed. Typically, you do not need to specify this parameter.
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
	return dara.Validate(s)
}

type RepairClusterNodePoolRequestOperations struct {
	// The parameters of a repair operation.
	Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
	// The ID of a repair operation.
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
