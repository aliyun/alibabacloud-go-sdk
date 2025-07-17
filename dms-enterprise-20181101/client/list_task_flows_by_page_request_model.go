// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowsByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagIdList(v []*int64) *ListTaskFlowsByPageRequest
	GetDagIdList() []*int64
	SetPageIndex(v int32) *ListTaskFlowsByPageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListTaskFlowsByPageRequest
	GetPageSize() *int32
	SetScenarioId(v int64) *ListTaskFlowsByPageRequest
	GetScenarioId() *int64
	SetSearchKey(v string) *ListTaskFlowsByPageRequest
	GetSearchKey() *string
	SetTid(v int64) *ListTaskFlowsByPageRequest
	GetTid() *int64
}

type ListTaskFlowsByPageRequest struct {
	// Filter condition, task flow ID list.
	DagIdList []*int64 `json:"DagIdList,omitempty" xml:"DagIdList,omitempty" type:"Repeated"`
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

func (s ListTaskFlowsByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageRequest) GetDagIdList() []*int64 {
	return s.DagIdList
}

func (s *ListTaskFlowsByPageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListTaskFlowsByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskFlowsByPageRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *ListTaskFlowsByPageRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListTaskFlowsByPageRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowsByPageRequest) SetDagIdList(v []*int64) *ListTaskFlowsByPageRequest {
	s.DagIdList = v
	return s
}

func (s *ListTaskFlowsByPageRequest) SetPageIndex(v int32) *ListTaskFlowsByPageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTaskFlowsByPageRequest) SetPageSize(v int32) *ListTaskFlowsByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskFlowsByPageRequest) SetScenarioId(v int64) *ListTaskFlowsByPageRequest {
	s.ScenarioId = &v
	return s
}

func (s *ListTaskFlowsByPageRequest) SetSearchKey(v string) *ListTaskFlowsByPageRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTaskFlowsByPageRequest) SetTid(v int64) *ListTaskFlowsByPageRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowsByPageRequest) Validate() error {
	return dara.Validate(s)
}
