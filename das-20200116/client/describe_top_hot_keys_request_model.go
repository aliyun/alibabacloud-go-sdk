// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopHotKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DescribeTopHotKeysRequest
	GetConsoleContext() *string
	SetEndTime(v string) *DescribeTopHotKeysRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeTopHotKeysRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeTopHotKeysRequest
	GetNodeId() *string
	SetStartTime(v string) *DescribeTopHotKeysRequest
	GetStartTime() *string
}

type DescribeTopHotKeysRequest struct {
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
	// r-****-db-0
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

func (s DescribeTopHotKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopHotKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DescribeTopHotKeysRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTopHotKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTopHotKeysRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeTopHotKeysRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTopHotKeysRequest) SetConsoleContext(v string) *DescribeTopHotKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetEndTime(v string) *DescribeTopHotKeysRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetInstanceId(v string) *DescribeTopHotKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetNodeId(v string) *DescribeTopHotKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeTopHotKeysRequest) SetStartTime(v string) *DescribeTopHotKeysRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopHotKeysRequest) Validate() error {
	return dara.Validate(s)
}
