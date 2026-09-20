// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInputOrOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIoType(v string) *ListNodeInputOrOutputRequest
	GetIoType() *string
	SetNodeId(v int64) *ListNodeInputOrOutputRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *ListNodeInputOrOutputRequest
	GetProjectEnv() *string
}

type ListNodeInputOrOutputRequest struct {
	// Specifies whether to query upstream or downstream nodes. Valid values:
	//
	// - input: upstream nodes.
	//
	// - output: downstream nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// output
	IoType *string `json:"IoType,omitempty" xml:"IoType,omitempty"`
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12314567
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment in which the node runs. Valid values: DEV (development environment) and PROD (production environment).
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListNodeInputOrOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInputOrOutputRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInputOrOutputRequest) GetIoType() *string {
	return s.IoType
}

func (s *ListNodeInputOrOutputRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodeInputOrOutputRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListNodeInputOrOutputRequest) SetIoType(v string) *ListNodeInputOrOutputRequest {
	s.IoType = &v
	return s
}

func (s *ListNodeInputOrOutputRequest) SetNodeId(v int64) *ListNodeInputOrOutputRequest {
	s.NodeId = &v
	return s
}

func (s *ListNodeInputOrOutputRequest) SetProjectEnv(v string) *ListNodeInputOrOutputRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListNodeInputOrOutputRequest) Validate() error {
	return dara.Validate(s)
}
