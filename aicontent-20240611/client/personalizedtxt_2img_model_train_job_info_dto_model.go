// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgModelTrainJobInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetCreateTime() *string
	SetCreateUserId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetCreateUserId() *string
	SetImageUrl(v []*string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetImageUrl() []*string
	SetInferenceJobList(v []*Personalizedtxt2imgInferenceJobInfoDTO) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetInferenceJobList() []*Personalizedtxt2imgInferenceJobInfoDTO
	SetJobStatus(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetJobStatus() *string
	SetName(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetName() *string
	SetObjectType(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetObjectType() *string
	SetId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO
	GetId() *string
}

type Personalizedtxt2imgModelTrainJobInfoDTO struct {
	CreateTime       *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId     *string                                   `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	ImageUrl         []*string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
	InferenceJobList []*Personalizedtxt2imgInferenceJobInfoDTO `json:"InferenceJobList,omitempty" xml:"InferenceJobList,omitempty" type:"Repeated"`
	JobStatus        *string                                   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Name             *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType       *string                                   `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetInferenceJobList() []*Personalizedtxt2imgInferenceJobInfoDTO {
	return s.InferenceJobList
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetName() *string {
	return s.Name
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetObjectType() *string {
	return s.ObjectType
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateTime(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetImageUrl(v []*string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetInferenceJobList(v []*Personalizedtxt2imgInferenceJobInfoDTO) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetName(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetObjectType(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) Validate() error {
	if s.InferenceJobList != nil {
		for _, item := range s.InferenceJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
