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
	// The end time of the query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	Notice:
	//
	// This parameter is a Long value. To prevent precision loss during serialization and deserialization, make sure that the value does not exceed 9007199254740991.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732069466000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// > Specify the node ID for a PolarDB for MySQL instance.
	//
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The maximum value is **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source of the task:
	//
	// - **MANUAL*	- or unspecified: queries tasks for recent deadlock analysis.
	//
	// - **AUTO**: queries tasks for full deadlock analysis.
	//
	// 	Notice:
	//
	// If you set this parameter to AUTO to query tasks for full deadlock analysis, the start time can be a maximum of seven days earlier than the end time.
	//
	// example:
	//
	// AUTO
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	Notice:
	//
	// This parameter is a Long value. To prevent precision loss during serialization and deserialization, make sure that the value does not exceed 9007199254740991.
	//
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
