// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHyperparameters interface {
	dara.Model
	String() string
	GoString() string
	SetBackupInterval(v int64) *Hyperparameters
	GetBackupInterval() *int64
	SetBatchSize(v int64) *Hyperparameters
	GetBatchSize() *int64
	SetDataLoaderWorkers(v int64) *Hyperparameters
	GetDataLoaderWorkers() *int64
	SetEvaluator(v *CustomParams) *Hyperparameters
	GetEvaluator() *CustomParams
	SetInputSize(v []*int64) *Hyperparameters
	GetInputSize() []*int64
	SetMaxEpoch(v int64) *Hyperparameters
	GetMaxEpoch() *int64
	SetOptimization(v *Optimization) *Hyperparameters
	GetOptimization() *Optimization
	SetSchedule(v *Schedule) *Hyperparameters
	GetSchedule() *Schedule
}

type Hyperparameters struct {
	// example:
	//
	// 1
	BackupInterval *int64 `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	// example:
	//
	// 32
	BatchSize *int64 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// example:
	//
	// 4
	DataLoaderWorkers *int64 `json:"DataLoaderWorkers,omitempty" xml:"DataLoaderWorkers,omitempty"`
	// This parameter is required.
	Evaluator *CustomParams `json:"Evaluator,omitempty" xml:"Evaluator,omitempty"`
	// This parameter is required.
	InputSize []*int64 `json:"InputSize,omitempty" xml:"InputSize,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxEpoch     *int64        `json:"MaxEpoch,omitempty" xml:"MaxEpoch,omitempty"`
	Optimization *Optimization `json:"Optimization,omitempty" xml:"Optimization,omitempty"`
	Schedule     *Schedule     `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s Hyperparameters) String() string {
	return dara.Prettify(s)
}

func (s Hyperparameters) GoString() string {
	return s.String()
}

func (s *Hyperparameters) GetBackupInterval() *int64 {
	return s.BackupInterval
}

func (s *Hyperparameters) GetBatchSize() *int64 {
	return s.BatchSize
}

func (s *Hyperparameters) GetDataLoaderWorkers() *int64 {
	return s.DataLoaderWorkers
}

func (s *Hyperparameters) GetEvaluator() *CustomParams {
	return s.Evaluator
}

func (s *Hyperparameters) GetInputSize() []*int64 {
	return s.InputSize
}

func (s *Hyperparameters) GetMaxEpoch() *int64 {
	return s.MaxEpoch
}

func (s *Hyperparameters) GetOptimization() *Optimization {
	return s.Optimization
}

func (s *Hyperparameters) GetSchedule() *Schedule {
	return s.Schedule
}

func (s *Hyperparameters) SetBackupInterval(v int64) *Hyperparameters {
	s.BackupInterval = &v
	return s
}

func (s *Hyperparameters) SetBatchSize(v int64) *Hyperparameters {
	s.BatchSize = &v
	return s
}

func (s *Hyperparameters) SetDataLoaderWorkers(v int64) *Hyperparameters {
	s.DataLoaderWorkers = &v
	return s
}

func (s *Hyperparameters) SetEvaluator(v *CustomParams) *Hyperparameters {
	s.Evaluator = v
	return s
}

func (s *Hyperparameters) SetInputSize(v []*int64) *Hyperparameters {
	s.InputSize = v
	return s
}

func (s *Hyperparameters) SetMaxEpoch(v int64) *Hyperparameters {
	s.MaxEpoch = &v
	return s
}

func (s *Hyperparameters) SetOptimization(v *Optimization) *Hyperparameters {
	s.Optimization = v
	return s
}

func (s *Hyperparameters) SetSchedule(v *Schedule) *Hyperparameters {
	s.Schedule = v
	return s
}

func (s *Hyperparameters) Validate() error {
	if s.Evaluator != nil {
		if err := s.Evaluator.Validate(); err != nil {
			return err
		}
	}
	if s.Optimization != nil {
		if err := s.Optimization.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
