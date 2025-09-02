// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *UpdateNodeOwnerRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *UpdateNodeOwnerRequest
	GetProjectEnv() *string
	SetUserId(v string) *UpdateNodeOwnerRequest
	GetUserId() *string
}

type UpdateNodeOwnerRequest struct {
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment in which the node runs. Valid values: DEV and PROD. The value DEV indicates the development environment, and the value PROD indicates the production environment.
	//
	// 	- PROD
	//
	// 	- DEV
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner. You can log on to the DataWorks console and move the pointer over the profile picture in the upper-right corner to view the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19337906836551
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateNodeOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOwnerRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeOwnerRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *UpdateNodeOwnerRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *UpdateNodeOwnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateNodeOwnerRequest) SetNodeId(v int64) *UpdateNodeOwnerRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateNodeOwnerRequest) SetProjectEnv(v string) *UpdateNodeOwnerRequest {
	s.ProjectEnv = &v
	return s
}

func (s *UpdateNodeOwnerRequest) SetUserId(v string) *UpdateNodeOwnerRequest {
	s.UserId = &v
	return s
}

func (s *UpdateNodeOwnerRequest) Validate() error {
	return dara.Validate(s)
}
