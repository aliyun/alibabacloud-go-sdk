// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionThreadDumpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobExecutionThreadDumpResponseBody
	GetCode() *int32
	SetData(v *GetJobExecutionThreadDumpResponseBodyData) *GetJobExecutionThreadDumpResponseBody
	GetData() *GetJobExecutionThreadDumpResponseBodyData
	SetMessage(v string) *GetJobExecutionThreadDumpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobExecutionThreadDumpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobExecutionThreadDumpResponseBody
	GetSuccess() *bool
}

type GetJobExecutionThreadDumpResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetJobExecutionThreadDumpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: appId is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobExecutionThreadDumpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionThreadDumpResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobExecutionThreadDumpResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobExecutionThreadDumpResponseBody) GetData() *GetJobExecutionThreadDumpResponseBodyData {
	return s.Data
}

func (s *GetJobExecutionThreadDumpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobExecutionThreadDumpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobExecutionThreadDumpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobExecutionThreadDumpResponseBody) SetCode(v int32) *GetJobExecutionThreadDumpResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBody) SetData(v *GetJobExecutionThreadDumpResponseBodyData) *GetJobExecutionThreadDumpResponseBody {
	s.Data = v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBody) SetMessage(v string) *GetJobExecutionThreadDumpResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBody) SetRequestId(v string) *GetJobExecutionThreadDumpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBody) SetSuccess(v bool) *GetJobExecutionThreadDumpResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionThreadDumpResponseBodyData struct {
	// example:
	//
	// \\"Thread-7\\" Id=67 TIMED_WAITING\\n\\tat java.base@17.0.5/java.lang.Thread.sleep(Native Method)\\n\\tat app//com.xxl.job.executor.service.jobhandler.SampleXxlJob.shardingJobHandler(SampleXxlJob.java:73)\\n\\tat java.base@17.0.5/jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method)\\n\\tat java.base@17.0.5/jdk.internal.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:77)\\n\\tat java.base@17.0.5/jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\\n\\tat java.base@17.0.5/java.lang.reflect.Method.invoke(Method.java:568)\\n\\tat app//com.xxl.job.core.handler.impl.MethodJobHandler.execute(MethodJobHandler.java:29)\\n\\tat app//com.xxl.job.core.thread.JobThread.run(JobThread.java:152)\\n
	Dump *string `json:"Dump,omitempty" xml:"Dump,omitempty"`
}

func (s GetJobExecutionThreadDumpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionThreadDumpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobExecutionThreadDumpResponseBodyData) GetDump() *string {
	return s.Dump
}

func (s *GetJobExecutionThreadDumpResponseBodyData) SetDump(v string) *GetJobExecutionThreadDumpResponseBodyData {
	s.Dump = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponseBodyData) Validate() error {
	return dara.Validate(s)
}
