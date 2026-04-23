// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeSubJobInfoVO interface {
	dara.Model
	String() string
	GoString() string
	SetErrMessage(v string) *KnowledgeSubJobInfoVO
	GetErrMessage() *string
	SetJobType(v string) *KnowledgeSubJobInfoVO
	GetJobType() *string
	SetName(v string) *KnowledgeSubJobInfoVO
	GetName() *string
	SetProgress(v int32) *KnowledgeSubJobInfoVO
	GetProgress() *int32
	SetStatus(v string) *KnowledgeSubJobInfoVO
	GetStatus() *string
	SetSubJobId(v int32) *KnowledgeSubJobInfoVO
	GetSubJobId() *int32
}

type KnowledgeSubJobInfoVO struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	JobType    *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress   *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubJobId   *int32  `json:"SubJobId,omitempty" xml:"SubJobId,omitempty"`
}

func (s KnowledgeSubJobInfoVO) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeSubJobInfoVO) GoString() string {
	return s.String()
}

func (s *KnowledgeSubJobInfoVO) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *KnowledgeSubJobInfoVO) GetJobType() *string {
	return s.JobType
}

func (s *KnowledgeSubJobInfoVO) GetName() *string {
	return s.Name
}

func (s *KnowledgeSubJobInfoVO) GetProgress() *int32 {
	return s.Progress
}

func (s *KnowledgeSubJobInfoVO) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeSubJobInfoVO) GetSubJobId() *int32 {
	return s.SubJobId
}

func (s *KnowledgeSubJobInfoVO) SetErrMessage(v string) *KnowledgeSubJobInfoVO {
	s.ErrMessage = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) SetJobType(v string) *KnowledgeSubJobInfoVO {
	s.JobType = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) SetName(v string) *KnowledgeSubJobInfoVO {
	s.Name = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) SetProgress(v int32) *KnowledgeSubJobInfoVO {
	s.Progress = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) SetStatus(v string) *KnowledgeSubJobInfoVO {
	s.Status = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) SetSubJobId(v int32) *KnowledgeSubJobInfoVO {
	s.SubJobId = &v
	return s
}

func (s *KnowledgeSubJobInfoVO) Validate() error {
	return dara.Validate(s)
}
