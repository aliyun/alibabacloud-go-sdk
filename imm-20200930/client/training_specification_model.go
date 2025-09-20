// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainingSpecification interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *TrainingSpecification
	GetDatasetName() *string
	SetEndpoint(v string) *TrainingSpecification
	GetEndpoint() *string
	SetModelSpecification(v *ModelSpecification) *TrainingSpecification
	GetModelSpecification() *ModelSpecification
	SetRuntime(v *Runtime) *TrainingSpecification
	GetRuntime() *Runtime
	SetSourceURI(v string) *TrainingSpecification
	GetSourceURI() *string
	SetTargetURI(v string) *TrainingSpecification
	GetTargetURI() *string
	SetTransforms(v []*CustomParams) *TrainingSpecification
	GetTransforms() []*CustomParams
	SetValidationSourceURI(v string) *TrainingSpecification
	GetValidationSourceURI() *string
	SetValidationSplit(v float32) *TrainingSpecification
	GetValidationSplit() *float32
}

type TrainingSpecification struct {
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// This parameter is required.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// This parameter is required.
	ModelSpecification *ModelSpecification `json:"ModelSpecification,omitempty" xml:"ModelSpecification,omitempty"`
	// This parameter is required.
	Runtime *Runtime `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-alg-dataset-bj/cifar10/test_index.json
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-alg-dataset-bj/model_training_test/
	TargetURI  *string         `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	Transforms []*CustomParams `json:"Transforms,omitempty" xml:"Transforms,omitempty" type:"Repeated"`
	// example:
	//
	// oss://imm-alg-dataset-bj/cifar10/test_index.json
	ValidationSourceURI *string `json:"ValidationSourceURI,omitempty" xml:"ValidationSourceURI,omitempty"`
	// example:
	//
	// 0.95
	ValidationSplit *float32 `json:"ValidationSplit,omitempty" xml:"ValidationSplit,omitempty"`
}

func (s TrainingSpecification) String() string {
	return dara.Prettify(s)
}

func (s TrainingSpecification) GoString() string {
	return s.String()
}

func (s *TrainingSpecification) GetDatasetName() *string {
	return s.DatasetName
}

func (s *TrainingSpecification) GetEndpoint() *string {
	return s.Endpoint
}

func (s *TrainingSpecification) GetModelSpecification() *ModelSpecification {
	return s.ModelSpecification
}

func (s *TrainingSpecification) GetRuntime() *Runtime {
	return s.Runtime
}

func (s *TrainingSpecification) GetSourceURI() *string {
	return s.SourceURI
}

func (s *TrainingSpecification) GetTargetURI() *string {
	return s.TargetURI
}

func (s *TrainingSpecification) GetTransforms() []*CustomParams {
	return s.Transforms
}

func (s *TrainingSpecification) GetValidationSourceURI() *string {
	return s.ValidationSourceURI
}

func (s *TrainingSpecification) GetValidationSplit() *float32 {
	return s.ValidationSplit
}

func (s *TrainingSpecification) SetDatasetName(v string) *TrainingSpecification {
	s.DatasetName = &v
	return s
}

func (s *TrainingSpecification) SetEndpoint(v string) *TrainingSpecification {
	s.Endpoint = &v
	return s
}

func (s *TrainingSpecification) SetModelSpecification(v *ModelSpecification) *TrainingSpecification {
	s.ModelSpecification = v
	return s
}

func (s *TrainingSpecification) SetRuntime(v *Runtime) *TrainingSpecification {
	s.Runtime = v
	return s
}

func (s *TrainingSpecification) SetSourceURI(v string) *TrainingSpecification {
	s.SourceURI = &v
	return s
}

func (s *TrainingSpecification) SetTargetURI(v string) *TrainingSpecification {
	s.TargetURI = &v
	return s
}

func (s *TrainingSpecification) SetTransforms(v []*CustomParams) *TrainingSpecification {
	s.Transforms = v
	return s
}

func (s *TrainingSpecification) SetValidationSourceURI(v string) *TrainingSpecification {
	s.ValidationSourceURI = &v
	return s
}

func (s *TrainingSpecification) SetValidationSplit(v float32) *TrainingSpecification {
	s.ValidationSplit = &v
	return s
}

func (s *TrainingSpecification) Validate() error {
	return dara.Validate(s)
}
