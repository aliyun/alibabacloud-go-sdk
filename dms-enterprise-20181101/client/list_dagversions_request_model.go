// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDAGVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListDAGVersionsRequest
	GetDagId() *int64
	SetPageIndex(v int32) *ListDAGVersionsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDAGVersionsRequest
	GetPageSize() *int32
	SetTid(v int64) *ListDAGVersionsRequest
	GetTid() *int64
}

type ListDAGVersionsRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDAGVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDAGVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDAGVersionsRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListDAGVersionsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDAGVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDAGVersionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDAGVersionsRequest) SetDagId(v int64) *ListDAGVersionsRequest {
	s.DagId = &v
	return s
}

func (s *ListDAGVersionsRequest) SetPageIndex(v int32) *ListDAGVersionsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDAGVersionsRequest) SetPageSize(v int32) *ListDAGVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDAGVersionsRequest) SetTid(v int64) *ListDAGVersionsRequest {
	s.Tid = &v
	return s
}

func (s *ListDAGVersionsRequest) Validate() error {
	return dara.Validate(s)
}
