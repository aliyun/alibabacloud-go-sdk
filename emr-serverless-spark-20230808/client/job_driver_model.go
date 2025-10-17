// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobDriver interface {
	dara.Model
	String() string
	GoString() string
	SetSparkSubmit(v *JobDriverSparkSubmit) *JobDriver
	GetSparkSubmit() *JobDriverSparkSubmit
}

type JobDriver struct {
	SparkSubmit *JobDriverSparkSubmit `json:"sparkSubmit,omitempty" xml:"sparkSubmit,omitempty" type:"Struct"`
}

func (s JobDriver) String() string {
	return dara.Prettify(s)
}

func (s JobDriver) GoString() string {
	return s.String()
}

func (s *JobDriver) GetSparkSubmit() *JobDriverSparkSubmit {
	return s.SparkSubmit
}

func (s *JobDriver) SetSparkSubmit(v *JobDriverSparkSubmit) *JobDriver {
	s.SparkSubmit = v
	return s
}

func (s *JobDriver) Validate() error {
	if s.SparkSubmit != nil {
		if err := s.SparkSubmit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JobDriverSparkSubmit struct {
	EntryPoint            *string   `json:"entryPoint,omitempty" xml:"entryPoint,omitempty"`
	EntryPointArguments   []*string `json:"entryPointArguments,omitempty" xml:"entryPointArguments,omitempty" type:"Repeated"`
	SparkSubmitParameters *string   `json:"sparkSubmitParameters,omitempty" xml:"sparkSubmitParameters,omitempty"`
}

func (s JobDriverSparkSubmit) String() string {
	return dara.Prettify(s)
}

func (s JobDriverSparkSubmit) GoString() string {
	return s.String()
}

func (s *JobDriverSparkSubmit) GetEntryPoint() *string {
	return s.EntryPoint
}

func (s *JobDriverSparkSubmit) GetEntryPointArguments() []*string {
	return s.EntryPointArguments
}

func (s *JobDriverSparkSubmit) GetSparkSubmitParameters() *string {
	return s.SparkSubmitParameters
}

func (s *JobDriverSparkSubmit) SetEntryPoint(v string) *JobDriverSparkSubmit {
	s.EntryPoint = &v
	return s
}

func (s *JobDriverSparkSubmit) SetEntryPointArguments(v []*string) *JobDriverSparkSubmit {
	s.EntryPointArguments = v
	return s
}

func (s *JobDriverSparkSubmit) SetSparkSubmitParameters(v string) *JobDriverSparkSubmit {
	s.SparkSubmitParameters = &v
	return s
}

func (s *JobDriverSparkSubmit) Validate() error {
	return dara.Validate(s)
}
