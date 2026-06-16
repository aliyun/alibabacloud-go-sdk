// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskNameByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *TaskNameByUserIdRequest
	GetLang() *string
	SetRegId(v string) *TaskNameByUserIdRequest
	GetRegId() *string
	SetTaskName(v string) *TaskNameByUserIdRequest
	GetTaskName() *string
}

type TaskNameByUserIdRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// The task name.
	//
	// example:
	//
	// Methylation_node_5_3
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s TaskNameByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskNameByUserIdRequest) GoString() string {
	return s.String()
}

func (s *TaskNameByUserIdRequest) GetLang() *string {
	return s.Lang
}

func (s *TaskNameByUserIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *TaskNameByUserIdRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *TaskNameByUserIdRequest) SetLang(v string) *TaskNameByUserIdRequest {
	s.Lang = &v
	return s
}

func (s *TaskNameByUserIdRequest) SetRegId(v string) *TaskNameByUserIdRequest {
	s.RegId = &v
	return s
}

func (s *TaskNameByUserIdRequest) SetTaskName(v string) *TaskNameByUserIdRequest {
	s.TaskName = &v
	return s
}

func (s *TaskNameByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
