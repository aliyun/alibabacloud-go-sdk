// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotBigKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DescribeHotBigKeysRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *DescribeHotBigKeysRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeHotBigKeysRequest
	GetNodeId() *string
}

type DescribeHotBigKeysRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
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
}

func (s DescribeHotBigKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DescribeHotBigKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHotBigKeysRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotBigKeysRequest) SetConsoleContext(v string) *DescribeHotBigKeysRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DescribeHotBigKeysRequest) SetInstanceId(v string) *DescribeHotBigKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHotBigKeysRequest) SetNodeId(v string) *DescribeHotBigKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeHotBigKeysRequest) Validate() error {
	return dara.Validate(s)
}
