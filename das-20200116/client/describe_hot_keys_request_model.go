// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHotKeysRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeHotKeysRequest
	GetNodeId() *string
}

type DescribeHotKeysRequest struct {
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
}

func (s DescribeHotKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHotKeysRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotKeysRequest) SetInstanceId(v string) *DescribeHotKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHotKeysRequest) SetNodeId(v string) *DescribeHotKeysRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeHotKeysRequest) Validate() error {
	return dara.Validate(s)
}
