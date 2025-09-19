// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotRequest
	GetCurrentPage() *int32
	SetHoneypotIds(v []*string) *ListHoneypotRequest
	GetHoneypotIds() []*string
	SetHoneypotName(v string) *ListHoneypotRequest
	GetHoneypotName() *string
	SetNodeId(v string) *ListHoneypotRequest
	GetNodeId() *string
	SetNodeName(v string) *ListHoneypotRequest
	GetNodeName() *string
	SetPageSize(v int32) *ListHoneypotRequest
	GetPageSize() *int32
}

type ListHoneypotRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IDs of the honeypots.
	HoneypotIds []*string `json:"HoneypotIds,omitempty" xml:"HoneypotIds,omitempty" type:"Repeated"`
	// The name of the honeypot.
	//
	// example:
	//
	// mx-rouyi
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// honeypot_master
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotRequest) GetHoneypotIds() []*string {
	return s.HoneypotIds
}

func (s *ListHoneypotRequest) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *ListHoneypotRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListHoneypotRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *ListHoneypotRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotRequest) SetCurrentPage(v int32) *ListHoneypotRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotRequest) SetHoneypotIds(v []*string) *ListHoneypotRequest {
	s.HoneypotIds = v
	return s
}

func (s *ListHoneypotRequest) SetHoneypotName(v string) *ListHoneypotRequest {
	s.HoneypotName = &v
	return s
}

func (s *ListHoneypotRequest) SetNodeId(v string) *ListHoneypotRequest {
	s.NodeId = &v
	return s
}

func (s *ListHoneypotRequest) SetNodeName(v string) *ListHoneypotRequest {
	s.NodeName = &v
	return s
}

func (s *ListHoneypotRequest) SetPageSize(v int32) *ListHoneypotRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
