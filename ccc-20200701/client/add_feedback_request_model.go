// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeedback(v string) *AddFeedbackRequest
	GetFeedback() *string
	SetInstanceId(v string) *AddFeedbackRequest
	GetInstanceId() *string
	SetRating(v int32) *AddFeedbackRequest
	GetRating() *int32
	SetTaskId(v string) *AddFeedbackRequest
	GetTaskId() *string
	SetTaskName(v string) *AddFeedbackRequest
	GetTaskName() *string
}

type AddFeedbackRequest struct {
	// The feedback provided by returning users.
	//
	// example:
	//
	// {"问题描述":"客户询问沙发生产周期并尝试加快", "客服方案":"订单确认，建议联系在线客服", "完成度判断":"否"}
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Rating:
	//
	// - thumbsDown: Thumbs down.
	//
	// - thumbsUp: Thumbs up.
	//
	// example:
	//
	// thumbsUp
	Rating *int32 `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// AI task ID.
	//
	// example:
	//
	// f780ade8-****-458b-b067-63077946a570
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task type.
	//
	// - Abstract:fields (Field extraction)
	//
	// - Abstract:keywords (Hot keywords)
	//
	// - Abstract:title_summary (Summary)
	//
	// example:
	//
	// Abstract:fields
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s AddFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFeedbackRequest) GoString() string {
	return s.String()
}

func (s *AddFeedbackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *AddFeedbackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddFeedbackRequest) GetRating() *int32 {
	return s.Rating
}

func (s *AddFeedbackRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AddFeedbackRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *AddFeedbackRequest) SetFeedback(v string) *AddFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *AddFeedbackRequest) SetInstanceId(v string) *AddFeedbackRequest {
	s.InstanceId = &v
	return s
}

func (s *AddFeedbackRequest) SetRating(v int32) *AddFeedbackRequest {
	s.Rating = &v
	return s
}

func (s *AddFeedbackRequest) SetTaskId(v string) *AddFeedbackRequest {
	s.TaskId = &v
	return s
}

func (s *AddFeedbackRequest) SetTaskName(v string) *AddFeedbackRequest {
	s.TaskName = &v
	return s
}

func (s *AddFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
