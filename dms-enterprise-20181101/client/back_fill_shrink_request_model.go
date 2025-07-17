// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackFillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *BackFillShrinkRequest
	GetAsc() *bool
	SetBackFillDate(v string) *BackFillShrinkRequest
	GetBackFillDate() *string
	SetBackFillDateBegin(v string) *BackFillShrinkRequest
	GetBackFillDateBegin() *string
	SetBackFillDateEnd(v string) *BackFillShrinkRequest
	GetBackFillDateEnd() *string
	SetDagId(v int64) *BackFillShrinkRequest
	GetDagId() *int64
	SetFilterNodeIdsShrink(v string) *BackFillShrinkRequest
	GetFilterNodeIdsShrink() *string
	SetHistoryDagId(v int64) *BackFillShrinkRequest
	GetHistoryDagId() *int64
	SetInterval(v int32) *BackFillShrinkRequest
	GetInterval() *int32
	SetIsTriggerSubTree(v bool) *BackFillShrinkRequest
	GetIsTriggerSubTree() *bool
	SetStartNodeIdsShrink(v string) *BackFillShrinkRequest
	GetStartNodeIdsShrink() *string
	SetTid(v int64) *BackFillShrinkRequest
	GetTid() *int64
}

type BackFillShrinkRequest struct {
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
	FilterNodeIdsShrink *string `json:"FilterNodeIds,omitempty" xml:"FilterNodeIds,omitempty"`
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
	StartNodeIdsShrink *string `json:"StartNodeIds,omitempty" xml:"StartNodeIds,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s BackFillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BackFillShrinkRequest) GoString() string {
	return s.String()
}

func (s *BackFillShrinkRequest) GetAsc() *bool {
	return s.Asc
}

func (s *BackFillShrinkRequest) GetBackFillDate() *string {
	return s.BackFillDate
}

func (s *BackFillShrinkRequest) GetBackFillDateBegin() *string {
	return s.BackFillDateBegin
}

func (s *BackFillShrinkRequest) GetBackFillDateEnd() *string {
	return s.BackFillDateEnd
}

func (s *BackFillShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *BackFillShrinkRequest) GetFilterNodeIdsShrink() *string {
	return s.FilterNodeIdsShrink
}

func (s *BackFillShrinkRequest) GetHistoryDagId() *int64 {
	return s.HistoryDagId
}

func (s *BackFillShrinkRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *BackFillShrinkRequest) GetIsTriggerSubTree() *bool {
	return s.IsTriggerSubTree
}

func (s *BackFillShrinkRequest) GetStartNodeIdsShrink() *string {
	return s.StartNodeIdsShrink
}

func (s *BackFillShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BackFillShrinkRequest) SetAsc(v bool) *BackFillShrinkRequest {
	s.Asc = &v
	return s
}

func (s *BackFillShrinkRequest) SetBackFillDate(v string) *BackFillShrinkRequest {
	s.BackFillDate = &v
	return s
}

func (s *BackFillShrinkRequest) SetBackFillDateBegin(v string) *BackFillShrinkRequest {
	s.BackFillDateBegin = &v
	return s
}

func (s *BackFillShrinkRequest) SetBackFillDateEnd(v string) *BackFillShrinkRequest {
	s.BackFillDateEnd = &v
	return s
}

func (s *BackFillShrinkRequest) SetDagId(v int64) *BackFillShrinkRequest {
	s.DagId = &v
	return s
}

func (s *BackFillShrinkRequest) SetFilterNodeIdsShrink(v string) *BackFillShrinkRequest {
	s.FilterNodeIdsShrink = &v
	return s
}

func (s *BackFillShrinkRequest) SetHistoryDagId(v int64) *BackFillShrinkRequest {
	s.HistoryDagId = &v
	return s
}

func (s *BackFillShrinkRequest) SetInterval(v int32) *BackFillShrinkRequest {
	s.Interval = &v
	return s
}

func (s *BackFillShrinkRequest) SetIsTriggerSubTree(v bool) *BackFillShrinkRequest {
	s.IsTriggerSubTree = &v
	return s
}

func (s *BackFillShrinkRequest) SetStartNodeIdsShrink(v string) *BackFillShrinkRequest {
	s.StartNodeIdsShrink = &v
	return s
}

func (s *BackFillShrinkRequest) SetTid(v int64) *BackFillShrinkRequest {
	s.Tid = &v
	return s
}

func (s *BackFillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
