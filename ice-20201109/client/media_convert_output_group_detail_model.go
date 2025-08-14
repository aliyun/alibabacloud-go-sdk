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
	Code       *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime *string                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime *string                     `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message    *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Name       *string                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Outputs    []*MediaConvertOutputDetail `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	Status     *string                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string                     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	return dara.Validate(s)
}
