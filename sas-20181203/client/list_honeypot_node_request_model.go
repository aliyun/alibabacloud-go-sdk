// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotNodeRequest
	GetCurrentPage() *int32
	SetNodeId(v string) *ListHoneypotNodeRequest
	GetNodeId() *string
	SetNodeName(v string) *ListHoneypotNodeRequest
	GetNodeName() *string
	SetPageSize(v int32) *ListHoneypotNodeRequest
	GetPageSize() *int32
}

type ListHoneypotNodeRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// 7d110ca6-05ee-4149-8042-13ad1a41fd****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// cyct_cnymu
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotNodeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotNodeRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListHoneypotNodeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotNodeRequest) SetCurrentPage(v int32) *ListHoneypotNodeRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotNodeRequest) SetNodeId(v string) *ListHoneypotNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotNodeRequest) SetNodeName(v string) *ListHoneypotNodeRequest {
	s.NodeName = &v
	return s
}

func (s *ListHoneypotNodeRequest) SetPageSize(v int32) *ListHoneypotNodeRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
