// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgInferenceJobInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetCreateUserId(v string) *Personalizedtxt2imgInferenceJobInfoDTO
	GetCreateUserId() *string
	SetId(v string) *Personalizedtxt2imgInferenceJobInfoDTO
	GetId() *string
	SetJobStatus(v string) *Personalizedtxt2imgInferenceJobInfoDTO
	GetJobStatus() *string
	SetModelId(v string) *Personalizedtxt2imgInferenceJobInfoDTO
	GetModelId() *string
	SetResultImageUrl(v []*string) *Personalizedtxt2imgInferenceJobInfoDTO
	GetResultImageUrl() []*string
}

type Personalizedtxt2imgInferenceJobInfoDTO struct {
	// example:
	//
	// 123456
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 123456
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 123456
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 123456
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) GetId() *string {
	return s.Id
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) GetJobStatus() *string {
	return s.JobStatus
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) GetResultImageUrl() []*string {
	return s.ResultImageUrl
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetModelId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetResultImageUrl(v []*string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ResultImageUrl = v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) Validate() error {
	return dara.Validate(s)
}
