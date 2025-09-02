// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSmokeTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v string) *RunSmokeTestRequest
	GetBizdate() *string
	SetName(v string) *RunSmokeTestRequest
	GetName() *string
	SetNodeId(v int64) *RunSmokeTestRequest
	GetNodeId() *int64
	SetNodeParams(v string) *RunSmokeTestRequest
	GetNodeParams() *string
	SetProjectEnv(v string) *RunSmokeTestRequest
	GetProjectEnv() *string
}

type RunSmokeTestRequest struct {
	// The data timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-26 00:00:00
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The parameters related to the node. Set this parameter to a JSON string. A key in the string indicates a parameter, and a value in the string indicates the value of the related parameter.
	//
	// example:
	//
	// bizdate=$bizdate tbods=$tbods
	NodeParams *string `json:"NodeParams,omitempty" xml:"NodeParams,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV. The value PROD indicates the production environment, and the value DEV indicates the development environment. A workspace in basic mode does not have a development environment. For more information, see [Differences between workspaces in basic mode and workspaces in standard mode](https://help.aliyun.com/document_detail/85772.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s RunSmokeTestRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSmokeTestRequest) GoString() string {
	return s.String()
}

func (s *RunSmokeTestRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *RunSmokeTestRequest) GetName() *string {
	return s.Name
}

func (s *RunSmokeTestRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *RunSmokeTestRequest) GetNodeParams() *string {
	return s.NodeParams
}

func (s *RunSmokeTestRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *RunSmokeTestRequest) SetBizdate(v string) *RunSmokeTestRequest {
	s.Bizdate = &v
	return s
}

func (s *RunSmokeTestRequest) SetName(v string) *RunSmokeTestRequest {
	s.Name = &v
	return s
}

func (s *RunSmokeTestRequest) SetNodeId(v int64) *RunSmokeTestRequest {
	s.NodeId = &v
	return s
}

func (s *RunSmokeTestRequest) SetNodeParams(v string) *RunSmokeTestRequest {
	s.NodeParams = &v
	return s
}

func (s *RunSmokeTestRequest) SetProjectEnv(v string) *RunSmokeTestRequest {
	s.ProjectEnv = &v
	return s
}

func (s *RunSmokeTestRequest) Validate() error {
	return dara.Validate(s)
}
