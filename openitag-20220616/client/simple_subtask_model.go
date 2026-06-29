// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleSubtask interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*SimpleSubtaskItems) *SimpleSubtask
	GetItems() []*SimpleSubtaskItems
	SetStatus(v string) *SimpleSubtask
	GetStatus() *string
	SetSubtaskId(v int64) *SimpleSubtask
	GetSubtaskId() *int64
}

type SimpleSubtask struct {
	// List of items for subtasks.
	Items []*SimpleSubtaskItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Status.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Subtask ID.
	//
	// example:
	//
	// 1500***457270333440
	SubtaskId *int64 `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
}

func (s SimpleSubtask) String() string {
	return dara.Prettify(s)
}

func (s SimpleSubtask) GoString() string {
	return s.String()
}

func (s *SimpleSubtask) GetItems() []*SimpleSubtaskItems {
	return s.Items
}

func (s *SimpleSubtask) GetStatus() *string {
	return s.Status
}

func (s *SimpleSubtask) GetSubtaskId() *int64 {
	return s.SubtaskId
}

func (s *SimpleSubtask) SetItems(v []*SimpleSubtaskItems) *SimpleSubtask {
	s.Items = v
	return s
}

func (s *SimpleSubtask) SetStatus(v string) *SimpleSubtask {
	s.Status = &v
	return s
}

func (s *SimpleSubtask) SetSubtaskId(v int64) *SimpleSubtask {
	s.SubtaskId = &v
	return s
}

func (s *SimpleSubtask) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SimpleSubtaskItems struct {
	// Abandon flag.
	//
	// example:
	//
	// false
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// Abandonment remark.
	//
	// example:
	//
	// 原始数据有问题
	AbandonRemark *string `json:"AbandonRemark,omitempty" xml:"AbandonRemark,omitempty"`
	// Date ID.
	//
	// example:
	//
	// 175296157992643****
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Feedback flag.
	//
	// example:
	//
	// false
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// Validation feedback.
	//
	// example:
	//
	// 验收完成
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// Failed mark.
	//
	// example:
	//
	// false
	FixedFlag *bool `json:"FixedFlag,omitempty" xml:"FixedFlag,omitempty"`
	// Data item ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 15116***94667356160
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// Assigned to me:
	//
	// - 0: Not assigned to me.
	//
	// - 1: Assigned to me.
	//
	// example:
	//
	// 0
	Mine *int64 `json:"Mine,omitempty" xml:"Mine,omitempty"`
	// Skip flag.
	//
	// example:
	//
	// false
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// Status:
	//
	// - INIT: Initial status.
	//
	// - TOPUBLISH: Pending publish.
	//
	// - CREATED: Created.
	//
	// - HANDLING: In progress.
	//
	// - VOTING: Voting in progress.
	//
	// - FINISHED: Completed.
	//
	// - FAIL: Failed.
	//
	// - EXPIRE: Timeout.
	//
	// - DISCARDED: Passively abandoned.
	//
	// - DISABLE: Actively abandoned.
	//
	// - LOCKED: Edit Lock.
	//
	// - OFFLINE: Unpublished.
	//
	// - MERGING: Merging results.
	//
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Weight.
	//
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SimpleSubtaskItems) String() string {
	return dara.Prettify(s)
}

func (s SimpleSubtaskItems) GoString() string {
	return s.String()
}

func (s *SimpleSubtaskItems) GetAbandonFlag() *bool {
	return s.AbandonFlag
}

func (s *SimpleSubtaskItems) GetAbandonRemark() *string {
	return s.AbandonRemark
}

func (s *SimpleSubtaskItems) GetDataId() *string {
	return s.DataId
}

func (s *SimpleSubtaskItems) GetFeedbackFlag() *bool {
	return s.FeedbackFlag
}

func (s *SimpleSubtaskItems) GetFeedbackRemark() *string {
	return s.FeedbackRemark
}

func (s *SimpleSubtaskItems) GetFixedFlag() *bool {
	return s.FixedFlag
}

func (s *SimpleSubtaskItems) GetItemId() *int64 {
	return s.ItemId
}

func (s *SimpleSubtaskItems) GetMine() *int64 {
	return s.Mine
}

func (s *SimpleSubtaskItems) GetRejectFlag() *bool {
	return s.RejectFlag
}

func (s *SimpleSubtaskItems) GetState() *string {
	return s.State
}

func (s *SimpleSubtaskItems) GetWeight() *int64 {
	return s.Weight
}

func (s *SimpleSubtaskItems) SetAbandonFlag(v bool) *SimpleSubtaskItems {
	s.AbandonFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetAbandonRemark(v string) *SimpleSubtaskItems {
	s.AbandonRemark = &v
	return s
}

func (s *SimpleSubtaskItems) SetDataId(v string) *SimpleSubtaskItems {
	s.DataId = &v
	return s
}

func (s *SimpleSubtaskItems) SetFeedbackFlag(v bool) *SimpleSubtaskItems {
	s.FeedbackFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetFeedbackRemark(v string) *SimpleSubtaskItems {
	s.FeedbackRemark = &v
	return s
}

func (s *SimpleSubtaskItems) SetFixedFlag(v bool) *SimpleSubtaskItems {
	s.FixedFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetItemId(v int64) *SimpleSubtaskItems {
	s.ItemId = &v
	return s
}

func (s *SimpleSubtaskItems) SetMine(v int64) *SimpleSubtaskItems {
	s.Mine = &v
	return s
}

func (s *SimpleSubtaskItems) SetRejectFlag(v bool) *SimpleSubtaskItems {
	s.RejectFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetState(v string) *SimpleSubtaskItems {
	s.State = &v
	return s
}

func (s *SimpleSubtaskItems) SetWeight(v int64) *SimpleSubtaskItems {
	s.Weight = &v
	return s
}

func (s *SimpleSubtaskItems) Validate() error {
	return dara.Validate(s)
}
