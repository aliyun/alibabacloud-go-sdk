// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopBigKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DescribeTopBigKeysRequest
	GetConsoleContext() *string
	SetEndTime(v string) *DescribeTopBigKeysRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeTopBigKeysRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeTopBigKeysRequest
	GetNodeId() *string
	SetStartTime(v string) *DescribeTopBigKeysRequest
	GetStartTime() *string
}

type DescribeTopBigKeysRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >
	//
	// 	- The end time must be later than the start time.
	//
	// 	- Only data within the last four days can be queried.
	//
	// 	- The maximum interval between the **start time*	- and the*	- end time*	- is 3 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993001
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraDB for Redis instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the data shard on the ApsaraDB for Redis instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopBigKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopBigKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DescribeTopBigKeysRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTopBigKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTopBigKeysRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeTopBigKeysRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTopBigKeysRequest) SetConsoleContext(v string) *DescribeTopBigKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetEndTime(v string) *DescribeTopBigKeysRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetInstanceId(v string) *DescribeTopBigKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetNodeId(v string) *DescribeTopBigKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeTopBigKeysRequest) SetStartTime(v string) *DescribeTopBigKeysRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopBigKeysRequest) Validate() error {
	return dara.Validate(s)
}
