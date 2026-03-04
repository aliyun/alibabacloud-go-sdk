// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputGroupDetail interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MediaConvertOutputGroupDetail
	GetCode() *string
	SetCreateTime(v string) *MediaConvertOutputGroupDetail
	GetCreateTime() *string
	SetFinishTime(v string) *MediaConvertOutputGroupDetail
	GetFinishTime() *string
	SetMessage(v string) *MediaConvertOutputGroupDetail
	GetMessage() *string
	SetName(v string) *MediaConvertOutputGroupDetail
	GetName() *string
	SetOutputs(v []*MediaConvertOutputDetail) *MediaConvertOutputGroupDetail
	GetOutputs() []*MediaConvertOutputDetail
	SetStatus(v string) *MediaConvertOutputGroupDetail
	GetStatus() *string
	SetTaskId(v string) *MediaConvertOutputGroupDetail
	GetTaskId() *string
}

type MediaConvertOutputGroupDetail struct {
	// The error code for the failed output group task.
	//
	// example:
	//
	// InvalidParameter.ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time the task was created.
	//
	// example:
	//
	// 2025-03-21T01:48:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2025-03-21T01:48:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The reason for a task failure.
	//
	// example:
	//
	// The resource operated "InputFile" is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the output group.
	//
	// example:
	//
	// hls-group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output details.
	Outputs []*MediaConvertOutputDetail `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The status of the output group task.
	//
	// 	- Init: The task is submitted.
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Failed
	//
	// 	- Skipped
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the output group task.
	//
	// example:
	//
	// ******22dad741d086a50325f9******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MediaConvertOutputGroupDetail) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputGroupDetail) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputGroupDetail) GetCode() *string {
	return s.Code
}

func (s *MediaConvertOutputGroupDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MediaConvertOutputGroupDetail) GetFinishTime() *string {
	return s.FinishTime
}

func (s *MediaConvertOutputGroupDetail) GetMessage() *string {
	return s.Message
}

func (s *MediaConvertOutputGroupDetail) GetName() *string {
	return s.Name
}

func (s *MediaConvertOutputGroupDetail) GetOutputs() []*MediaConvertOutputDetail {
	return s.Outputs
}

func (s *MediaConvertOutputGroupDetail) GetStatus() *string {
	return s.Status
}

func (s *MediaConvertOutputGroupDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *MediaConvertOutputGroupDetail) SetCode(v string) *MediaConvertOutputGroupDetail {
	s.Code = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetCreateTime(v string) *MediaConvertOutputGroupDetail {
	s.CreateTime = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetFinishTime(v string) *MediaConvertOutputGroupDetail {
	s.FinishTime = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetMessage(v string) *MediaConvertOutputGroupDetail {
	s.Message = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetName(v string) *MediaConvertOutputGroupDetail {
	s.Name = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetOutputs(v []*MediaConvertOutputDetail) *MediaConvertOutputGroupDetail {
	s.Outputs = v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetStatus(v string) *MediaConvertOutputGroupDetail {
	s.Status = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) SetTaskId(v string) *MediaConvertOutputGroupDetail {
	s.TaskId = &v
	return s
}

func (s *MediaConvertOutputGroupDetail) Validate() error {
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
