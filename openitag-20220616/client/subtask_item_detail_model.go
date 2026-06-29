// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubtaskItemDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v []*SubtaskItemDetailAnnotations) *SubtaskItemDetail
	GetAnnotations() []*SubtaskItemDetailAnnotations
	SetDataSource(v map[string]interface{}) *SubtaskItemDetail
	GetDataSource() map[string]interface{}
	SetItemId(v int64) *SubtaskItemDetail
	GetItemId() *int64
}

type SubtaskItemDetail struct {
	// List of annotation results
	Annotations []*SubtaskItemDetailAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// data source
	//
	// example:
	//
	// None
	DataSource map[string]interface{} `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// Item ID
	//
	// example:
	//
	// 1500***47176994816
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s SubtaskItemDetail) String() string {
	return dara.Prettify(s)
}

func (s SubtaskItemDetail) GoString() string {
	return s.String()
}

func (s *SubtaskItemDetail) GetAnnotations() []*SubtaskItemDetailAnnotations {
	return s.Annotations
}

func (s *SubtaskItemDetail) GetDataSource() map[string]interface{} {
	return s.DataSource
}

func (s *SubtaskItemDetail) GetItemId() *int64 {
	return s.ItemId
}

func (s *SubtaskItemDetail) SetAnnotations(v []*SubtaskItemDetailAnnotations) *SubtaskItemDetail {
	s.Annotations = v
	return s
}

func (s *SubtaskItemDetail) SetDataSource(v map[string]interface{}) *SubtaskItemDetail {
	s.DataSource = v
	return s
}

func (s *SubtaskItemDetail) SetItemId(v int64) *SubtaskItemDetail {
	s.ItemId = &v
	return s
}

func (s *SubtaskItemDetail) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubtaskItemDetailAnnotations struct {
	// Abandon mark
	//
	// example:
	//
	// False
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// Abandonment remark
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
	// Feedback mark
	//
	// example:
	//
	// False
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// Validation feedback
	//
	// example:
	//
	// 验收成功/所有投票结果都未被采纳
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// Failure mark
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
	// Skip mark
	//
	// example:
	//
	// False
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// Status
	//
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// weight
	//
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SubtaskItemDetailAnnotations) String() string {
	return dara.Prettify(s)
}

func (s SubtaskItemDetailAnnotations) GoString() string {
	return s.String()
}

func (s *SubtaskItemDetailAnnotations) GetAbandonFlag() *bool {
	return s.AbandonFlag
}

func (s *SubtaskItemDetailAnnotations) GetAbandonRemark() *string {
	return s.AbandonRemark
}

func (s *SubtaskItemDetailAnnotations) GetDataId() *string {
	return s.DataId
}

func (s *SubtaskItemDetailAnnotations) GetFeedbackFlag() *bool {
	return s.FeedbackFlag
}

func (s *SubtaskItemDetailAnnotations) GetFeedbackRemark() *string {
	return s.FeedbackRemark
}

func (s *SubtaskItemDetailAnnotations) GetFixedFlag() *bool {
	return s.FixedFlag
}

func (s *SubtaskItemDetailAnnotations) GetMine() *int64 {
	return s.Mine
}

func (s *SubtaskItemDetailAnnotations) GetRejectFlag() *bool {
	return s.RejectFlag
}

func (s *SubtaskItemDetailAnnotations) GetState() *string {
	return s.State
}

func (s *SubtaskItemDetailAnnotations) GetWeight() *int64 {
	return s.Weight
}

func (s *SubtaskItemDetailAnnotations) SetAbandonFlag(v bool) *SubtaskItemDetailAnnotations {
	s.AbandonFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetAbandonRemark(v string) *SubtaskItemDetailAnnotations {
	s.AbandonRemark = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetDataId(v string) *SubtaskItemDetailAnnotations {
	s.DataId = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFeedbackFlag(v bool) *SubtaskItemDetailAnnotations {
	s.FeedbackFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFeedbackRemark(v string) *SubtaskItemDetailAnnotations {
	s.FeedbackRemark = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFixedFlag(v bool) *SubtaskItemDetailAnnotations {
	s.FixedFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetMine(v int64) *SubtaskItemDetailAnnotations {
	s.Mine = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetRejectFlag(v bool) *SubtaskItemDetailAnnotations {
	s.RejectFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetState(v string) *SubtaskItemDetailAnnotations {
	s.State = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetWeight(v int64) *SubtaskItemDetailAnnotations {
	s.Weight = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) Validate() error {
	return dara.Validate(s)
}
