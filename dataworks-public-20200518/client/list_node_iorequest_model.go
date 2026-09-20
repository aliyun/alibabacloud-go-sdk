// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeIORequest interface {
	dara.Model
	String() string
	GoString() string
	SetIoType(v string) *ListNodeIORequest
	GetIoType() *string
	SetNodeId(v int64) *ListNodeIORequest
	GetNodeId() *int64
	SetProjectEnv(v string) *ListNodeIORequest
	GetProjectEnv() *string
}

type ListNodeIORequest struct {
	// Specifies whether to query upstream or downstream nodes. Valid values: input and output.
	//
	// This parameter is required.
	//
	// example:
	//
	// output
	IoType *string `json:"IoType,omitempty" xml:"IoType,omitempty"`
	// The ID of the node. You can call [ListNodes](https://help.aliyun.com/document_detail/173979.html) to query the NodeId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
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

func (s ListNodeIORequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeIORequest) GoString() string {
	return s.String()
}

func (s *ListNodeIORequest) GetIoType() *string {
	return s.IoType
}

func (s *ListNodeIORequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodeIORequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListNodeIORequest) SetIoType(v string) *ListNodeIORequest {
	s.IoType = &v
	return s
}

func (s *ListNodeIORequest) SetNodeId(v int64) *ListNodeIORequest {
	s.NodeId = &v
	return s
}

func (s *ListNodeIORequest) SetProjectEnv(v string) *ListNodeIORequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListNodeIORequest) Validate() error {
	return dara.Validate(s)
}
