// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamVodListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeStreamVodListRequest
	GetEndTime() *int64
	SetId(v string) *DescribeStreamVodListRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeStreamVodListRequest
	GetOwnerId() *int64
	SetStartTime(v int64) *DescribeStreamVodListRequest
	GetStartTime() *int64
}

type DescribeStreamVodListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1634873413
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18526049*****219118918-cn-beijing
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1639077653
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStreamVodListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamVodListRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeStreamVodListRequest) GetId() *string {
	return s.Id
}

func (s *DescribeStreamVodListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStreamVodListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeStreamVodListRequest) SetEndTime(v int64) *DescribeStreamVodListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetId(v string) *DescribeStreamVodListRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetOwnerId(v int64) *DescribeStreamVodListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamVodListRequest) SetStartTime(v int64) *DescribeStreamVodListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeStreamVodListRequest) Validate() error {
	return dara.Validate(s)
}
