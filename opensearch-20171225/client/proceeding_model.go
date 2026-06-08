// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProceeding interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v map[string]interface{}) *Proceeding
	GetDetail() map[string]interface{}
	SetProgress(v float32) *Proceeding
	GetProgress() *float32
	SetStatus(v string) *Proceeding
	GetStatus() *string
	SetSubTasks(v *Proceeding) *Proceeding
	GetSubTasks() *Proceeding
	SetType(v string) *Proceeding
	GetType() *string
}

type Proceeding struct {
	Detail   map[string]interface{} `json:"detail,omitempty" xml:"detail,omitempty"`
	Progress *float32               `json:"progress,omitempty" xml:"progress,omitempty"`
	Status   *string                `json:"status,omitempty" xml:"status,omitempty"`
	SubTasks *Proceeding            `json:"subTasks,omitempty" xml:"subTasks,omitempty"`
	Type     *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Proceeding) String() string {
	return dara.Prettify(s)
}

func (s Proceeding) GoString() string {
	return s.String()
}

func (s *Proceeding) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *Proceeding) GetProgress() *float32 {
	return s.Progress
}

func (s *Proceeding) GetStatus() *string {
	return s.Status
}

func (s *Proceeding) GetSubTasks() *Proceeding {
	return s.SubTasks
}

func (s *Proceeding) GetType() *string {
	return s.Type
}

func (s *Proceeding) SetDetail(v map[string]interface{}) *Proceeding {
	s.Detail = v
	return s
}

func (s *Proceeding) SetProgress(v float32) *Proceeding {
	s.Progress = &v
	return s
}

func (s *Proceeding) SetStatus(v string) *Proceeding {
	s.Status = &v
	return s
}

func (s *Proceeding) SetSubTasks(v *Proceeding) *Proceeding {
	s.SubTasks = v
	return s
}

func (s *Proceeding) SetType(v string) *Proceeding {
	s.Type = &v
	return s
}

func (s *Proceeding) Validate() error {
	if s.SubTasks != nil {
		if err := s.SubTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}
