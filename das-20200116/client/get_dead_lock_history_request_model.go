// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetDeadLockHistoryRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetDeadLockHistoryRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetDeadLockHistoryRequest
	GetNodeId() *string
	SetPageNo(v int32) *GetDeadLockHistoryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetDeadLockHistoryRequest
	GetPageSize() *int32
	SetSource(v string) *GetDeadLockHistoryRequest
	GetSource() *string
	SetStartTime(v int64) *GetDeadLockHistoryRequest
	GetStartTime() *int64
}

type GetDeadLockHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1732069466000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AUTO
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731983066000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetDeadLockHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetDeadLockHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDeadLockHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeadLockHistoryRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDeadLockHistoryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetDeadLockHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDeadLockHistoryRequest) GetSource() *string {
	return s.Source
}

func (s *GetDeadLockHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDeadLockHistoryRequest) SetEndTime(v int64) *GetDeadLockHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetInstanceId(v string) *GetDeadLockHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetNodeId(v string) *GetDeadLockHistoryRequest {
	s.NodeId = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetPageNo(v int32) *GetDeadLockHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetPageSize(v int32) *GetDeadLockHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetSource(v string) *GetDeadLockHistoryRequest {
	s.Source = &v
	return s
}

func (s *GetDeadLockHistoryRequest) SetStartTime(v int64) *GetDeadLockHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *GetDeadLockHistoryRequest) Validate() error {
	return dara.Validate(s)
}
