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
	SetStartTime(v int64) *GetStackRequest
	GetStartTime() *int64
	SetTraceID(v string) *GetStackRequest
	GetTraceID() *string
}

type GetStackRequest struct {
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Pid     *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RpcID     *string `json:"RpcID,omitempty" xml:"RpcID,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
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
