// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIncidentTracingSubNodesCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVertexIdAndTypeList(v [][]*string) *QueryIncidentTracingSubNodesCountRequest
	GetVertexIdAndTypeList() [][]*string
}

type QueryIncidentTracingSubNodesCountRequest struct {
	// The key-value pairs that consist of node IDs and node types. A key-value pair is an array.
	VertexIdAndTypeList [][]*string `json:"VertexIdAndTypeList,omitempty" xml:"VertexIdAndTypeList,omitempty" type:"Repeated"`
}

func (s QueryIncidentTracingSubNodesCountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIncidentTracingSubNodesCountRequest) GoString() string {
	return s.String()
}

func (s *QueryIncidentTracingSubNodesCountRequest) GetVertexIdAndTypeList() [][]*string {
	return s.VertexIdAndTypeList
}

func (s *QueryIncidentTracingSubNodesCountRequest) SetVertexIdAndTypeList(v [][]*string) *QueryIncidentTracingSubNodesCountRequest {
	s.VertexIdAndTypeList = v
	return s
}

func (s *QueryIncidentTracingSubNodesCountRequest) Validate() error {
	return dara.Validate(s)
}
