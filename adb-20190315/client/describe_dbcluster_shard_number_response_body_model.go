// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterShardNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableShardNumberList(v []*DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) *DescribeDBClusterShardNumberResponseBody
	GetAvailableShardNumberList() []*DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList
	SetAvailableShardNumbers(v []*int32) *DescribeDBClusterShardNumberResponseBody
	GetAvailableShardNumbers() []*int32
	SetRequestId(v string) *DescribeDBClusterShardNumberResponseBody
	GetRequestId() *string
	SetShardNumber(v int32) *DescribeDBClusterShardNumberResponseBody
	GetShardNumber() *int32
}

type DescribeDBClusterShardNumberResponseBody struct {
	// The supported numbers of shards, including the number of current shards and the number of desired shards.
	AvailableShardNumberList []*DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList `json:"AvailableShardNumberList,omitempty" xml:"AvailableShardNumberList,omitempty" type:"Repeated"`
	// The number of desired shards, excluding the number of current shards.
	AvailableShardNumbers []*int32 `json:"AvailableShardNumbers,omitempty" xml:"AvailableShardNumbers,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CBE843D8-964D-5EA3-9D31-822125611B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of shards that you want to change during the data migration.
	//
	// example:
	//
	// 128
	ShardNumber *int32 `json:"ShardNumber,omitempty" xml:"ShardNumber,omitempty"`
}

func (s DescribeDBClusterShardNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterShardNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterShardNumberResponseBody) GetAvailableShardNumberList() []*DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList {
	return s.AvailableShardNumberList
}

func (s *DescribeDBClusterShardNumberResponseBody) GetAvailableShardNumbers() []*int32 {
	return s.AvailableShardNumbers
}

func (s *DescribeDBClusterShardNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterShardNumberResponseBody) GetShardNumber() *int32 {
	return s.ShardNumber
}

func (s *DescribeDBClusterShardNumberResponseBody) SetAvailableShardNumberList(v []*DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) *DescribeDBClusterShardNumberResponseBody {
	s.AvailableShardNumberList = v
	return s
}

func (s *DescribeDBClusterShardNumberResponseBody) SetAvailableShardNumbers(v []*int32) *DescribeDBClusterShardNumberResponseBody {
	s.AvailableShardNumbers = v
	return s
}

func (s *DescribeDBClusterShardNumberResponseBody) SetRequestId(v string) *DescribeDBClusterShardNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterShardNumberResponseBody) SetShardNumber(v int32) *DescribeDBClusterShardNumberResponseBody {
	s.ShardNumber = &v
	return s
}

func (s *DescribeDBClusterShardNumberResponseBody) Validate() error {
	if s.AvailableShardNumberList != nil {
		for _, item := range s.AvailableShardNumberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList struct {
	// The number of shards.
	//
	// example:
	//
	// 128
	ShardNumber *int32 `json:"ShardNumber,omitempty" xml:"ShardNumber,omitempty"`
}

func (s DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) GetShardNumber() *int32 {
	return s.ShardNumber
}

func (s *DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) SetShardNumber(v int32) *DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList {
	s.ShardNumber = &v
	return s
}

func (s *DescribeDBClusterShardNumberResponseBodyAvailableShardNumberList) Validate() error {
	return dara.Validate(s)
}
