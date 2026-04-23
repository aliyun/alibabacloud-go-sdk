// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeJobInfoVO interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *KnowledgeJobInfoVO
	GetCreateTime() *string
	SetCreator(v string) *KnowledgeJobInfoVO
	GetCreator() *string
	SetDescription(v string) *KnowledgeJobInfoVO
	GetDescription() *string
	SetEndTime(v string) *KnowledgeJobInfoVO
	GetEndTime() *string
	SetJobId(v int32) *KnowledgeJobInfoVO
	GetJobId() *int32
	SetKnowledgeCnt(v int32) *KnowledgeJobInfoVO
	GetKnowledgeCnt() *int32
	SetProgress(v int32) *KnowledgeJobInfoVO
	GetProgress() *int32
	SetShowJobId(v string) *KnowledgeJobInfoVO
	GetShowJobId() *string
	SetStatus(v string) *KnowledgeJobInfoVO
	GetStatus() *string
	SetSupplement(v string) *KnowledgeJobInfoVO
	GetSupplement() *string
	SetTableCnt(v int32) *KnowledgeJobInfoVO
	GetTableCnt() *int32
}

type KnowledgeJobInfoVO struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId        *int32  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	KnowledgeCnt *int32  `json:"KnowledgeCnt,omitempty" xml:"KnowledgeCnt,omitempty"`
	Progress     *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ShowJobId    *string `json:"ShowJobId,omitempty" xml:"ShowJobId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Supplement   *string `json:"Supplement,omitempty" xml:"Supplement,omitempty"`
	TableCnt     *int32  `json:"TableCnt,omitempty" xml:"TableCnt,omitempty"`
}

func (s KnowledgeJobInfoVO) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeJobInfoVO) GoString() string {
	return s.String()
}

func (s *KnowledgeJobInfoVO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *KnowledgeJobInfoVO) GetCreator() *string {
	return s.Creator
}

func (s *KnowledgeJobInfoVO) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeJobInfoVO) GetEndTime() *string {
	return s.EndTime
}

func (s *KnowledgeJobInfoVO) GetJobId() *int32 {
	return s.JobId
}

func (s *KnowledgeJobInfoVO) GetKnowledgeCnt() *int32 {
	return s.KnowledgeCnt
}

func (s *KnowledgeJobInfoVO) GetProgress() *int32 {
	return s.Progress
}

func (s *KnowledgeJobInfoVO) GetShowJobId() *string {
	return s.ShowJobId
}

func (s *KnowledgeJobInfoVO) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeJobInfoVO) GetSupplement() *string {
	return s.Supplement
}

func (s *KnowledgeJobInfoVO) GetTableCnt() *int32 {
	return s.TableCnt
}

func (s *KnowledgeJobInfoVO) SetCreateTime(v string) *KnowledgeJobInfoVO {
	s.CreateTime = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetCreator(v string) *KnowledgeJobInfoVO {
	s.Creator = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetDescription(v string) *KnowledgeJobInfoVO {
	s.Description = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetEndTime(v string) *KnowledgeJobInfoVO {
	s.EndTime = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetJobId(v int32) *KnowledgeJobInfoVO {
	s.JobId = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetKnowledgeCnt(v int32) *KnowledgeJobInfoVO {
	s.KnowledgeCnt = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetProgress(v int32) *KnowledgeJobInfoVO {
	s.Progress = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetShowJobId(v string) *KnowledgeJobInfoVO {
	s.ShowJobId = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetStatus(v string) *KnowledgeJobInfoVO {
	s.Status = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetSupplement(v string) *KnowledgeJobInfoVO {
	s.Supplement = &v
	return s
}

func (s *KnowledgeJobInfoVO) SetTableCnt(v int32) *KnowledgeJobInfoVO {
	s.TableCnt = &v
	return s
}

func (s *KnowledgeJobInfoVO) Validate() error {
	return dara.Validate(s)
}
