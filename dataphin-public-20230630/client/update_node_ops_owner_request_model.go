// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOpsOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v *UpdateNodeOpsOwnerRequestCommand) *UpdateNodeOpsOwnerRequest
	GetCommand() *UpdateNodeOpsOwnerRequestCommand
	SetOpTenantId(v int64) *UpdateNodeOpsOwnerRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateNodeOpsOwnerRequest
	GetOpUserId() *string
}

type UpdateNodeOpsOwnerRequest struct {
	// The command for updating O&M owners.
	//
	// This parameter is required.
	Command *UpdateNodeOpsOwnerRequestCommand `json:"Command,omitempty" xml:"Command,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s UpdateNodeOpsOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerRequest) GetCommand() *UpdateNodeOpsOwnerRequestCommand {
	return s.Command
}

func (s *UpdateNodeOpsOwnerRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateNodeOpsOwnerRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateNodeOpsOwnerRequest) SetCommand(v *UpdateNodeOpsOwnerRequestCommand) *UpdateNodeOpsOwnerRequest {
	s.Command = v
	return s
}

func (s *UpdateNodeOpsOwnerRequest) SetOpTenantId(v int64) *UpdateNodeOpsOwnerRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateNodeOpsOwnerRequest) SetOpUserId(v string) *UpdateNodeOpsOwnerRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateNodeOpsOwnerRequest) Validate() error {
	if s.Command != nil {
		if err := s.Command.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateNodeOpsOwnerRequestCommand struct {
	// The list of nodes. Only offline nodes are supported.
	//
	// This parameter is required.
	NodeIdList []*UpdateNodeOpsOwnerRequestCommandNodeIdList `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// The updated O&M owners. Specify a list of user account IDs. A maximum of 50 IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OpsOwnerList []*string `json:"OpsOwnerList,omitempty" xml:"OpsOwnerList,omitempty" type:"Repeated"`
}

func (s UpdateNodeOpsOwnerRequestCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerRequestCommand) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerRequestCommand) GetNodeIdList() []*UpdateNodeOpsOwnerRequestCommandNodeIdList {
	return s.NodeIdList
}

func (s *UpdateNodeOpsOwnerRequestCommand) GetOpsOwnerList() []*string {
	return s.OpsOwnerList
}

func (s *UpdateNodeOpsOwnerRequestCommand) SetNodeIdList(v []*UpdateNodeOpsOwnerRequestCommandNodeIdList) *UpdateNodeOpsOwnerRequestCommand {
	s.NodeIdList = v
	return s
}

func (s *UpdateNodeOpsOwnerRequestCommand) SetOpsOwnerList(v []*string) *UpdateNodeOpsOwnerRequestCommand {
	s.OpsOwnerList = v
	return s
}

func (s *UpdateNodeOpsOwnerRequestCommand) Validate() error {
	if s.NodeIdList != nil {
		for _, item := range s.NodeIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNodeOpsOwnerRequestCommandNodeIdList struct {
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n_8198365584737107968
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node source type. Only offline nodes are supported. Valid values:
	//
	// - DATA_PROCESS: compute node.
	//
	// - PIPELINE: integration node.
	//
	// - BLACK_BOX: logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	NodeFromType *string `json:"NodeFromType,omitempty" xml:"NodeFromType,omitempty"`
	// The node type. Valid values:
	//
	// - DATA_PROCESS: compute node.
	//
	// - PIPELINE_NODE: integration node.
	//
	// - BBOX_LOGIC_TABLE_NODE: logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s UpdateNodeOpsOwnerRequestCommandNodeIdList) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerRequestCommandNodeIdList) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) GetId() *string {
	return s.Id
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) GetNodeFromType() *string {
	return s.NodeFromType
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) GetNodeType() *string {
	return s.NodeType
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) SetId(v string) *UpdateNodeOpsOwnerRequestCommandNodeIdList {
	s.Id = &v
	return s
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) SetNodeFromType(v string) *UpdateNodeOpsOwnerRequestCommandNodeIdList {
	s.NodeFromType = &v
	return s
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) SetNodeType(v string) *UpdateNodeOpsOwnerRequestCommandNodeIdList {
	s.NodeType = &v
	return s
}

func (s *UpdateNodeOpsOwnerRequestCommandNodeIdList) Validate() error {
	return dara.Validate(s)
}
