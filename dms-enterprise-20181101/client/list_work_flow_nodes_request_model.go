// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkFlowNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchName(v string) *ListWorkFlowNodesRequest
	GetSearchName() *string
	SetTid(v int64) *ListWorkFlowNodesRequest
	GetTid() *int64
}

type ListWorkFlowNodesRequest struct {
	// The name that is used to search for approval nodes.
	//
	// example:
	//
	// admin
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListWorkFlowNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkFlowNodesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkFlowNodesRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListWorkFlowNodesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListWorkFlowNodesRequest) SetSearchName(v string) *ListWorkFlowNodesRequest {
	s.SearchName = &v
	return s
}

func (s *ListWorkFlowNodesRequest) SetTid(v int64) *ListWorkFlowNodesRequest {
	s.Tid = &v
	return s
}

func (s *ListWorkFlowNodesRequest) Validate() error {
	return dara.Validate(s)
}
