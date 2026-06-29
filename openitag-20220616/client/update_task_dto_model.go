// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskDTO interface {
	dara.Model
	String() string
	GoString() string
	SetExif(v map[string]*string) *UpdateTaskDTO
	GetExif() map[string]*string
	SetRemark(v string) *UpdateTaskDTO
	GetRemark() *string
	SetTags(v []*string) *UpdateTaskDTO
	GetTags() []*string
	SetTaskName(v string) *UpdateTaskDTO
	GetTaskName() *string
}

type UpdateTaskDTO struct {
	// Extended field
	Exif map[string]*string `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Remark information
	//
	// example:
	//
	// demo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// List of labels
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Task Name
	//
	// example:
	//
	// demo
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateTaskDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskDTO) GoString() string {
	return s.String()
}

func (s *UpdateTaskDTO) GetExif() map[string]*string {
	return s.Exif
}

func (s *UpdateTaskDTO) GetRemark() *string {
	return s.Remark
}

func (s *UpdateTaskDTO) GetTags() []*string {
	return s.Tags
}

func (s *UpdateTaskDTO) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateTaskDTO) SetExif(v map[string]*string) *UpdateTaskDTO {
	s.Exif = v
	return s
}

func (s *UpdateTaskDTO) SetRemark(v string) *UpdateTaskDTO {
	s.Remark = &v
	return s
}

func (s *UpdateTaskDTO) SetTags(v []*string) *UpdateTaskDTO {
	s.Tags = v
	return s
}

func (s *UpdateTaskDTO) SetTaskName(v string) *UpdateTaskDTO {
	s.TaskName = &v
	return s
}

func (s *UpdateTaskDTO) Validate() error {
	return dara.Validate(s)
}
