// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TaskSnapshot
	GetBizId() *string
	SetCommiter(v int64) *TaskSnapshot
	GetCommiter() *int64
	SetGmtCreated(v string) *TaskSnapshot
	GetGmtCreated() *string
	SetItem(v *Task) *TaskSnapshot
	GetItem() *Task
	SetMessage(v string) *TaskSnapshot
	GetMessage() *string
	SetTaskBizId(v string) *TaskSnapshot
	GetTaskBizId() *string
	SetVersion(v string) *TaskSnapshot
	GetVersion() *string
}

type TaskSnapshot struct {
	BizId      *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Commiter   *int64  `json:"commiter,omitempty" xml:"commiter,omitempty"`
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	Item       *Task   `json:"item,omitempty" xml:"item,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	TaskBizId  *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TaskSnapshot) String() string {
	return dara.Prettify(s)
}

func (s TaskSnapshot) GoString() string {
	return s.String()
}

func (s *TaskSnapshot) GetBizId() *string {
	return s.BizId
}

func (s *TaskSnapshot) GetCommiter() *int64 {
	return s.Commiter
}

func (s *TaskSnapshot) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *TaskSnapshot) GetItem() *Task {
	return s.Item
}

func (s *TaskSnapshot) GetMessage() *string {
	return s.Message
}

func (s *TaskSnapshot) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *TaskSnapshot) GetVersion() *string {
	return s.Version
}

func (s *TaskSnapshot) SetBizId(v string) *TaskSnapshot {
	s.BizId = &v
	return s
}

func (s *TaskSnapshot) SetCommiter(v int64) *TaskSnapshot {
	s.Commiter = &v
	return s
}

func (s *TaskSnapshot) SetGmtCreated(v string) *TaskSnapshot {
	s.GmtCreated = &v
	return s
}

func (s *TaskSnapshot) SetItem(v *Task) *TaskSnapshot {
	s.Item = v
	return s
}

func (s *TaskSnapshot) SetMessage(v string) *TaskSnapshot {
	s.Message = &v
	return s
}

func (s *TaskSnapshot) SetTaskBizId(v string) *TaskSnapshot {
	s.TaskBizId = &v
	return s
}

func (s *TaskSnapshot) SetVersion(v string) *TaskSnapshot {
	s.Version = &v
	return s
}

func (s *TaskSnapshot) Validate() error {
	return dara.Validate(s)
}
