// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *GetNodeCodeRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *GetNodeCodeRequest
	GetProjectEnv() *string
}

type GetNodeCodeRequest struct {
	// The ID of the node.
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

func (s GetNodeCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeCodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeCodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeCodeRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetNodeCodeRequest) SetNodeId(v int64) *GetNodeCodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetNodeCodeRequest) SetProjectEnv(v string) *GetNodeCodeRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetNodeCodeRequest) Validate() error {
	return dara.Validate(s)
}
