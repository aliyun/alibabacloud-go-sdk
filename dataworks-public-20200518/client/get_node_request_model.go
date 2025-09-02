// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *GetNodeRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *GetNodeRequest
	GetProjectEnv() *string
}

type GetNodeRequest struct {
	// The interval at which the node is rerun after the node fails to run.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s GetNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetNodeRequest) SetNodeId(v int64) *GetNodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeRequest) SetProjectEnv(v string) *GetNodeRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetNodeRequest) Validate() error {
	return dara.Validate(s)
}
