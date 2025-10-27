// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgSize(v int64) *DescribeTableDetailResponseBody
	GetAvgSize() *int64
	SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody
	GetItems() *DescribeTableDetailResponseBodyItems
	SetRequestId(v string) *DescribeTableDetailResponseBody
	GetRequestId() *string
}

type DescribeTableDetailResponseBody struct {
	// The average number of rows in partitions.
	//
	// example:
	//
	// 0
	AvgSize *int64 `json:"AvgSize,omitempty" xml:"AvgSize,omitempty"`
	// The list of partitions.
	Items *DescribeTableDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBody) GetAvgSize() *int64 {
	return s.AvgSize
}

func (s *DescribeTableDetailResponseBody) GetItems() *DescribeTableDetailResponseBodyItems {
	return s.Items
}

func (s *DescribeTableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableDetailResponseBody) SetAvgSize(v int64) *DescribeTableDetailResponseBody {
	s.AvgSize = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableDetailResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTableDetailResponseBodyItems struct {
	Shard []*DescribeTableDetailResponseBodyItemsShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Repeated"`
}

func (s DescribeTableDetailResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItems) GetShard() []*DescribeTableDetailResponseBodyItemsShard {
	return s.Shard
}

func (s *DescribeTableDetailResponseBodyItems) SetShard(v []*DescribeTableDetailResponseBodyItemsShard) *DescribeTableDetailResponseBodyItems {
	s.Shard = v
	return s
}

func (s *DescribeTableDetailResponseBodyItems) Validate() error {
	if s.Shard != nil {
		for _, item := range s.Shard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTableDetailResponseBodyItemsShard struct {
	Id   *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTableDetailResponseBodyItemsShard) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItemsShard) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItemsShard) GetId() *int32 {
	return s.Id
}

func (s *DescribeTableDetailResponseBodyItemsShard) GetSize() *int64 {
	return s.Size
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetId(v int32) *DescribeTableDetailResponseBodyItemsShard {
	s.Id = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetSize(v int64) *DescribeTableDetailResponseBodyItemsShard {
	s.Size = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) Validate() error {
	return dara.Validate(s)
}
