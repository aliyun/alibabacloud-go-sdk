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
	// Specifies whether to query the information about ancestor or descendant nodes of the current node. Valid values: input and output.
	//
	// This parameter is required.
	//
	// example:
	//
	// output
	IoType *string `json:"IoType,omitempty" xml:"IoType,omitempty"`
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The runtime environment. Valid values: DEV and PROD.
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
