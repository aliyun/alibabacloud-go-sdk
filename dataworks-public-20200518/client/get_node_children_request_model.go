// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeChildrenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *GetNodeChildrenRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *GetNodeChildrenRequest
	GetProjectEnv() *string
}

type GetNodeChildrenRequest struct {
	// The node ID. You can go to the Operation Center page in the DataWorks console to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456878
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment type of Operation Center. Valid values: PROD and DEV. The value PROD indicates the production environment, and the value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s GetNodeChildrenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeChildrenRequest) GoString() string {
	return s.String()
}

func (s *GetNodeChildrenRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeChildrenRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetNodeChildrenRequest) SetNodeId(v int64) *GetNodeChildrenRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeChildrenRequest) SetProjectEnv(v string) *GetNodeChildrenRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetNodeChildrenRequest) Validate() error {
	return dara.Validate(s)
}
