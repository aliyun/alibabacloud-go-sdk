// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetNodeContent(v string) *UpdateTaskContentV2Request
	GetNodeContent() *string
	SetNodeId(v string) *UpdateTaskContentV2Request
	GetNodeId() *string
}

type UpdateTaskContentV2Request struct {
	// example:
	//
	// { "dbId":12****, "sql":"select 	- from test_table",   "dbType":"lindorm_sql"  }
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// example:
	//
	// 449***
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s UpdateTaskContentV2Request) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentV2Request) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentV2Request) GetNodeContent() *string {
	return s.NodeContent
}

func (s *UpdateTaskContentV2Request) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskContentV2Request) SetNodeContent(v string) *UpdateTaskContentV2Request {
	s.NodeContent = &v
	return s
}

func (s *UpdateTaskContentV2Request) SetNodeId(v string) *UpdateTaskContentV2Request {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskContentV2Request) Validate() error {
	return dara.Validate(s)
}
