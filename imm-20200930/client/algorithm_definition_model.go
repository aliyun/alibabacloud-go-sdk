// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlgorithmDefinition interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmDefinitionId(v string) *AlgorithmDefinition
	GetAlgorithmDefinitionId() *string
	SetCreateTime(v string) *AlgorithmDefinition
	GetCreateTime() *string
	SetCustomLabels(v []map[string]*string) *AlgorithmDefinition
	GetCustomLabels() []map[string]*string
	SetDescription(v string) *AlgorithmDefinition
	GetDescription() *string
	SetName(v string) *AlgorithmDefinition
	GetName() *string
	SetOwnerId(v string) *AlgorithmDefinition
	GetOwnerId() *string
	SetProjectName(v string) *AlgorithmDefinition
	GetProjectName() *string
	SetTrainingSpecification(v *TrainingSpecification) *AlgorithmDefinition
	GetTrainingSpecification() *TrainingSpecification
	SetUpdateTime(v string) *AlgorithmDefinition
	GetUpdateTime() *string
}

type AlgorithmDefinition struct {
	// example:
	//
	// 8fc6e718-8d19-495f-a510-bcee3c598588
	AlgorithmDefinitionId *string `json:"AlgorithmDefinitionId,omitempty" xml:"AlgorithmDefinitionId,omitempty"`
	// example:
	//
	// 2023-05-31T10:19:40.572325888+08:00
	CreateTime   *string              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomLabels []map[string]*string `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// algoName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// user1
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// traningtest
	ProjectName           *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TrainingSpecification *TrainingSpecification `json:"TrainingSpecification,omitempty" xml:"TrainingSpecification,omitempty"`
	// example:
	//
	// 2023-05-31T10:19:40.572325888+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AlgorithmDefinition) String() string {
	return dara.Prettify(s)
}

func (s AlgorithmDefinition) GoString() string {
	return s.String()
}

func (s *AlgorithmDefinition) GetAlgorithmDefinitionId() *string {
	return s.AlgorithmDefinitionId
}

func (s *AlgorithmDefinition) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AlgorithmDefinition) GetCustomLabels() []map[string]*string {
	return s.CustomLabels
}

func (s *AlgorithmDefinition) GetDescription() *string {
	return s.Description
}

func (s *AlgorithmDefinition) GetName() *string {
	return s.Name
}

func (s *AlgorithmDefinition) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AlgorithmDefinition) GetProjectName() *string {
	return s.ProjectName
}

func (s *AlgorithmDefinition) GetTrainingSpecification() *TrainingSpecification {
	return s.TrainingSpecification
}

func (s *AlgorithmDefinition) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *AlgorithmDefinition) SetAlgorithmDefinitionId(v string) *AlgorithmDefinition {
	s.AlgorithmDefinitionId = &v
	return s
}

func (s *AlgorithmDefinition) SetCreateTime(v string) *AlgorithmDefinition {
	s.CreateTime = &v
	return s
}

func (s *AlgorithmDefinition) SetCustomLabels(v []map[string]*string) *AlgorithmDefinition {
	s.CustomLabels = v
	return s
}

func (s *AlgorithmDefinition) SetDescription(v string) *AlgorithmDefinition {
	s.Description = &v
	return s
}

func (s *AlgorithmDefinition) SetName(v string) *AlgorithmDefinition {
	s.Name = &v
	return s
}

func (s *AlgorithmDefinition) SetOwnerId(v string) *AlgorithmDefinition {
	s.OwnerId = &v
	return s
}

func (s *AlgorithmDefinition) SetProjectName(v string) *AlgorithmDefinition {
	s.ProjectName = &v
	return s
}

func (s *AlgorithmDefinition) SetTrainingSpecification(v *TrainingSpecification) *AlgorithmDefinition {
	s.TrainingSpecification = v
	return s
}

func (s *AlgorithmDefinition) SetUpdateTime(v string) *AlgorithmDefinition {
	s.UpdateTime = &v
	return s
}

func (s *AlgorithmDefinition) Validate() error {
	if s.TrainingSpecification != nil {
		if err := s.TrainingSpecification.Validate(); err != nil {
			return err
		}
	}
	return nil
}
