// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowsByPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagIdListShrink(v string) *ListTaskFlowsByPageShrinkRequest
	GetDagIdListShrink() *string
	SetPageIndex(v int32) *ListTaskFlowsByPageShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListTaskFlowsByPageShrinkRequest
	GetPageSize() *int32
	SetScenarioId(v int64) *ListTaskFlowsByPageShrinkRequest
	GetScenarioId() *int64
	SetSearchKey(v string) *ListTaskFlowsByPageShrinkRequest
	GetSearchKey() *string
	SetTid(v int64) *ListTaskFlowsByPageShrinkRequest
	GetTid() *int64
}

type ListTaskFlowsByPageShrinkRequest struct {
	// Filter condition, task flow ID list.
	DagIdListShrink *string `json:"DagIdList,omitempty" xml:"DagIdList,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filter condition, application scenario ID.
	//
	// example:
	//
	// 12***
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The keyword that is used to search for task flow names.
	//
	// example:
	//
	// Test node
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTaskFlowsByPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageShrinkRequest) GetDagIdListShrink() *string {
	return s.DagIdListShrink
}

func (s *ListTaskFlowsByPageShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListTaskFlowsByPageShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskFlowsByPageShrinkRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *ListTaskFlowsByPageShrinkRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListTaskFlowsByPageShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowsByPageShrinkRequest) SetDagIdListShrink(v string) *ListTaskFlowsByPageShrinkRequest {
	s.DagIdListShrink = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) SetPageIndex(v int32) *ListTaskFlowsByPageShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) SetPageSize(v int32) *ListTaskFlowsByPageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) SetScenarioId(v int64) *ListTaskFlowsByPageShrinkRequest {
	s.ScenarioId = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) SetSearchKey(v string) *ListTaskFlowsByPageShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) SetTid(v int64) *ListTaskFlowsByPageShrinkRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowsByPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
