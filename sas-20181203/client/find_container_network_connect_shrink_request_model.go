// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindContainerNetworkConnectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaType(v string) *FindContainerNetworkConnectShrinkRequest
	GetCriteriaType() *string
	SetCurrentPage(v int64) *FindContainerNetworkConnectShrinkRequest
	GetCurrentPage() *int64
	SetDstNodeShrink(v string) *FindContainerNetworkConnectShrinkRequest
	GetDstNodeShrink() *string
	SetEndTime(v int64) *FindContainerNetworkConnectShrinkRequest
	GetEndTime() *int64
	SetPageSize(v int64) *FindContainerNetworkConnectShrinkRequest
	GetPageSize() *int64
	SetSrcNodeShrink(v string) *FindContainerNetworkConnectShrinkRequest
	GetSrcNodeShrink() *string
	SetStartTime(v int64) *FindContainerNetworkConnectShrinkRequest
	GetStartTime() *int64
}

type FindContainerNetworkConnectShrinkRequest struct {
	// The type of the information that you want to query. Valid values:
	//
	// 	- **EDGE**: connection information
	//
	// example:
	//
	// EDGE
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The information about the destination node.
	DstNodeShrink *string `json:"DstNode,omitempty" xml:"DstNode,omitempty"`
	// The end time of the network connection.
	//
	// example:
	//
	// 1649260799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the source node.
	SrcNodeShrink *string `json:"SrcNode,omitempty" xml:"SrcNode,omitempty"`
	// The start time of the network connection.
	//
	// example:
	//
	// 1666886400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s FindContainerNetworkConnectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FindContainerNetworkConnectShrinkRequest) GoString() string {
	return s.String()
}

func (s *FindContainerNetworkConnectShrinkRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *FindContainerNetworkConnectShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *FindContainerNetworkConnectShrinkRequest) GetDstNodeShrink() *string {
	return s.DstNodeShrink
}

func (s *FindContainerNetworkConnectShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *FindContainerNetworkConnectShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *FindContainerNetworkConnectShrinkRequest) GetSrcNodeShrink() *string {
	return s.SrcNodeShrink
}

func (s *FindContainerNetworkConnectShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *FindContainerNetworkConnectShrinkRequest) SetCriteriaType(v string) *FindContainerNetworkConnectShrinkRequest {
	s.CriteriaType = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetCurrentPage(v int64) *FindContainerNetworkConnectShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetDstNodeShrink(v string) *FindContainerNetworkConnectShrinkRequest {
	s.DstNodeShrink = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetEndTime(v int64) *FindContainerNetworkConnectShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetPageSize(v int64) *FindContainerNetworkConnectShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetSrcNodeShrink(v string) *FindContainerNetworkConnectShrinkRequest {
	s.SrcNodeShrink = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) SetStartTime(v int64) *FindContainerNetworkConnectShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *FindContainerNetworkConnectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
