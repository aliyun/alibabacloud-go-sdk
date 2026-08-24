// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadlockHistogramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetDeadlockHistogramRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetDeadlockHistogramRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetDeadlockHistogramRequest
	GetNodeId() *string
	SetStartTime(v int64) *GetDeadlockHistogramRequest
	GetStartTime() *int64
	SetStatus(v string) *GetDeadlockHistogramRequest
	GetStatus() *string
}

type GetDeadlockHistogramRequest struct {
	// The end of the time range to query. The format is a UNIX timestamp in milliseconds.
	//
	// 	Notice: The value is of the Long type. Precision loss may occur during the serialization/deserialization procedure. The value must not be greater than 9007199254740991.</notice>
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
	// > For PolarDB for MySQL instances, you must specify the node ID.
	//
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. The format is a UNIX timestamp in milliseconds.
	//
	// > The start time can be at most 7 days earlier than the end time.
	//
	// 	Notice: The value is of the Long type. Precision loss may occur during the serialization/deserialization procedure. The value must not be greater than 9007199254740991.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1731983066000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The analysis status of the task.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeadlockHistogramRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeadlockHistogramRequest) GoString() string {
	return s.String()
}

func (s *GetDeadlockHistogramRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDeadlockHistogramRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeadlockHistogramRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDeadlockHistogramRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDeadlockHistogramRequest) GetStatus() *string {
	return s.Status
}

func (s *GetDeadlockHistogramRequest) SetEndTime(v int64) *GetDeadlockHistogramRequest {
	s.EndTime = &v
	return s
}

func (s *GetDeadlockHistogramRequest) SetInstanceId(v string) *GetDeadlockHistogramRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeadlockHistogramRequest) SetNodeId(v string) *GetDeadlockHistogramRequest {
	s.NodeId = &v
	return s
}

func (s *GetDeadlockHistogramRequest) SetStartTime(v int64) *GetDeadlockHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *GetDeadlockHistogramRequest) SetStatus(v string) *GetDeadlockHistogramRequest {
	s.Status = &v
	return s
}

func (s *GetDeadlockHistogramRequest) Validate() error {
	return dara.Validate(s)
}
