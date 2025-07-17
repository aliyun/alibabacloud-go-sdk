// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DeleteTaskRequest
	GetNodeId() *string
	SetTid(v int64) *DeleteTaskRequest
	GetTid() *int64
}

type DeleteTaskRequest struct {
	// The ID of the node you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteTaskRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteTaskRequest) SetNodeId(v string) *DeleteTaskRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteTaskRequest) SetTid(v int64) *DeleteTaskRequest {
	s.Tid = &v
	return s
}

func (s *DeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
