// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *StartAIAnalysisRequest
	GetXDebugId() *string
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
	SetXSysomInvokeSource(v string) *StartAIAnalysisRequest
	GetXSysomInvokeSource() *string
}

type StartAIAnalysisRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The analysis tool. This parameter does not need to be specified when you use OpenAPI.
	//
	// example:
	//
	// gp/pyki/analysis
	AnalysisTool *string `json:"analysisTool,omitempty" xml:"analysisTool,omitempty"`
	// The data richness level.
	AnalysisParams []*string `json:"analysis_params,omitempty" xml:"analysis_params,omitempty" type:"Repeated"`
	// The channel name.
	//
	// example:
	//
	// ecs_sysom
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The process name. This parameter is optional.
	//
	// example:
	//
	// python_test
	Comms *string `json:"comms,omitempty" xml:"comms,omitempty"`
	// The creator. This parameter does not need to be specified when you use OpenAPI.
	//
	// example:
	//
	// 445333
	CreatedBy *string `json:"created_by,omitempty" xml:"created_by,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9dej066kii4goqxxxx
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The instance type. This parameter does not need to be specified when you use OpenAPI.
	//
	// example:
	//
	// ip/sn/hostname
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The iteration entry function. This parameter is required only in iteration mode and can be left empty.
	//
	// example:
	//
	// Class.function
	IterationFunc *string `json:"iteration_func,omitempty" xml:"iteration_func,omitempty"`
	// The iteration entry module. This parameter is required only in iteration mode and can be left empty.
	//
	// example:
	//
	// a.b.module
	IterationMod *string `json:"iteration_mod,omitempty" xml:"iteration_mod,omitempty"`
	// The iteration range. The iteration count refers to the number of iterations when the data collection module is activated, which is independent of the AI job iteration count.
	IterationRange []*int32 `json:"iteration_range,omitempty" xml:"iteration_range,omitempty" type:"Repeated"`
	// The process IDs (PIDs) of the AI job. Multiple PIDs are supported, separated by commas.
	//
	// example:
	//
	// 2421,36547,10043
	Pids *string `json:"pids,omitempty" xml:"pids,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The AI Infra analysis duration. Unit: milliseconds. Default value: 2000.
	//
	// example:
	//
	// 2000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The Alibaba Cloud user ID. This parameter does not need to be specified when you use OpenAPI.
	//
	// example:
	//
	// 123423414
	Uid                *string `json:"uid,omitempty" xml:"uid,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s StartAIAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAnalysisRequest) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *StartAIAnalysisRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *StartAIAnalysisRequest) SetXDebugId(v string) *StartAIAnalysisRequest {
	s.XDebugId = &v
	return s
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

func (s *StartAIAnalysisRequest) SetXSysomInvokeSource(v string) *StartAIAnalysisRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *StartAIAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
