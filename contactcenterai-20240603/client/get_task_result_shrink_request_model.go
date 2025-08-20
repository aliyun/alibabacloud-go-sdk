// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequiredFieldListShrink(v string) *GetTaskResultShrinkRequest
	GetRequiredFieldListShrink() *string
	SetTaskId(v string) *GetTaskResultShrinkRequest
	GetTaskId() *string
}

type GetTaskResultShrinkRequest struct {
	RequiredFieldListShrink *string `json:"requiredFieldList,omitempty" xml:"requiredFieldList,omitempty"`
	// example:
	//
	// 20240905-********-93E9-5D45-B4EF-045743A34071
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultShrinkRequest) GetRequiredFieldListShrink() *string {
	return s.RequiredFieldListShrink
}

func (s *GetTaskResultShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResultShrinkRequest) SetRequiredFieldListShrink(v string) *GetTaskResultShrinkRequest {
	s.RequiredFieldListShrink = &v
	return s
}

func (s *GetTaskResultShrinkRequest) SetTaskId(v string) *GetTaskResultShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
