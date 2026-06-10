// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisTool(v string) *StartAIAnalysisRequest
	GetAnalysisTool() *string
	SetAnalysisParams(v []*string) *StartAIAnalysisRequest
	GetAnalysisParams() []*string
	SetChannel(v string) *StartAIAnalysisRequest
	GetChannel() *string
	SetComms(v string) *StartAIAnalysisRequest
	GetComms() *string
	SetCreatedBy(v string) *StartAIAnalysisRequest
	GetCreatedBy() *string
	SetInstance(v string) *StartAIAnalysisRequest
	GetInstance() *string
	SetInstanceType(v string) *StartAIAnalysisRequest
	GetInstanceType() *string
	SetIterationFunc(v string) *StartAIAnalysisRequest
	GetIterationFunc() *string
	SetIterationMod(v string) *StartAIAnalysisRequest
	GetIterationMod() *string
	SetIterationRange(v []*int32) *StartAIAnalysisRequest
	GetIterationRange() []*int32
	SetPids(v string) *StartAIAnalysisRequest
	GetPids() *string
	SetRegion(v string) *StartAIAnalysisRequest
	GetRegion() *string
	SetTimeout(v int32) *StartAIAnalysisRequest
	GetTimeout() *int32
	SetUid(v string) *StartAIAnalysisRequest
	GetUid() *string
}

type StartAIAnalysisRequest struct {
	// Analysis tool. This field does not need to be filled in when using OpenAPI.
	//
	// example:
	//
	// gp/pyki/analysis
	AnalysisTool *string `json:"analysisTool,omitempty" xml:"analysisTool,omitempty"`
	// Data richness
	AnalysisParams []*string `json:"analysis_params,omitempty" xml:"analysis_params,omitempty" type:"Repeated"`
	// Channel name
	//
	// example:
	//
	// ecs_sysom
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// Process name. This is an optional parameter.
	//
	// example:
	//
	// python_test
	Comms *string `json:"comms,omitempty" xml:"comms,omitempty"`
	// Creator. You do not need to specify this parameter when using OpenAPI.
	//
	// example:
	//
	// 445333
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	// Instance ID
	//
	// example:
	//
	// i-wz9dej066kii4goqxxxx
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// Instance type. This field does not need to be filled in when using OpenAPI.
	//
	// example:
	//
	// ip/sn/hostname
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// Iteration entry function. Required only in iteration mode. This parameter is optional.
	//
	// example:
	//
	// Class.function
	IterationFunc *string `json:"iteration_func,omitempty" xml:"iteration_func,omitempty"`
	// Iteration entry module. Required only in iteration mode. This parameter is optional.
	//
	// example:
	//
	// a.b.module
	IterationMod *string `json:"iteration_mod,omitempty" xml:"iteration_mod,omitempty"`
	// Iteration range (iteration count: the number of iterations when the data collection module is activated, independent of the AI job\\"s iteration count)
	IterationRange []*int32 `json:"iteration_range,omitempty" xml:"iteration_range,omitempty" type:"Repeated"`
	// Process IDs (PIDs) of the AI job. Batch input is supported, separated by commas.
	//
	// example:
	//
	// 2421,36547,10043
	Pids *string `json:"pids,omitempty" xml:"pids,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// AI Infra analysis duration.
	//
	// Unit: milliseconds.
	//
	// Default value: 2000.
	//
	// example:
	//
	// 2000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Alibaba Cloud User ID. You do not need to specify this parameter when using OpenAPI.
	//
	// example:
	//
	// 123423414
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s StartAIAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAnalysisRequest) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisRequest) GetAnalysisTool() *string {
	return s.AnalysisTool
}

func (s *StartAIAnalysisRequest) GetAnalysisParams() []*string {
	return s.AnalysisParams
}

func (s *StartAIAnalysisRequest) GetChannel() *string {
	return s.Channel
}

func (s *StartAIAnalysisRequest) GetComms() *string {
	return s.Comms
}

func (s *StartAIAnalysisRequest) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *StartAIAnalysisRequest) GetInstance() *string {
	return s.Instance
}

func (s *StartAIAnalysisRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *StartAIAnalysisRequest) GetIterationFunc() *string {
	return s.IterationFunc
}

func (s *StartAIAnalysisRequest) GetIterationMod() *string {
	return s.IterationMod
}

func (s *StartAIAnalysisRequest) GetIterationRange() []*int32 {
	return s.IterationRange
}

func (s *StartAIAnalysisRequest) GetPids() *string {
	return s.Pids
}

func (s *StartAIAnalysisRequest) GetRegion() *string {
	return s.Region
}

func (s *StartAIAnalysisRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *StartAIAnalysisRequest) GetUid() *string {
	return s.Uid
}

func (s *StartAIAnalysisRequest) SetAnalysisTool(v string) *StartAIAnalysisRequest {
	s.AnalysisTool = &v
	return s
}

func (s *StartAIAnalysisRequest) SetAnalysisParams(v []*string) *StartAIAnalysisRequest {
	s.AnalysisParams = v
	return s
}

func (s *StartAIAnalysisRequest) SetChannel(v string) *StartAIAnalysisRequest {
	s.Channel = &v
	return s
}

func (s *StartAIAnalysisRequest) SetComms(v string) *StartAIAnalysisRequest {
	s.Comms = &v
	return s
}

func (s *StartAIAnalysisRequest) SetCreatedBy(v string) *StartAIAnalysisRequest {
	s.CreatedBy = &v
	return s
}

func (s *StartAIAnalysisRequest) SetInstance(v string) *StartAIAnalysisRequest {
	s.Instance = &v
	return s
}

func (s *StartAIAnalysisRequest) SetInstanceType(v string) *StartAIAnalysisRequest {
	s.InstanceType = &v
	return s
}

func (s *StartAIAnalysisRequest) SetIterationFunc(v string) *StartAIAnalysisRequest {
	s.IterationFunc = &v
	return s
}

func (s *StartAIAnalysisRequest) SetIterationMod(v string) *StartAIAnalysisRequest {
	s.IterationMod = &v
	return s
}

func (s *StartAIAnalysisRequest) SetIterationRange(v []*int32) *StartAIAnalysisRequest {
	s.IterationRange = v
	return s
}

func (s *StartAIAnalysisRequest) SetPids(v string) *StartAIAnalysisRequest {
	s.Pids = &v
	return s
}

func (s *StartAIAnalysisRequest) SetRegion(v string) *StartAIAnalysisRequest {
	s.Region = &v
	return s
}

func (s *StartAIAnalysisRequest) SetTimeout(v int32) *StartAIAnalysisRequest {
	s.Timeout = &v
	return s
}

func (s *StartAIAnalysisRequest) SetUid(v string) *StartAIAnalysisRequest {
	s.Uid = &v
	return s
}

func (s *StartAIAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
