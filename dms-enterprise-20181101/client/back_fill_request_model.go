// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackFillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *BackFillRequest
	GetAsc() *bool
	SetBackFillDate(v string) *BackFillRequest
	GetBackFillDate() *string
	SetBackFillDateBegin(v string) *BackFillRequest
	GetBackFillDateBegin() *string
	SetBackFillDateEnd(v string) *BackFillRequest
	GetBackFillDateEnd() *string
	SetDagId(v int64) *BackFillRequest
	GetDagId() *int64
	SetFilterNodeIds(v []*int64) *BackFillRequest
	GetFilterNodeIds() []*int64
	SetHistoryDagId(v int64) *BackFillRequest
	GetHistoryDagId() *int64
	SetInterval(v int32) *BackFillRequest
	GetInterval() *int32
	SetIsTriggerSubTree(v bool) *BackFillRequest
	GetIsTriggerSubTree() *bool
	SetStartNodeIds(v []*int64) *BackFillRequest
	GetStartNodeIds() []*int64
	SetTid(v int64) *BackFillRequest
	GetTid() *int64
}

type BackFillRequest struct {
	// The running sequence of task flows for data backfill. Valid values:
	//
	// 	- **0**: reverse chronological order.
	//
	// 	- **1**: chronological order. This is the default value.
	//
	// example:
	//
	// 0
	Asc *bool `json:"Asc,omitempty" xml:"Asc,omitempty"`
	// The date for the data to be backfilled. This parameter is required if you specify a date for data backfill.
	//
	// example:
	//
	// 2022-01-14
	BackFillDate *string `json:"BackFillDate,omitempty" xml:"BackFillDate,omitempty"`
	// The start date of the date range for the data to be backfilled. This parameter is required if you specify a date range for data backfill.
	//
	// example:
	//
	// 2022-01-14
	BackFillDateBegin *string `json:"BackFillDateBegin,omitempty" xml:"BackFillDateBegin,omitempty"`
	// The end date of the date range for the data to be backfilled. This parameter is required if you specify a date range for data backfill.
	//
	// example:
	//
	// 2022-09-29
	BackFillDateEnd *string `json:"BackFillDateEnd,omitempty" xml:"BackFillDateEnd,omitempty"`
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// Filter condition, which specifies the list of node IDs in the task flow that do not need to supplement data.
	FilterNodeIds []*int64 `json:"FilterNodeIds,omitempty" xml:"FilterNodeIds,omitempty" type:"Repeated"`
	// The ID of the historical task flow.
	//
	// example:
	//
	// 16***
	HistoryDagId *int64 `json:"HistoryDagId,omitempty" xml:"HistoryDagId,omitempty"`
	// The interval at which data backfill is performed. Unit: hours. Minimum value: 1. Default value: 24.
	//
	// example:
	//
	// 24
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Specifies whether to run descendant nodes. Default value: true.
	//
	// example:
	//
	// true
	IsTriggerSubTree *bool `json:"IsTriggerSubTree,omitempty" xml:"IsTriggerSubTree,omitempty"`
	// The number of nodes for which you want to backfill data.
	StartNodeIds []*int64 `json:"StartNodeIds,omitempty" xml:"StartNodeIds,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s BackFillRequest) String() string {
	return dara.Prettify(s)
}

func (s BackFillRequest) GoString() string {
	return s.String()
}

func (s *BackFillRequest) GetAsc() *bool {
	return s.Asc
}

func (s *BackFillRequest) GetBackFillDate() *string {
	return s.BackFillDate
}

func (s *BackFillRequest) GetBackFillDateBegin() *string {
	return s.BackFillDateBegin
}

func (s *BackFillRequest) GetBackFillDateEnd() *string {
	return s.BackFillDateEnd
}

func (s *BackFillRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *BackFillRequest) GetFilterNodeIds() []*int64 {
	return s.FilterNodeIds
}

func (s *BackFillRequest) GetHistoryDagId() *int64 {
	return s.HistoryDagId
}

func (s *BackFillRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *BackFillRequest) GetIsTriggerSubTree() *bool {
	return s.IsTriggerSubTree
}

func (s *BackFillRequest) GetStartNodeIds() []*int64 {
	return s.StartNodeIds
}

func (s *BackFillRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BackFillRequest) SetAsc(v bool) *BackFillRequest {
	s.Asc = &v
	return s
}

func (s *BackFillRequest) SetBackFillDate(v string) *BackFillRequest {
	s.BackFillDate = &v
	return s
}

func (s *BackFillRequest) SetBackFillDateBegin(v string) *BackFillRequest {
	s.BackFillDateBegin = &v
	return s
}

func (s *BackFillRequest) SetBackFillDateEnd(v string) *BackFillRequest {
	s.BackFillDateEnd = &v
	return s
}

func (s *BackFillRequest) SetDagId(v int64) *BackFillRequest {
	s.DagId = &v
	return s
}

func (s *BackFillRequest) SetFilterNodeIds(v []*int64) *BackFillRequest {
	s.FilterNodeIds = v
	return s
}

func (s *BackFillRequest) SetHistoryDagId(v int64) *BackFillRequest {
	s.HistoryDagId = &v
	return s
}

func (s *BackFillRequest) SetInterval(v int32) *BackFillRequest {
	s.Interval = &v
	return s
}

func (s *BackFillRequest) SetIsTriggerSubTree(v bool) *BackFillRequest {
	s.IsTriggerSubTree = &v
	return s
}

func (s *BackFillRequest) SetStartNodeIds(v []*int64) *BackFillRequest {
	s.StartNodeIds = v
	return s
}

func (s *BackFillRequest) SetTid(v int64) *BackFillRequest {
	s.Tid = &v
	return s
}

func (s *BackFillRequest) Validate() error {
	return dara.Validate(s)
}
