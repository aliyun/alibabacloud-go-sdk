// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobElasticSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAIMasterDockerImage(v string) *JobElasticSpec
	GetAIMasterDockerImage() *string
	SetAIMasterType(v string) *JobElasticSpec
	GetAIMasterType() *string
	SetEDPMaxParallelism(v int32) *JobElasticSpec
	GetEDPMaxParallelism() *int32
	SetEDPMinParallelism(v int32) *JobElasticSpec
	GetEDPMinParallelism() *int32
	SetElasticStrategy(v string) *JobElasticSpec
	GetElasticStrategy() *string
	SetEnableAIMaster(v bool) *JobElasticSpec
	GetEnableAIMaster() *bool
	SetEnableEDP(v bool) *JobElasticSpec
	GetEnableEDP() *bool
	SetEnableElasticTraining(v bool) *JobElasticSpec
	GetEnableElasticTraining() *bool
	SetEnablePsJobElasticPS(v bool) *JobElasticSpec
	GetEnablePsJobElasticPS() *bool
	SetEnablePsJobElasticWorker(v bool) *JobElasticSpec
	GetEnablePsJobElasticWorker() *bool
	SetEnablePsResourceEstimate(v bool) *JobElasticSpec
	GetEnablePsResourceEstimate() *bool
	SetMaxParallelism(v int32) *JobElasticSpec
	GetMaxParallelism() *int32
	SetMinParallelism(v int32) *JobElasticSpec
	GetMinParallelism() *int32
	SetPSMaxParallelism(v int32) *JobElasticSpec
	GetPSMaxParallelism() *int32
	SetPSMinParallelism(v int32) *JobElasticSpec
	GetPSMinParallelism() *int32
}

type JobElasticSpec struct {
	// The AIMaster Docker image.
	AIMasterDockerImage *string `json:"AIMasterDockerImage,omitempty" xml:"AIMasterDockerImage,omitempty"`
	// The resource specifications for the AIMaster role.
	//
	// example:
	//
	// ecs.c6.large
	AIMasterType *string `json:"AIMasterType,omitempty" xml:"AIMasterType,omitempty"`
	// The maximum EDP parallelism.
	//
	// example:
	//
	// 16
	EDPMaxParallelism *int32 `json:"EDPMaxParallelism,omitempty" xml:"EDPMaxParallelism,omitempty"`
	// The minimum EDP parallelism.
	//
	// example:
	//
	// 8
	EDPMinParallelism *int32 `json:"EDPMinParallelism,omitempty" xml:"EDPMinParallelism,omitempty"`
	// The elastic training strategy.
	//
	// example:
	//
	// chief
	ElasticStrategy *string `json:"ElasticStrategy,omitempty" xml:"ElasticStrategy,omitempty"`
	// Specifies whether to enable AIMaster.
	EnableAIMaster *bool `json:"EnableAIMaster,omitempty" xml:"EnableAIMaster,omitempty"`
	// Specifies whether to enable EDP.
	//
	// example:
	//
	// true
	EnableEDP *bool `json:"EnableEDP,omitempty" xml:"EnableEDP,omitempty"`
	// Specifies whether to enable elastic training.
	//
	// example:
	//
	// true
	EnableElasticTraining *bool `json:"EnableElasticTraining,omitempty" xml:"EnableElasticTraining,omitempty"`
	// Specifies whether to enable the elastic PS feature.
	//
	// example:
	//
	// true
	EnablePsJobElasticPS *bool `json:"EnablePsJobElasticPS,omitempty" xml:"EnablePsJobElasticPS,omitempty"`
	// Specifies whether to enable elastic workers for PS jobs.
	EnablePsJobElasticWorker *bool `json:"EnablePsJobElasticWorker,omitempty" xml:"EnablePsJobElasticWorker,omitempty"`
	// Specifies whether to enable PS resource estimation.
	//
	// example:
	//
	// true
	EnablePsResourceEstimate *bool `json:"EnablePsResourceEstimate,omitempty" xml:"EnablePsResourceEstimate,omitempty"`
	// The maximum parallelism.
	//
	// example:
	//
	// 8
	MaxParallelism *int32 `json:"MaxParallelism,omitempty" xml:"MaxParallelism,omitempty"`
	// The minimum parallelism.
	//
	// example:
	//
	// 1
	MinParallelism *int32 `json:"MinParallelism,omitempty" xml:"MinParallelism,omitempty"`
	// The maximum PS parallelism.
	//
	// example:
	//
	// 10
	PSMaxParallelism *int32 `json:"PSMaxParallelism,omitempty" xml:"PSMaxParallelism,omitempty"`
	// The minimum PS parallelism.
	//
	// example:
	//
	// 4
	PSMinParallelism *int32 `json:"PSMinParallelism,omitempty" xml:"PSMinParallelism,omitempty"`
}

