// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIllustrationTask interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v string) *IllustrationTask
	GetGmtCreate() *string
	SetGmtModified(v string) *IllustrationTask
	GetGmtModified() *string
	SetIllustrationIds(v []*int64) *IllustrationTask
	GetIllustrationIds() []*int64
	SetIllustrationTaskId(v int64) *IllustrationTask
	GetIllustrationTaskId() *int64
	SetTaskStatus(v string) *IllustrationTask
	GetTaskStatus() *string
	SetTextId(v int64) *IllustrationTask
	GetTextId() *int64
}

type IllustrationTask struct {
	GmtCreate          *string  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IllustrationIds    []*int64 `json:"illustrationIds,omitempty" xml:"illustrationIds,omitempty" type:"Repeated"`
	IllustrationTaskId *int64   `json:"illustrationTaskId,omitempty" xml:"illustrationTaskId,omitempty"`
	// example:
	//
	// Success
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	TextId     *int64  `json:"textId,omitempty" xml:"textId,omitempty"`
}

func (s IllustrationTask) String() string {
	return dara.Prettify(s)
}

func (s IllustrationTask) GoString() string {
	return s.String()
}

func (s *IllustrationTask) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *IllustrationTask) GetGmtModified() *string {
	return s.GmtModified
}

func (s *IllustrationTask) GetIllustrationIds() []*int64 {
	return s.IllustrationIds
}

func (s *IllustrationTask) GetIllustrationTaskId() *int64 {
	return s.IllustrationTaskId
}

func (s *IllustrationTask) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *IllustrationTask) GetTextId() *int64 {
	return s.TextId
}

func (s *IllustrationTask) SetGmtCreate(v string) *IllustrationTask {
	s.GmtCreate = &v
	return s
}

func (s *IllustrationTask) SetGmtModified(v string) *IllustrationTask {
	s.GmtModified = &v
	return s
}

func (s *IllustrationTask) SetIllustrationIds(v []*int64) *IllustrationTask {
	s.IllustrationIds = v
	return s
}

func (s *IllustrationTask) SetIllustrationTaskId(v int64) *IllustrationTask {
	s.IllustrationTaskId = &v
	return s
}

func (s *IllustrationTask) SetTaskStatus(v string) *IllustrationTask {
	s.TaskStatus = &v
	return s
}

func (s *IllustrationTask) SetTextId(v int64) *IllustrationTask {
	s.TextId = &v
	return s
}

func (s *IllustrationTask) Validate() error {
	return dara.Validate(s)
}
