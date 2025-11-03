// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceFuncStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceFuncStatusResponseBody
	GetCode() *string
	SetData(v *GetServiceFuncStatusResponseBodyData) *GetServiceFuncStatusResponseBody
	GetData() *GetServiceFuncStatusResponseBodyData
	SetMessage(v string) *GetServiceFuncStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceFuncStatusResponseBody
	GetRequestId() *string
}

type GetServiceFuncStatusResponseBody struct {
	// example:
	//
	// Success
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetServiceFuncStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceFuncStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceFuncStatusResponseBody) GetData() *GetServiceFuncStatusResponseBodyData {
	return s.Data
}

func (s *GetServiceFuncStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceFuncStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceFuncStatusResponseBody) SetCode(v string) *GetServiceFuncStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetData(v *GetServiceFuncStatusResponseBodyData) *GetServiceFuncStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetMessage(v string) *GetServiceFuncStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetRequestId(v string) *GetServiceFuncStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceFuncStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceFuncStatusResponseBodyData struct {
	Args *GetServiceFuncStatusResponseBodyDataArgs `json:"args,omitempty" xml:"args,omitempty" type:"Struct"`
}

func (s GetServiceFuncStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBodyData) GetArgs() *GetServiceFuncStatusResponseBodyDataArgs {
	return s.Args
}

func (s *GetServiceFuncStatusResponseBodyData) SetArgs(v *GetServiceFuncStatusResponseBodyDataArgs) *GetServiceFuncStatusResponseBodyData {
	s.Args = v
	return s
}

func (s *GetServiceFuncStatusResponseBodyData) Validate() error {
	if s.Args != nil {
		if err := s.Args.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceFuncStatusResponseBodyDataArgs struct {
	// example:
	//
	// java
	AddCmd *string `json:"add_cmd,omitempty" xml:"add_cmd,omitempty"`
	// example:
	//
	// true
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
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
	// example:
	//
	// true
	SystemProfiling *string `json:"system_profiling,omitempty" xml:"system_profiling,omitempty"`
}

func (s GetServiceFuncStatusResponseBodyDataArgs) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusResponseBodyDataArgs) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetAddCmd() *string {
	return s.AddCmd
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetCpu() *string {
	return s.Cpu
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetJavaStorePath() *string {
	return s.JavaStorePath
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetLocks() *string {
	return s.Locks
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetLoop() *int32 {
	return s.Loop
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetMem() *string {
	return s.Mem
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) GetSystemProfiling() *string {
	return s.SystemProfiling
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetAddCmd(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.AddCmd = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetCpu(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Cpu = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetJavaStorePath(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.JavaStorePath = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetLocks(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Locks = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetLoop(v int32) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Loop = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetMem(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Mem = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetSystemProfiling(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.SystemProfiling = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) Validate() error {
	return dara.Validate(s)
}
