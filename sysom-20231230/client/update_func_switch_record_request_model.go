// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFuncSwitchRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *UpdateFuncSwitchRecordRequest
	GetChannel() *string
	SetParams(v *UpdateFuncSwitchRecordRequestParams) *UpdateFuncSwitchRecordRequest
	GetParams() *UpdateFuncSwitchRecordRequestParams
	SetServiceName(v string) *UpdateFuncSwitchRecordRequest
	GetServiceName() *string
}

type UpdateFuncSwitchRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	Params *UpdateFuncSwitchRecordRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s UpdateFuncSwitchRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequest) GetChannel() *string {
	return s.Channel
}

func (s *UpdateFuncSwitchRecordRequest) GetParams() *UpdateFuncSwitchRecordRequestParams {
	return s.Params
}

func (s *UpdateFuncSwitchRecordRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateFuncSwitchRecordRequest) SetChannel(v string) *UpdateFuncSwitchRecordRequest {
	s.Channel = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequest) SetParams(v *UpdateFuncSwitchRecordRequestParams) *UpdateFuncSwitchRecordRequest {
	s.Params = v
	return s
}

func (s *UpdateFuncSwitchRecordRequest) SetServiceName(v string) *UpdateFuncSwitchRecordRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFuncSwitchRecordRequestParams struct {
	Args *UpdateFuncSwitchRecordRequestParamsArgs `json:"args,omitempty" xml:"args,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// mullprof
	FunctionName *string `json:"function_name,omitempty" xml:"function_name,omitempty"`
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// restart
	Op     *string `json:"op,omitempty" xml:"op,omitempty"`
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1664516888213680
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s UpdateFuncSwitchRecordRequestParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequestParams) GetArgs() *UpdateFuncSwitchRecordRequestParamsArgs {
	return s.Args
}

func (s *UpdateFuncSwitchRecordRequestParams) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UpdateFuncSwitchRecordRequestParams) GetInstance() *string {
	return s.Instance
}

func (s *UpdateFuncSwitchRecordRequestParams) GetOp() *string {
	return s.Op
}

func (s *UpdateFuncSwitchRecordRequestParams) GetRegion() *string {
	return s.Region
}

func (s *UpdateFuncSwitchRecordRequestParams) GetUid() *string {
	return s.Uid
}

func (s *UpdateFuncSwitchRecordRequestParams) SetArgs(v *UpdateFuncSwitchRecordRequestParamsArgs) *UpdateFuncSwitchRecordRequestParams {
	s.Args = v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetFunctionName(v string) *UpdateFuncSwitchRecordRequestParams {
	s.FunctionName = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetInstance(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Instance = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetOp(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Op = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetRegion(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Region = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetUid(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Uid = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) Validate() error {
	if s.Args != nil {
		if err := s.Args.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFuncSwitchRecordRequestParamsArgs struct {
	// example:
	//
	// java
	AddCmd *string `json:"add_cmd,omitempty" xml:"add_cmd,omitempty"`
	// example:
	//
	// true
	Cpu      *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Duration *int32  `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// /tmp/sysom/java-profiler
	JavaStorePath *string `json:"java_store_path,omitempty" xml:"java_store_path,omitempty"`
	// example:
	//
	// true
	Locks *string `json:"locks,omitempty" xml:"locks,omitempty"`
	// example:
	//
	// -1
	Loop *int32 `json:"loop,omitempty" xml:"loop,omitempty"`
	// example:
	//
	// true
	Mem *string `json:"mem,omitempty" xml:"mem,omitempty"`
	Pid *int32  `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// true
	SystemProfiling *string `json:"system_profiling,omitempty" xml:"system_profiling,omitempty"`
}

func (s UpdateFuncSwitchRecordRequestParamsArgs) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequestParamsArgs) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetAddCmd() *string {
	return s.AddCmd
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetCpu() *string {
	return s.Cpu
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetDuration() *int32 {
	return s.Duration
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetJavaStorePath() *string {
	return s.JavaStorePath
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetLocks() *string {
	return s.Locks
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetLoop() *int32 {
	return s.Loop
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetMem() *string {
	return s.Mem
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetPid() *int32 {
	return s.Pid
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) GetSystemProfiling() *string {
	return s.SystemProfiling
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetAddCmd(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.AddCmd = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetCpu(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Cpu = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetDuration(v int32) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Duration = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetJavaStorePath(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.JavaStorePath = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetLocks(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Locks = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetLoop(v int32) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Loop = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetMem(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Mem = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetPid(v int32) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Pid = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetSystemProfiling(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.SystemProfiling = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) Validate() error {
	return dara.Validate(s)
}
