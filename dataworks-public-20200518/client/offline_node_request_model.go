// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *OfflineNodeRequest
	GetNodeId() *int64
	SetProjectId(v int64) *OfflineNodeRequest
	GetProjectId() *int64
}

type OfflineNodeRequest struct {
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to obtain the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s OfflineNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineNodeRequest) GoString() string {
	return s.String()
}

func (s *OfflineNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *OfflineNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *OfflineNodeRequest) SetNodeId(v int64) *OfflineNodeRequest {
	s.NodeId = &v
	return s
}

func (s *OfflineNodeRequest) SetProjectId(v int64) *OfflineNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *OfflineNodeRequest) Validate() error {
	return dara.Validate(s)
}
