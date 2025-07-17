// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetStackRequest
	GetEndTime() *int64
	SetPid(v string) *GetStackRequest
	GetPid() *string
	SetRegionId(v string) *GetStackRequest
	GetRegionId() *string
	SetRpcID(v string) *GetStackRequest
	GetRpcID() *string
	SetSpanID(v string) *GetStackRequest
	GetSpanID() *string
	SetStartTime(v int64) *GetStackRequest
	GetStartTime() *int64
	SetTraceID(v string) *GetStackRequest
	GetTraceID() *string
}

type GetStackRequest struct {
	// The exit timestamp of the method call. Unit: milliseconds.
	//
	// example:
	//
	// 1653641800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The process identifier (PID) of the application. For more information about how to obtain the PID, see [Obtain the PID of an application](https://www.alibabacloud.com/help/zh/doc-detail/186100.htm?spm=a2cdw.13409063.0.0.7a72281f0bkTfx#title-imy-7gj-qhr).
	//
	// example:
	//
	// eb4zdose6v@36bab313a******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the remote procedure call (RPC) mode. You can obtain the ID by calling the **GetTrace*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.1
	RpcID *string `json:"RpcID,omitempty" xml:"RpcID,omitempty"`
	// The span ID of the trace. It is displayed on the Trace Explorer page in the ARMS console.
	//
	// example:
	//
	// 88c32dfa4b******
	SpanID *string `json:"SpanID,omitempty" xml:"SpanID,omitempty"`
	// The entry timestamp of the method call. Unit: milliseconds.
	//
	// example:
	//
	// 1653555396
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trace ID. It is displayed on the **Trace Explorer*	- page in the Application Real-Time Monitoring Service (ARMS) console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0a5800611641470044457853******
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetStackRequest) GetPid() *string {
	return s.Pid
}

func (s *GetStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackRequest) GetRpcID() *string {
	return s.RpcID
}

func (s *GetStackRequest) GetSpanID() *string {
	return s.SpanID
}

func (s *GetStackRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetStackRequest) GetTraceID() *string {
	return s.TraceID
}

func (s *GetStackRequest) SetEndTime(v int64) *GetStackRequest {
	s.EndTime = &v
	return s
}

func (s *GetStackRequest) SetPid(v string) *GetStackRequest {
	s.Pid = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetRpcID(v string) *GetStackRequest {
	s.RpcID = &v
	return s
}

func (s *GetStackRequest) SetSpanID(v string) *GetStackRequest {
	s.SpanID = &v
	return s
}

func (s *GetStackRequest) SetStartTime(v int64) *GetStackRequest {
	s.StartTime = &v
	return s
}

func (s *GetStackRequest) SetTraceID(v string) *GetStackRequest {
	s.TraceID = &v
	return s
}

func (s *GetStackRequest) Validate() error {
	return dara.Validate(s)
}
