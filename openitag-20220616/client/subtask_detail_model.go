// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubtaskDetail interface {
	dara.Model
	String() string
	GoString() string
	SetCanDiscard(v bool) *SubtaskDetail
	GetCanDiscard() *bool
	SetCanReassign(v bool) *SubtaskDetail
	GetCanReassign() *bool
	SetCanRelease(v bool) *SubtaskDetail
	GetCanRelease() *bool
	SetCurrentWorkNode(v string) *SubtaskDetail
	GetCurrentWorkNode() *string
	SetExtConfigs(v string) *SubtaskDetail
	GetExtConfigs() *string
	SetItems(v []*SubtaskDetailItems) *SubtaskDetail
	GetItems() []*SubtaskDetailItems
	SetStatus(v string) *SubtaskDetail
	GetStatus() *string
	SetSubtaskId(v string) *SubtaskDetail
	GetSubtaskId() *string
	SetTaskId(v string) *SubtaskDetail
	GetTaskId() *string
	SetWeight(v int64) *SubtaskDetail
	GetWeight() *int64
	SetWorkNodeState(v string) *SubtaskDetail
	GetWorkNodeState() *string
	SetWorkforce(v []*Workforce) *SubtaskDetail
	GetWorkforce() []*Workforce
}

type SubtaskDetail struct {
	// is discardable
	//
	// example:
	//
	// false
	CanDiscard *bool `json:"CanDiscard,omitempty" xml:"CanDiscard,omitempty"`
	// Can assign
	//
	// example:
	//
	// false
	CanReassign *bool `json:"CanReassign,omitempty" xml:"CanReassign,omitempty"`
	// is releasable
	//
	// example:
	//
	// false
	CanRelease *bool `json:"CanRelease,omitempty" xml:"CanRelease,omitempty"`
	// current File Type
	//
	// example:
	//
	// MARK
	CurrentWorkNode *string `json:"CurrentWorkNode,omitempty" xml:"CurrentWorkNode,omitempty"`
	// extra parameters
	//
	// if can be null:
	// true
	//
	// example:
	//
	// null
	ExtConfigs *string `json:"ExtConfigs,omitempty" xml:"ExtConfigs,omitempty"`
	// List of items in the sub-job
	Items []*SubtaskDetailItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// status
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Subtask ID
	//
	// example:
	//
	// 1500***457270333440
	SubtaskId *string `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
	// parent job ID of the sub-job
	//
	// example:
	//
	// 1511***994667356160
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Job weight
	//
	// example:
	//
	// 631548
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// Current edge zone status
	//
	// example:
	//
	// FINISHED
	WorkNodeState *string `json:"WorkNodeState,omitempty" xml:"WorkNodeState,omitempty"`
	// list of annotators assigned to the sub-job
	Workforce []*Workforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s SubtaskDetail) String() string {
	return dara.Prettify(s)
}

func (s SubtaskDetail) GoString() string {
	return s.String()
}

func (s *SubtaskDetail) GetCanDiscard() *bool {
	return s.CanDiscard
}

func (s *SubtaskDetail) GetCanReassign() *bool {
	return s.CanReassign
}

func (s *SubtaskDetail) GetCanRelease() *bool {
	return s.CanRelease
}

func (s *SubtaskDetail) GetCurrentWorkNode() *string {
	return s.CurrentWorkNode
}

func (s *SubtaskDetail) GetExtConfigs() *string {
	return s.ExtConfigs
}

func (s *SubtaskDetail) GetItems() []*SubtaskDetailItems {
	return s.Items
}

func (s *SubtaskDetail) GetStatus() *string {
	return s.Status
}

func (s *SubtaskDetail) GetSubtaskId() *string {
	return s.SubtaskId
}

func (s *SubtaskDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *SubtaskDetail) GetWeight() *int64 {
	return s.Weight
}

func (s *SubtaskDetail) GetWorkNodeState() *string {
	return s.WorkNodeState
}

func (s *SubtaskDetail) GetWorkforce() []*Workforce {
	return s.Workforce
}

func (s *SubtaskDetail) SetCanDiscard(v bool) *SubtaskDetail {
	s.CanDiscard = &v
	return s
}

func (s *SubtaskDetail) SetCanReassign(v bool) *SubtaskDetail {
	s.CanReassign = &v
	return s
}

func (s *SubtaskDetail) SetCanRelease(v bool) *SubtaskDetail {
	s.CanRelease = &v
	return s
}

func (s *SubtaskDetail) SetCurrentWorkNode(v string) *SubtaskDetail {
	s.CurrentWorkNode = &v
	return s
}

func (s *SubtaskDetail) SetExtConfigs(v string) *SubtaskDetail {
	s.ExtConfigs = &v
	return s
}

func (s *SubtaskDetail) SetItems(v []*SubtaskDetailItems) *SubtaskDetail {
	s.Items = v
	return s
}

func (s *SubtaskDetail) SetStatus(v string) *SubtaskDetail {
	s.Status = &v
	return s
}

func (s *SubtaskDetail) SetSubtaskId(v string) *SubtaskDetail {
	s.SubtaskId = &v
	return s
}

func (s *SubtaskDetail) SetTaskId(v string) *SubtaskDetail {
	s.TaskId = &v
	return s
}

func (s *SubtaskDetail) SetWeight(v int64) *SubtaskDetail {
	s.Weight = &v
	return s
}

func (s *SubtaskDetail) SetWorkNodeState(v string) *SubtaskDetail {
	s.WorkNodeState = &v
	return s
}

func (s *SubtaskDetail) SetWorkforce(v []*Workforce) *SubtaskDetail {
	s.Workforce = v
	return s
}

func (s *SubtaskDetail) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Workforce != nil {
		for _, item := range s.Workforce {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubtaskDetailItems struct {
	// Abandon mark
	//
	// example:
	//
	// False
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// discard reason
	//
	// example:
	//
	// 原始数据错误
	AbandonRemark *string `json:"AbandonRemark,omitempty" xml:"AbandonRemark,omitempty"`
	// Date ID
	//
	// example:
	//
	// 1957578084
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// feedback mark
	//
	// example:
	//
	// False
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// Validate feedback
	//
	// example:
	//
	// 标注内容合格
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// Failed mark
	//
	// example:
	//
	// False
	FixedFlag *bool `json:"FixedFlag,omitempty" xml:"FixedFlag,omitempty"`
	// Is assigned to me
	//
	// example:
	//
	// 0
	Mine *int64 `json:"Mine,omitempty" xml:"Mine,omitempty"`
	// skip mark
	//
	// example:
	//
	// False
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// status
	//
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Weight
	//
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SubtaskDetailItems) String() string {
	return dara.Prettify(s)
}

func (s SubtaskDetailItems) GoString() string {
	return s.String()
}

func (s *SubtaskDetailItems) GetAbandonFlag() *bool {
	return s.AbandonFlag
}

func (s *SubtaskDetailItems) GetAbandonRemark() *string {
	return s.AbandonRemark
}

func (s *SubtaskDetailItems) GetDataId() *string {
	return s.DataId
}

func (s *SubtaskDetailItems) GetFeedbackFlag() *bool {
	return s.FeedbackFlag
}

func (s *SubtaskDetailItems) GetFeedbackRemark() *string {
	return s.FeedbackRemark
}

func (s *SubtaskDetailItems) GetFixedFlag() *bool {
	return s.FixedFlag
}

func (s *SubtaskDetailItems) GetMine() *int64 {
	return s.Mine
}

func (s *SubtaskDetailItems) GetRejectFlag() *bool {
	return s.RejectFlag
}

func (s *SubtaskDetailItems) GetState() *string {
	return s.State
}

func (s *SubtaskDetailItems) GetWeight() *int64 {
	return s.Weight
}

func (s *SubtaskDetailItems) SetAbandonFlag(v bool) *SubtaskDetailItems {
	s.AbandonFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetAbandonRemark(v string) *SubtaskDetailItems {
	s.AbandonRemark = &v
	return s
}

func (s *SubtaskDetailItems) SetDataId(v string) *SubtaskDetailItems {
	s.DataId = &v
	return s
}

func (s *SubtaskDetailItems) SetFeedbackFlag(v bool) *SubtaskDetailItems {
	s.FeedbackFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetFeedbackRemark(v string) *SubtaskDetailItems {
	s.FeedbackRemark = &v
	return s
}

func (s *SubtaskDetailItems) SetFixedFlag(v bool) *SubtaskDetailItems {
	s.FixedFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetMine(v int64) *SubtaskDetailItems {
	s.Mine = &v
	return s
}

func (s *SubtaskDetailItems) SetRejectFlag(v bool) *SubtaskDetailItems {
	s.RejectFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetState(v string) *SubtaskDetailItems {
	s.State = &v
	return s
}

func (s *SubtaskDetailItems) SetWeight(v int64) *SubtaskDetailItems {
	s.Weight = &v
	return s
}

func (s *SubtaskDetailItems) Validate() error {
	return dara.Validate(s)
}