func (s JobElasticSpec) String() string {
	return dara.Prettify(s)
}

func (s JobElasticSpec) GoString() string {
	return s.String()
}

func (s *JobElasticSpec) GetAIMasterDockerImage() *string {
	return s.AIMasterDockerImage
}

func (s *JobElasticSpec) GetAIMasterType() *string {
	return s.AIMasterType
}

func (s *JobElasticSpec) GetEDPMaxParallelism() *int32 {
	return s.EDPMaxParallelism
}

func (s *JobElasticSpec) GetEDPMinParallelism() *int32 {
	return s.EDPMinParallelism
}

func (s *JobElasticSpec) GetElasticStrategy() *string {
	return s.ElasticStrategy
}

func (s *JobElasticSpec) GetEnableAIMaster() *bool {
	return s.EnableAIMaster
}

func (s *JobElasticSpec) GetEnableEDP() *bool {
	return s.EnableEDP
}

func (s *JobElasticSpec) GetEnableElasticTraining() *bool {
	return s.EnableElasticTraining
}

func (s *JobElasticSpec) GetEnablePsJobElasticPS() *bool {
	return s.EnablePsJobElasticPS
}

func (s *JobElasticSpec) GetEnablePsJobElasticWorker() *bool {
	return s.EnablePsJobElasticWorker
}

func (s *JobElasticSpec) GetEnablePsResourceEstimate() *bool {
	return s.EnablePsResourceEstimate
}

func (s *JobElasticSpec) GetMaxParallelism() *int32 {
	return s.MaxParallelism
}

func (s *JobElasticSpec) GetMinParallelism() *int32 {
	return s.MinParallelism
}

func (s *JobElasticSpec) GetPSMaxParallelism() *int32 {
	return s.PSMaxParallelism
}

func (s *JobElasticSpec) GetPSMinParallelism() *int32 {
	return s.PSMinParallelism
}

func (s *JobElasticSpec) SetAIMasterDockerImage(v string) *JobElasticSpec {
	s.AIMasterDockerImage = &v
	return s
}

func (s *JobElasticSpec) SetAIMasterType(v string) *JobElasticSpec {
	s.AIMasterType = &v
	return s
}

func (s *JobElasticSpec) SetEDPMaxParallelism(v int32) *JobElasticSpec {
	s.EDPMaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetEDPMinParallelism(v int32) *JobElasticSpec {
	s.EDPMinParallelism = &v
	return s
}

func (s *JobElasticSpec) SetElasticStrategy(v string) *JobElasticSpec {
	s.ElasticStrategy = &v
	return s
}

func (s *JobElasticSpec) SetEnableAIMaster(v bool) *JobElasticSpec {
	s.EnableAIMaster = &v
	return s
}

func (s *JobElasticSpec) SetEnableEDP(v bool) *JobElasticSpec {
	s.EnableEDP = &v
	return s
}

func (s *JobElasticSpec) SetEnableElasticTraining(v bool) *JobElasticSpec {
	s.EnableElasticTraining = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsJobElasticPS(v bool) *JobElasticSpec {
	s.EnablePsJobElasticPS = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsJobElasticWorker(v bool) *JobElasticSpec {
	s.EnablePsJobElasticWorker = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsResourceEstimate(v bool) *JobElasticSpec {
	s.EnablePsResourceEstimate = &v
	return s
}

func (s *JobElasticSpec) SetMaxParallelism(v int32) *JobElasticSpec {
	s.MaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetMinParallelism(v int32) *JobElasticSpec {
	s.MinParallelism = &v
	return s
}

func (s *JobElasticSpec) SetPSMaxParallelism(v int32) *JobElasticSpec {
	s.PSMaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetPSMinParallelism(v int32) *JobElasticSpec {
	s.PSMinParallelism = &v
	return s
}

func (s *JobElasticSpec) Validate() error {
	return dara.Validate(s)
}
