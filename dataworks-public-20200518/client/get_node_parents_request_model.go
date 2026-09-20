// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeParentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *GetNodeParentsRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *GetNodeParentsRequest
	GetProjectEnv() *string
}

type GetNodeParentsRequest struct {
	// The node ID. You can view the node ID in the Operation Center of the DataWorks console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345456211234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment identifier of the Operation Center. PROD indicates the production environment. DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s GetNodeParentsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeParentsRequest) GoString() string {
	return s.String()
}

func (s *GetNodeParentsRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeParentsRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetNodeParentsRequest) SetNodeId(v int64) *GetNodeParentsRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeParentsRequest) SetProjectEnv(v string) *GetNodeParentsRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetNodeParentsRequest) Validate() error {
	return dara.Validate(s)
}
